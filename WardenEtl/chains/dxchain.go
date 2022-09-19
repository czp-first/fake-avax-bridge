package chains

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type DxClient struct {
	HttpClient *ethclient.Client
	WssClient  *ethclient.Client
}

type OffboardTxn struct {
	Type            TxnType
	BlockHash       string
	TxnHash         string
	ContractAddress string
	AccountAddress  string
	ChainId         int64
	BlockNumber     int64
	TxIndex         int
	Amount          int64
	OffboardTxnHash string
	Nonce           int64
	Batch           int64
}

func NewDxClient() (*DxClient, error) {
	httpClient, err := ethclient.Dial(os.Getenv("DxChainHttps"))
	if err != nil {
		log.Fatalf("Fail connect dxchain https client: %v", err)
		return nil, err
	}

	wssClient, err := ethclient.Dial(os.Getenv("DxChainWss"))
	if err != nil {
		log.Fatalf("Fail connect dxchain wss client: %v", err)
		return nil, err
	}
	log.Infoln("Successfully connected to dxchain")

	return &DxClient{
		HttpClient: httpClient,
		WssClient:  wssClient,
	}, nil
}

// func (dx *DxClient) SeeBlock(query ethereum.FilterQuery, bs *settings.BridgeSettings, ch chan *OffboardTxn) {

// 	logs, err := dx.WssClient.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Errorf("Fail filter dxchain log: %v\n", err)
// 	}

// 	for _, vlog := range logs {

// 		dxErc20Contract, _ := layer1.NewBoundContract(dx.WssClient, vlog.Address, contracts.DxErc20MetaData.ABI)
// 		unwrapEvent := &contracts.DxErc20Unwrap{}

// 		err = dxErc20Contract.ParseEvent("Unwrap", vlog, unwrapEvent)
// 		if err != nil {
// 			log.Errorf("Fail to parse unwrap event: %v\n", err)
// 		}

// 		offboardFeeDollars := bs.OffboardAssetsConfig.Assets[vlog.Address].FeeDollars
// 		chainlinkFeedAddress := bs.OffboardAssetsConfig.Assets[vlog.Address].ChainlinkFeedAddress
// 		currentTokenPrice := big.NewInt(0)
// 		if chainlinkFeedAddress == common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			currentTokenPrice = bs.OffboardAssetsConfig.CurrentEthPrice
// 		} else {
// 			currentTokenPrice, _ = chainlinkfeed.GetFeedData(chainlinkFeedAddress)
// 		}
// 		log.Infof("currentToken[%v], currentTokenPrice[%v]\n", bs.OffboardAssetsConfig.Assets[vlog.Address].Name, currentTokenPrice)

// 		chainId := bs.Networks.DxChain

// 		txn, _, _ := dx.HttpClient.TransactionByHash(context.Background(), vlog.TxHash)

// 		// TODO: diff signer with txn type
// 		msg, err := txn.AsMessage(types.NewLondonSigner(big.NewInt(int64(chainId))), nil)
// 		if err != nil {
// 			log.Errorf("Turn txn to msg: %v\n", err)
// 		}

// 		// cal offboard fee
// 		client, err := ethclient.Dial(os.Getenv("ETHHttps"))
// 		if err != nil {
// 			log.Fatalf("Oops! There was a problem %s\n", err)
// 		} else {
// 			log.Infoln("Success! You are connected to Netwrok")
// 		}

// 		gasPrice, err := client.SuggestGasPrice(context.Background())
// 		if err != nil {
// 			log.Fatalf("Fail get gas price: %v\n", err)
// 		}
// 		erc20TokenAbi, err := abi.JSON(strings.NewReader(contracts.Erc20ABI))

// 		input, err := erc20TokenAbi.Pack("transfer", msg.From(), unwrapEvent.Amount)
// 		if err != nil {
// 			log.Fatalf("Fail pack transfer: %v\n", err)
// 		}
// 		futureMsg := ethereum.CallMsg{
// 			From:  bs.Wallets.DxChain,
// 			To:    &vlog.Address,
// 			Value: big.NewInt(0),
// 			Data:  input,
// 		}
// 		estimateGas, err := client.EstimateGas(context.TODO(), futureMsg)

// 		feeToken := utils.GetOffboardFeeToken(bs.OffboardAssetsConfig.CurrentEthPrice, currentTokenPrice, gasPrice, offboardFeeDollars, estimateGas)

// 		log.Infof("amount[%v], fee[%v]\n", unwrapEvent.Amount, feeToken)
// 		if decimal.NewFromBigInt(unwrapEvent.Amount, 0).Cmp(decimal.NewFromBigInt(feeToken, 0)) <= 0 {
// 			log.Warnf("Insufficient Amount, pass")
// 			continue
// 		}

// 		log.Infof("account: %v\n", msg.From().String())
// 		ch <- &OffboardTxn{
// 			Type:            WardenSee,
// 			BlockHash:       vlog.BlockHash.String(),
// 			TxnHash:         vlog.TxHash.String(),
// 			ContractAddress: vlog.Address.String(),
// 			AccountAddress:  msg.From().String(),
// 			ChainId:         int64(chainId),
// 			BlockNumber:     query.FromBlock.Int64(),
// 			TxIndex:         int(vlog.TxIndex),
// 			Amount:          unwrapEvent.Amount.Int64(),
// 			Batch:           1,
// 		}
// 	}
// }

func (dx *DxClient) ConfirmTxn(txnHash string) (bool, error) {
	txn, isPending, err := dx.HttpClient.TransactionByHash(context.Background(), common.HexToHash(txnHash))
	if err != nil {
		log.Errorf("Fail get txn by hash: %v\n", err)
		return false, err
	}

	if isPending {
		return false, nil
	}

	receipt, err := dx.HttpClient.TransactionReceipt(context.Background(), txn.Hash())
	if err != nil {
		log.Errorf("Fail get txn receipt: %v\n", err)
		return false, err
	}

	if receipt.Status == 1 {
		return true, nil
	}

	return false, nil
}
