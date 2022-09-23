package chains

import (
	"os"

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

func NewEthClient() (*EthClient, error) {
	httpClient, err := ethclient.Dial(os.Getenv("ETHHttps"))
	if err != nil {
		log.Errorf("Fail connect eth https client: %v\n", err)
		return nil, err
	}

	wssClient, err := ethclient.Dial(os.Getenv("ETHWss"))
	if err != nil {
		log.Errorf("Fail connect eth wss client: %v\n", err)
		return nil, err
	}

	log.Infoln("Successfully connected to eth")

	return &EthClient{
		HttpClient: httpClient,
		WssClient:  wssClient,
	}, nil

}
