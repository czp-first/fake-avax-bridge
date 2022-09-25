package core

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chainlinkfeed"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/contracts"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/sqldb"
	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"
)

func (c *WardenContext) GetCredential(ctx context.Context, in *pb.GetCredentialReq) (*pb.GetCredentialResp, error) {
	err := ioutil.WriteFile(os.Getenv("IdentificationFilePath"), []byte(in.GetIdentification()), 0644)
	if err != nil {
		return nil, err
	}

	credentialJson, err := json.Marshal(c.credential)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.GetCredentialResp{
		Type:       os.Getenv("CredentialType"),
		Credential: string(credentialJson),
	}, nil

}

// enclave get aws credential of warden
func (c *WardenContext) GetAwsCredential(ctx context.Context, in *pb.AwsCredentialReq) (*pb.AwsCredentialResp, error) {
	err := ioutil.WriteFile(os.Getenv("IdentificationFilePath"), []byte(in.GetIdentification()), 0644)
	if err != nil {
		return nil, err
	}

	return &pb.AwsCredentialResp{
		AccessKeyId:     os.Getenv("AWSAccessKeyId"),
		SecretAccessKey: os.Getenv("AWSAccessKeySecret"),
		Region:          os.Getenv("AWSRegion"),
		KmsKeyId:        os.Getenv("AWSKMSKeyId"),
	}, nil
}

// warden save share from enclave
func (c *WardenContext) SaveShare(ctx context.Context, in *pb.SaveShareReq) (*pb.Empty, error) {
	decryptText := c.credential.Decrypt(in.GetShare())

	err := ioutil.WriteFile(os.Getenv("ShareFilePath"), []byte(decryptText), 0644)
	if err != nil {
		return nil, err
	}

	jsonSchema := pulsar.NewJSONSchema(middleware.SettingsSchemaDef, nil)
	producer, err := c.pulsarCli.CreateProducer(pulsar.ProducerOptions{
		Topic:  os.Getenv("PulsarSettingsTopic"),
		Schema: jsonSchema,
	})
	if err != nil {
		log.Errorf("Could not instance producer: %v", err)
		return nil, err
	}
	defer producer.Close()

	fields := &middleware.SettingsJSON{Fields: make([]*middleware.SettingsField, 0)}
	fields.Fields = append(fields.Fields, &middleware.SettingsField{Path: []string{"critical", "walletAddress", "ethereum"}, Value: in.GetOnboardAccountAddress(), Type: "string"})
	fields.Fields = append(fields.Fields, &middleware.SettingsField{Path: []string{"critical", "walletAddress", "dxchain"}, Value: in.GetOffboardAccountAddress(), Type: "string"})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Value: fields,
	})
	if err != nil {
		log.Errorf("produce msg, err:%v", err)
		return nil, err
	}
	log.Info("publish message ok")

	return &pb.Empty{}, nil
}

// enclave request share of warden
func (s *WardenContext) GetShare(ctx context.Context, in *pb.Empty) (*pb.ShareResp, error) {
	identification, err := ioutil.ReadFile(os.Getenv("IdentificationFilePath"))
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(os.Getenv("ShareFilePath"))
	if err != nil {
		return nil, err
	}

	encryptText := s.credential.Encrypt(string(content))

	return &pb.ShareResp{
		Share:          encryptText,
		Identification: string(identification),
	}, nil
}

