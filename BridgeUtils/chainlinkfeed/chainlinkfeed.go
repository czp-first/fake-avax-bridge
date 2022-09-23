package chainlinkfeed

import (
	"math/big"
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// https://avawarden-prod.s3.amazonaws.com/bridge_settings.json
// aave: 0x6df09e975c830ecae5bd4ed9d90f3a95a4f88012
func GetFeedData(contractAddr common.Address) (*big.Int, error) {

	client, err := ethclient.Dial(os.Getenv("MainnetETHHttps"))
	if err != nil {
		log.Errorf("Fail connect mainnet ethclient: %v", err)
		return nil, err
	}
	chainLinkFeedContract, err := contracts.NewChainlink(contractAddr, client)
	if err != nil {
		log.Errorf("Fail new chainlink contract: %v", err)
		return nil, err
	}
	feedData, err := chainLinkFeedContract.LatestRoundData(nil)
	if err != nil {
		log.Errorf("Fail query chainlink feed: %v", err)
		return nil, err
	}
	return feedData.Answer, nil
}
