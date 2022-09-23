package chainlinkfeed

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/WardenServer/contracts"
)

// https://avawarden-prod.s3.amazonaws.com/bridge_settings.json
// aave: 0x6df09e975c830ecae5bd4ed9d90f3a95a4f88012
func GetFeedData(contractAddr common.Address) (*big.Int, error) {
	// chainLinkFeedContractAddr := common.HexToAddress(contractAddr)

	client, err := ethclient.Dial(os.Getenv("MainnetETHHttps"))
	if err != nil {
		logrus.Errorf("Fail connect mainnet ethclient: %v\n", err)
		return nil, err
	}
	chainLinkFeedContract, err := contracts.NewChainlink(contractAddr, client)
	if err != nil {
		logrus.Errorf("Fail new chainlink contract: %v\n", err)
		return nil, err
	}
	feedData, err := chainLinkFeedContract.LatestRoundData(nil)
	if err != nil {
		logrus.Errorf("Fail query chainlink feed: %v\n", err)
		return nil, err
	}
	return feedData.Answer, nil
}