// enclave get onboard transaction of warden etl
func (s *WardenContext) GetWardenOnboard(ctx context.Context, in *pb.GetWardenOnboardReq) (*pb.GetWardenOnboardResp, error) {

	s.bridgeSettings.InitSettings()
	s.bridgeSettings.Get()

	wardenOnboard, exist, err := s.db.GetWardenOnboardByHash(in.BlockHash, in.TxnHash)
	if err != nil {
		log.Errorf("fail get warden onboard by hash, err:%v", err)
		return nil, err
	}
	if !exist {
		log.Error("get no warden onboard by hash")
		return nil, sqldb.ErrNoRows
	}

	assets := s.bridgeSettings.GetSettings().Critical.Assets
	var asset settings.Asset
	for assetName, assetInfo := range assets {
		if assetInfo.NativeContractAddress.Hex() == wardenOnboard.Contract {
			log.Infof("Get %v warden onboard txn", assetName)
			asset = assetInfo
			break
		}
	}

	// TODO: if not asset

	contract := asset.WrappedContractAddress
	chainlinkFeedAddress := asset.ChainlinkFeedAddress
	// currentTokenPrice := big.NewInt(0)
	var currentTokenPrice *big.Int
	if chainlinkFeedAddress == chain.ZeroAddress {
		currentTokenPrice = s.bridgeSettings.GetSettings().NonCritical.CurrentEthPrice
	} else {
		currentTokenPrice, _ = chainlinkfeed.GetFeedData(chainlinkFeedAddress)
	}
	log.Infof("currentTokenPrice:%v", currentTokenPrice)

	onboardFeeDollars := asset.OnboardFeeDollars
	feeToken := chain.GetOnboardFeeToken(currentTokenPrice, onboardFeeDollars)

	realAmount := big.NewInt(0)
	realAmount.Sub(wardenOnboard.Amount, feeToken)
	log.Infof("amount: %v, fee: %v, realAmount: %v", wardenOnboard.Amount, feeToken, realAmount)

	if realAmount.Cmp(big.NewInt(0)) < 0 {
		log.Fatal("onboard real amount less 0")
	}
	// identification, _ := ioutil.ReadFile(os.Getenv("IdentificationFilePath"))
	client, err := ethclient.Dial(os.Getenv("DxChainHttps"))
	if err != nil {
		log.Fatalf("Oops! There was a problem %s", err)
	} else {
		log.Infoln("Success! You are connected to heco testnet")
	}

	currentGasPrices := s.bridgeSettings.GetCurrentGasPrices()
	suggestedTip := currentGasPrices.Dxchain.SuggestedTip
	gasPrice := chain.ToWei(suggestedTip, 0)

	walletAddress := s.bridgeSettings.GetSettings().Critical.WalletAddress.Dxchain

	// TODO: get from db
	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Fatalf("Fail get pending nonce: %v", err)
	}
	log.Infof("Nonce=%v", nonce)

	content, err := ioutil.ReadFile(os.Getenv("ShareFilePath"))
	if err != nil {
		return nil, err
	}

	encryptText := s.credential.Encrypt(string(content))

	// TODO: rpc problem when nonce equals 0
	return &pb.GetWardenOnboardResp{
		ChainId:    s.bridgeSettings.GetSettings().Critical.Networks.Dxchain.Uint64(),
		RealAmount: realAmount.Uint64(),
		Account:    wardenOnboard.Account,
		GasPrice:   gasPrice.Uint64(),
		Nonce:      nonce,
		Share:      encryptText,
		Contract:   contract.Hex(),
		Fee:        feeToken.Uint64(),
		IsEip1559:  s.bridgeSettings.IsUseEip1559TransactionFormat(),
	}, nil
}

