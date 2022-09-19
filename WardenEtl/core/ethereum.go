package core

import (
	"WardenEtl/chainlinkfeed"
	"WardenEtl/consts"
	"WardenEtl/contracts"
	"WardenEtl/layer1"
	"WardenEtl/middleware"
	"WardenEtl/utils"

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

func (ctx *WardenContext) seeEthTransfer(query ethereum.FilterQuery) error {
	logs, err := ctx.EthClient.WssClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("onboard: Fail filter eth log: %v", err)
		return err
	}

	for _, vlog := range logs {
		log.Infof("onboard: transfer token[%v] in block[%d] transaction[%v]", vlog.Address, vlog.BlockNumber, vlog.TxHash)
		erc20Contract, err := layer1.NewBoundContract(ctx.EthClient.WssClient, vlog.Address, contracts.Erc20ABI)
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
		if asset.ChainlinkFeedAddress != consts.ZeroAddress {
			currentTokenPrice, err = chainlinkfeed.GetFeedData(asset.ChainlinkFeedAddress)
			if err != nil {
				log.Errorf("onboard: fail get token[%v] price: %v", vlog.Address, err)
				return err
			}
		}
		log.Infof("onboard: currentToken[%v], currentTokenPrice[%v]", onboardAssetsConfig.Assets[vlog.Address].Name, currentTokenPrice)

		feeAmount := utils.GetOnboardFeeToken(currentTokenPrice, asset.FeeDollars)
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

func (ctx *WardenContext) seeEthBlock() error {

	log.Info("onboard: get settings")
	ctx.bridgeSettings.InitSettings()
	ctx.bridgeSettings.Get()
	ctx.bridgeSettings.NewOnboardSettingsJSON()

	preSettings := ctx.bridgeSettings.GetSettings()
	ethMinConfirmations := preSettings.NonCritical.MinimumConfirmations.Ethereum
	preEthLastIndexedBlk := preSettings.NonCritical.NetworkViews.Ethereum.LastIndexedBlock
	preEthLastSeenBlk := preSettings.NonCritical.NetworkViews.Ethereum.LastSeenBlock
	chainlinkEthUsdFeedAddress := preSettings.NonCritical.ChainlinkEthUsdFeedAddress

	ethLatestBlockNum, err := ctx.EthClient.HttpClient.BlockNumber(context.Background())
	if err != nil {
		log.Errorf("onboard: Get eth blockNumber err: %v", err)
		return err
	}
	log.Infof(
		"onboard: ethMinConfirmations[%d], preEthLastIndexedBlock[%d], preEthLastSeenBlock[%d], ethLatestBlock[%d]",
		ethMinConfirmations, preEthLastIndexedBlk, preEthLastSeenBlk, ethLatestBlockNum,
	)

	currentEthPrice, err := chainlinkfeed.GetFeedData(chainlinkEthUsdFeedAddress)
	if err != nil {
		log.Errorf("onboard: get current eth price err: %v", err)
		return err
	}
	log.Infof("onboard: Current eth price: %v", currentEthPrice)
	onboardAssetsConfig := ctx.bridgeSettings.GetOnboardAssetsConfig()
	onboardAssetsConfig.CurrentEthPrice = currentEthPrice
	ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{
		Path:  []string{"nonCritical", "currentEthPrice"},
		Value: currentEthPrice.String(),
		Type:  "int",
	})

	preEthLastIndexedBlk = 12714289 - 1 // weth
	// preEthLastIndexedBlk = 12731462 - 1 // dx
	if (ethLatestBlockNum - preEthLastIndexedBlk) > ethMinConfirmations {
		readyEthIndexBlk := preEthLastIndexedBlk + 1
		log.Infof("onboard: Ready index eth block[%d]", readyEthIndexBlk)

		tokensAddr := ctx.bridgeSettings.GetTokenAddrs(preSettings.Critical.Networks.Ethereum)
		receiveAddressHash := preSettings.Critical.WalletAddress.Ethereum
		addTopic := [][]common.Hash{{}, {receiveAddressHash.Hash()}}
		query, err := layer1.MakeFilterQuery(tokensAddr, contracts.Erc20ABI, "Transfer", big.NewInt(int64(readyEthIndexBlk)), big.NewInt(int64(readyEthIndexBlk)), addTopic)
		if err != nil {
			log.Errorf("onboard: make filter query err: %v", err)
			return err
		}

		err = ctx.seeEthTransfer(query)
		if err != nil {
			log.Errorf("onboard: see eth transfer err: %v", err)
			return err
		}
		ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{Path: []string{"nonCritical", "networkViews", "ethereum", "lastIndexedBlock"}, Value: big.NewInt(int64(readyEthIndexBlk)).String(), Type: "int"})
		log.Infof("onboard: Finish index eth block[%d]", readyEthIndexBlk)
	}

	ctx.bridgeSettings.AppendOnboardSettingsJSON(&middleware.SettingsField{Path: []string{"nonCritical", "networkViews", "ethereum", "lastSeenBlock"}, Value: big.NewInt(int64(ethLatestBlockNum)).String(), Type: "int"})
	ctx.bridgeSettings.Update(ctx.pulsarCli, true)
	return nil

}

func (ctx *WardenContext) SeeEthBlock() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := ctx.seeEthBlock()
		if err != nil {
			break
		}
		break
	}
}
