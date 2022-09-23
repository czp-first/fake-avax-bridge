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

func NewDxClient() (*DxClient, error) {
	httpClient, err := ethclient.Dial(os.Getenv("DxChainHttps"))
	if err != nil {
		log.Fatalf("Fail connect dxchain https client: %v\n", err)
	}

	wssClient, err := ethclient.Dial(os.Getenv("DxChainWss"))
	if err != nil {
		log.Fatalf("Fail connect dxchain wss client: %v\n", err)
	}
	log.Infoln("Successfully connected to dxchain")

	return &DxClient{
		HttpClient: httpClient,
		WssClient:  wssClient,
	}, nil
}

func (dx *DxClient) ConfirmTxn(txnHash string) (bool, error) {
	txn, isPending, err := dx.HttpClient.TransactionByHash(context.Background(), common.HexToHash(txnHash))
	if err != nil {
		log.Errorf("fail get txn by hash: %v", err)
		return false, err
	}

	if isPending {
		return false, nil
	}

	receipt, err := dx.HttpClient.TransactionReceipt(context.Background(), txn.Hash())
	if err != nil {
		log.Errorf("Fail get txn receipt: %v", err)
		return false, err
	}

	if receipt.Status == 1 {
		return true, nil
	}

	return false, nil
}