func (s *WardenContext) Onboard(ctx context.Context, in *pb.OnboardReq) (*pb.Empty, error) {

	jsonSchema := pulsar.NewJSONSchema(middleware.OnboardTxnSchemaDef, nil)
	producer, err := s.pulsarCli.CreateProducer(pulsar.ProducerOptions{
		Topic:  os.Getenv("PulsarTopic"),
		Schema: jsonSchema,
	})
	if err != nil {
		log.Fatalf("Could not intance producer: %v", err)
	}
	defer producer.Close()
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Value: &middleware.OnboardTxnJSON{
			Type:            middleware.Enclave,
			BlockHash:       in.BlockHash,
			TxnHash:         in.TxnHash,
			ContractAddress: "",
			AccountAddress:  "",
			BlockNumber:     0,
			TxnIndex:        0,
			Amount:          big.NewInt(0),
			OnboardTxnHash:  in.OnboardTxnHash,
			Nonce:           in.Nonce,
			Batch:           in.Batch,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln("Pubished message")

	return &pb.Empty{}, nil
}

func (s *WardenContext) GetWardenOffboard(ctx context.Context, in *pb.GetWardenOffboardReq) (*pb.GetWardenOffboardResp, error) {
	wardenOffboard, exist, err := s.db.GetWardenOffboardByHash(in.BlockHash, in.TxnHash)
	if err != nil {
		log.Errorf("fail get warden offboard by hash, err:%v", err)
		return nil, err
	}
	if !exist {
		log.Error("get no warden offboard by hash")
		return nil, sqldb.ErrNoRows
	}

	bridgeSettings := s.bridgeSettings.GetSettings()
	var asset settings.Asset
	for assetName, assetInfo := range bridgeSettings.Critical.Assets {
		if assetInfo.WrappedContractAddress.Hex() == wardenOffboard.Contract {
			log.Infof("Get %s warden offboard txn", assetName)
			asset = assetInfo
			break
		}
	}
	// TODO: if not asset

	chainlinkFeedAddress := asset.ChainlinkFeedAddress
	var currentTokenPrice *big.Int
	etherPrice := bridgeSettings.NonCritical.CurrentEthPrice
	if chainlinkFeedAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		currentTokenPrice = etherPrice
	} else {
		currentTokenPrice, _ = chainlinkfeed.GetFeedData(chainlinkFeedAddress)
	}
	log.Infof("currentTokenPrice:%v", currentTokenPrice)

	// TODO
	client, err := ethclient.Dial(os.Getenv("ETHHttps"))
	if err != nil {
		log.Fatalf("Oops! There was a problem %s", err)
	} else {
		log.Infoln("Success! You are connected to ropsten")
	}

	gasPrice := big.NewInt(0)
	gasPrice.Add(
		chain.ToWei(bridgeSettings.NonCritical.CurrentGasPrices.Ethereum.NextBaseFee, 0),
		chain.ToWei(bridgeSettings.NonCritical.CurrentGasPrices.Ethereum.SuggestedTip, 0),
	)

	erc20TokenAbi, _ := abi.JSON(strings.NewReader(contracts.Erc20ABI))
	walletAddress := bridgeSettings.Critical.WalletAddress.Ethereum
	input, err := erc20TokenAbi.Pack("transfer", common.HexToAddress(wardenOffboard.Account), wardenOffboard.Amount)
	if err != nil {
		log.Fatalf("Fail pack transfer: %v", err)
	}
	contract := asset.NativeContractAddress
	futureMsg := ethereum.CallMsg{
		From:  walletAddress,
		To:    &contract,
		Value: big.NewInt(0),
		Data:  input,
	}
	estimateGas, _ := client.EstimateGas(context.TODO(), futureMsg)

	feeToken := chain.GetOffboardFeeToken(etherPrice, currentTokenPrice, gasPrice, asset.OffboardFeeDollars, estimateGas)
	realAmount := big.NewInt(0)
	realAmount.Sub(wardenOffboard.Amount, feeToken)

	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Fatalf("Fail get pending nonce: %v", err)
	}
	log.Infof("Nonce=%v", nonce)

	content, err := ioutil.ReadFile(os.Getenv("ShareFilePath"))
	if err != nil {
		return nil, err
	}

	encryptText := s.credential.Encrypt(string(content))

	// TODO: rpc problem when nonce equals 0
	return &pb.GetWardenOffboardResp{
		ChainId:    bridgeSettings.Critical.Networks.Ethereum.Uint64(),
		RealAmount: realAmount.Uint64(),
		Account:    wardenOffboard.Account,
		GasPrice:   gasPrice.Uint64(),
		Nonce:      nonce,
		Share:      encryptText,
		Contract:   contract.Hex(),
		IsEip1559:  s.bridgeSettings.IsUseEip1559TransactionFormat(),
	}, nil
}

func (s *WardenContext) Offboard(ctx context.Context, in *pb.OffboardReq) (*pb.Empty, error) {
	// s.offboard <- &chains.OffboardTxn{
	// 	Type:            chains.EnclaveTell,
	// 	BlockHash:       in.BlockHash,
	// 	TxnHash:         in.TxnHash,
	// 	OffboardTxnHash: in.OffboardTxnHash,
	// 	Nonce:           in.Nonce,
	// 	Batch:           in.Batch,
	// }
	return &pb.Empty{}, nil
}
