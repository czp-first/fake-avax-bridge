package chain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type ChainClient struct {
	HttpClient *ethclient.Client
	WssClient  *ethclient.Client
}

func NewChainClient(chainHttp, chainWss string) (*ChainClient, error) {
	httpClient, err := ethclient.Dial(chainHttp)
	if err != nil {
		log.Fatalf("Fail connect https client: %v", err)
		return nil, err
	}

	wssClient, err := ethclient.Dial(chainWss)
	if err != nil {
		log.Fatalf("Fail connect wss client: %v", err)
		return nil, err
	}

	log.Info("Successfully connected to chain client")

	return &ChainClient{
		HttpClient: httpClient,
		WssClient:  wssClient,
	}, nil
}

func (c *ChainClient) ConfirmTxn(txnHash string) (bool, error) {
	txn, isPending, err := c.HttpClient.TransactionByHash(context.Background(), common.HexToHash(txnHash))
	if err != nil {
		log.Errorf("Fail get txn by hash: %v", err)
		return false, err
	}

	if isPending {
		return false, nil
	}

	receipt, err := c.HttpClient.TransactionReceipt(context.Background(), txn.Hash())
	if err != nil {
		log.Errorf("Fail get txn receipt: %v", err)
		return false, err
	}

	if receipt.Status == 1 {
		return true, nil
	}

	return false, nil
}
