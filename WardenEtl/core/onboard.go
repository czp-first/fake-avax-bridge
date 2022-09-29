package core

import (
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chainlinkfeed"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/contracts"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/WardenEtl/layer1"

	"context"
	"math/big"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func (ctx *WardenContext) seeFromChainTransfer(query ethereum.FilterQuery) error {
	logs, err := ctx.FromChainClient.WssClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("onboard: Fail filter from chain log: %v", err)
		return err
	}

	for _, vlog := range logs {
		log.Infof("onboard: transfer token[%v] in block[%d] transaction[%v]", vlog.Address, vlog.BlockNumber, vlog.TxHash)
		erc20Contract, err := layer1.NewBoundContract(ctx.FromChainClient.WssClient, vlog.Address, contracts.Erc20ABI)
		if err != nil {
			log.Errorf("onboard: fail bound contract: %v", err)
			return err
		}
		transferEvent := &contracts.Erc20Transfer{}
		err = erc20Contract.ParseEvent("Transfer", vlog, transferEvent)
		if err != nil {
			log.Errorf("onboard: Fail to parse transfer event: %v", err)
			return err
		}
		log.Infof("onboard: transfer from[%v] to[%v] amount[%d]", transferEvent.From, transferEvent.To, transferEvent.Value)

		onboardAssetsConfig := ctx.bridgeSettings.GetOnboardAssetsConfig()
		asset, ok := onboardAssetsConfig.Assets[vlog.Address]
		// this if is not necessary because query already filter tokens
		if !ok {
			log.Warnf("onboard: token[%v] not in birdge onboard list", vlog.Address)
			continue
		}

		currentTokenPrice := onboardAssetsConfig.CurrentEthPrice
		if asset.ChainlinkFeedAddress != chain.ZeroAddress {
			currentTokenPrice, err = chainlinkfeed.GetFeedData(asset.ChainlinkFeedAddress)
			if err != nil {
				log.Errorf("onboard: fail get token[%v] price: %v", vlog.Address, err)
				return err
			}
		}
		log.Infof("onboard: currentToken[%v], currentTokenPrice[%v]", onboardAssetsConfig.Assets[vlog.Address].Name, currentTokenPrice)

		feeAmount := chain.GetOnboardFeeToken(currentTokenPrice, asset.FeeDollars)
		log.Infof("onboard: fee amount %d", feeAmount)
		if decimal.NewFromBigInt(transferEvent.Value, 0).Cmp(decimal.NewFromBigInt(feeAmount, 0)) <= 0 {
			log.Warnf("onboard: Insufficient Amount, feeAmount[%d]>amount[%d], pass", feeAmount, transferEvent.Value)
			continue
		}

		jsonSchema := pulsar.NewJSONSchema(middleware.OnboardTxnSchemaDef, nil)

		onboardTxnProducer, err := ctx.pulsarCli.CreateProducer(pulsar.ProducerOptions{
			Topic:  os.Getenv("PulsarOnboardTopic"),
			Schema: jsonSchema,
		})
		if err != nil {
			log.Errorf("Could not instance onboardTxn producer: %v", err)
			return err
		}
		defer onboardTxnProducer.Close()
		_, err = onboardTxnProducer.Send(context.Background(), &pulsar.ProducerMessage{
			Value: &middleware.OnboardTxnJSON{
				Type:            middleware.Warden,
				BlockHash:       vlog.BlockHash.String(),
				TxnHash:         vlog.TxHash.String(),
				ContractAddress: vlog.Address.String(),
				AccountAddress:  transferEvent.From.String(),
				ChainId:         ctx.bridgeSettings.GetSettings().Critical.Networks.Ethereum,
				BlockNumber:     query.FromBlock.Uint64(),
				TxnIndex:        vlog.TxIndex,
				Amount:          transferEvent.Value,
				Batch:           1,
			},
		})

		if err != nil {
			log.Errorf("onboard: publish %s message err: %v", os.Getenv("PulsarOnboardTopic"), err)
			return err
		}
		log.Infoln("onboard: published %s message", os.Getenv("PulsarOnboardTopic"))
	}
	return nil
}

