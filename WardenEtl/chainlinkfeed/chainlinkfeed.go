package chainlinkfeed

import (
	"WardenEtl/contracts"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

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
