package chains

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type EthClient struct {
	HttpClient *ethclient.Client
	WssClient  *ethclient.Client
}

type TxnType int

const (
	WardenSee TxnType = iota
	EnclaveTell
	WardenTimeout
)

type OnboardTxn struct {
	Type            TxnType
	BlockHash       string
	TxnHash         string
	ContractAddress string
	AccountAddress  string
	ChainId         int64
	BlockNumber     int64
	TxIndex         int
	Amount          int64
	OnboardTxnHash  string
	Nonce           int64
	Batch           int64
}

func NewEthClient() (*EthClient, error) {
	httpClient, err := ethclient.Dial(os.Getenv("ETHHttps"))
	if err != nil {
		log.Fatalf("Fail connect eth https client: %v", err)
		return nil, err
	}

	wssClient, err := ethclient.Dial(os.Getenv("ETHWss"))
	if err != nil {
		log.Fatalf("Fail connect eth wss client: %v", err)
		return nil, err
	}

	log.Infoln("Successfully connected to eth")

	return &EthClient{
		HttpClient: httpClient,
		WssClient:  wssClient,
	}, nil

}

// func (c *EthClient) SeeBlock(query ethereum.FilterQuery, bs settings.BridgeSettingsInterface, ch chan *OnboardTxn) {
// 	logs, err := c.WssClient.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Errorf("Fail filter eth log: %v\n", err)
// 		return
// 	}

// 	for _, vlog := range logs {
// 		// log.Debugf("%+v\n", vlog)
// 		log.Infof("onboard: transfer token[%v] in block[%d] transaction[%v]\n", vlog.Address, vlog.BlockNumber, vlog.TxHash)
// 		erc20Contract, _ := layer1.NewBoundContract(c.WssClient, vlog.Address, contracts.Erc20ABI)
// 		transferEvent := &contracts.Erc20Transfer{}
// 		err = erc20Contract.ParseEvent("Transfer", vlog, transferEvent)
// 		if err != nil {
// 			log.Errorf("Fail to parse transfer event: %v\n", err)
// 			return
// 		}
// 		// log.Infof("%+v\n", transferEvent)
// 		log.Infof("onboard: transfer from[%v] to[%v] amount[%d]\n", transferEvent.From, transferEvent.To, transferEvent.Value)

// 		onboardAssetsConfig := bs.GetOnboardAssetsConfig()
// 		onboardFeeDollars := onboardAssetsConfig.Assets[vlog.Address].FeeDollars

// 		chainlinkFeedAddress := onboardAssetsConfig.Assets[vlog.Address].ChainlinkFeedAddress
// 		// currentTokenPrice := big.NewInt(0)
// 		var currentTokenPrice *big.Int
// 		if chainlinkFeedAddress == common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			currentTokenPrice = onboardAssetsConfig.CurrentEthPrice
// 		} else {
// 			currentTokenPrice, _ = chainlinkfeed.GetFeedData(chainlinkFeedAddress)
// 		}
// 		log.Infof("currentToken[%v], currentTokenPrice[%v]\n", onboardAssetsConfig.Assets[vlog.Address].Name, currentTokenPrice)

// 		feeToken := utils.GetOnboardFeeToken(currentTokenPrice, onboardFeeDollars)
// 		if decimal.NewFromBigInt(transferEvent.Value, 0).Cmp(decimal.NewFromBigInt(feeToken, 0)) <= 0 {
// 			log.Warnf("Insufficient Amount, pass")
// 			continue
// 		}
// 		log.Infoln("aaaaa")
// 		// ch <- &OnboardTxn{
// 		// 	Type:            WardenSee,
// 		// 	BlockHash:       vlog.BlockHash.String(),
// 		// 	TxnHash:         vlog.TxHash.String(),
// 		// 	ContractAddress: vlog.Address.String(),
// 		// 	AccountAddress:  transferEvent.From.String(),
// 		// 	ChainId:         int64(bs.GetSettings().Critical.Networks.Ethereum),
// 		// 	BlockNumber:     query.FromBlock.Int64(),
// 		// 	TxIndex:         int(vlog.TxIndex),
// 		// 	Amount:          transferEvent.Value.Int64(),
// 		// 	Batch:           1,
// 		// }
// 		log.Infoln("dddddd")

// 	}
// }

func (c *EthClient) ConfirmTxn(txnHash string) (bool, error) {
	txn, isPending, err := c.HttpClient.TransactionByHash(context.Background(), common.HexToHash(txnHash))
	if err != nil {
		log.Errorf("Fail get txn by hash: %v\n", err)
		return false, err
	}

	if isPending {
		return false, nil
	}

	receipt, err := c.HttpClient.TransactionReceipt(context.Background(), txn.Hash())
	if err != nil {
		log.Errorf("Fail get txn receipt: %v\n", err)
		return false, err
	}

	if receipt.Status == 1 {
		return true, nil
	}

	return false, nil
}