func (ctx *WardenContext) seeFromChainBlock() error {

	log.Info("onboard: get settings")
	ctx.bridgeSettings.InitSettings()
	ctx.bridgeSettings.Get()
	ctx.bridgeSettings.NewOnboardSettingsJSON()

	preSettings := ctx.bridgeSettings.GetSettings()
	fromChainMinConfirmations := preSettings.NonCritical.MinimumConfirmations.Ethereum
	preFromChainLastIndexedBlk := preSettings.NonCritical.NetworkViews.Ethereum.LastIndexedBlock
	preFromChainLastSeenBlk := preSettings.NonCritical.NetworkViews.Ethereum.LastSeenBlock
	chainlinkEthUsdFeedAddress := preSettings.NonCritical.ChainlinkEthUsdFeedAddress

	fromChainLatestBlockNum, err := ctx.FromChainClient.HttpClient.BlockNumber(context.Background())
	if err != nil {
		log.Errorf("onboard: Get from chain blockNumber err: %v", err)
		return err
	}
	log.Infof(
		"onboard: fromChainMinConfirmations[%d], preFromChainLastIndexedBlock[%d], preFromChainLastSeenBlock[%d], fromChainLatestBlock[%d]",
		fromChainMinConfirmations, preFromChainLastIndexedBlk, preFromChainLastSeenBlk, fromChainLatestBlockNum,
	)

	currentFromChainPrice, err := chainlinkfeed.GetFeedData(chainlinkEthUsdFeedAddress)
	if err != nil {
		log.Errorf("onboard: get current from chain native price err: %v", err)
		return err
	}
	log.Infof("onboard: current from chain native price: %v", currentFromChainPrice)
	onboardAssetsConfig := ctx.bridgeSettings.GetOnboardAssetsConfig()
	onboardAssetsConfig.CurrentEthPrice = currentFromChainPrice
	ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{
		Path:  []string{"nonCritical", "currentEthPrice"},
		Value: currentFromChainPrice.String(),
		Type:  "int",
	})

	// preFromChainLastIndexedBlk = 12714289 - 1 // weth
	// preFromChainLastIndexedBlk = 12731462 - 1 // dx
	if (fromChainLatestBlockNum - preFromChainLastIndexedBlk) > fromChainMinConfirmations {
		readyFromChainIndexBlk := preFromChainLastIndexedBlk + 1
		log.Infof("onboard: Ready index from chain block[%d]", readyFromChainIndexBlk)

		tokensAddr := ctx.bridgeSettings.GetTokenAddrs(preSettings.Critical.Networks.Ethereum)
		receiveAddressHash := preSettings.Critical.WalletAddress.Ethereum
		addTopic := [][]common.Hash{{}, {receiveAddressHash.Hash()}}
		query, err := layer1.MakeFilterQuery(tokensAddr, contracts.Erc20ABI, "Transfer", big.NewInt(int64(readyFromChainIndexBlk)), big.NewInt(int64(readyFromChainIndexBlk)), addTopic)
		if err != nil {
			log.Errorf("onboard: make filter query err: %v", err)
			return err
		}

		err = ctx.seeFromChainTransfer(query)
		if err != nil {
			log.Errorf("onboard: see from chain transfer err: %v", err)
			return err
		}
		ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{Path: []string{"nonCritical", "networkViews", "ethereum", "lastIndexedBlock"}, Value: big.NewInt(int64(readyFromChainIndexBlk)).String(), Type: "int"})
		log.Infof("onboard: Finish index from chain block[%d]", readyFromChainIndexBlk)
	}

	ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{Path: []string{"nonCritical", "networkViews", "ethereum", "lastSeenBlock"}, Value: big.NewInt(int64(fromChainLatestBlockNum)).String(), Type: "int"})
	ctx.bridgeSettings.ProduceUpdate(ctx.pulsarCli, true)
	return nil

}

func (ctx *WardenContext) SeeFromChainBlock() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := ctx.seeFromChainBlock()
		if err != nil {
			log.Errorf("see from chain block err:%v", err)
			break
		}
	}
}
