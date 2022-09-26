package core

import (
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

/*
for example: if onboard's direction is from ethereum to heco, heco is to chain, ethereum is from chain
*/
type WardenContext struct {
	pulsarCli       pulsar.Client
	bridgeSettings  settings.BridgeSettingsInterface
	FromChainClient *chain.ChainClient
	ToChainClient   *chain.ChainClient
}

func NewWardenContext() (*WardenContext, error) {
	return &WardenContext{}, nil
}

// create pulsar client
func (ctx *WardenContext) initPulsarCli() error {
	pulsarClient, err := middleware.CreatePulsarClient(os.Getenv("PulsarURL"))
	if err != nil {
		return err
	}
	ctx.pulsarCli = pulsarClient
	return nil
}

// connect from chain
func (ctx *WardenContext) initFromChainClient() error {
	var err error
	ctx.FromChainClient, err = chain.NewChainClient(os.Getenv("FromChainHttps"), os.Getenv("FromChainWss"))
	return err
}

// connect to chain
func (ctx *WardenContext) initToChainClient() error {
	var err error
	ctx.ToChainClient, err = chain.NewChainClient(os.Getenv("ToChainHttps"), os.Getenv("ToChainWss"))
	return err
}

// init bridge settings
func (ctx *WardenContext) initBridgeSettings() error {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		return err
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()

	ctx.bridgeSettings = bridgeSettings
	return nil
}

func (ctx *WardenContext) Init() error {
	var err error
	log.Infoln("Initializing bridgeSettings...")
	err = ctx.initBridgeSettings()
	if err != nil {
		log.Errorf("fail init bridge settings, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize bridgeSettings")

	log.Infoln("Connecting from chain client...")
	err = ctx.initFromChainClient()
	if err != nil {
		log.Errorf("fail connect from chain, err:%v", err)
		return err
	}
	log.Infoln("Successfully connect from chain client")

	log.Infoln("Connecting to chain client...")
	err = ctx.initToChainClient()
	if err != nil {
		log.Errorf("fail connect to chain, err:%v", err)
		return err
	}
	log.Infoln("Successfully connect to chain client")

	log.Infoln("Instancing pulsar client...")
	err = ctx.initPulsarCli()
	if err != nil {
		log.Errorf("fail create pulsar client, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize pulsar client")
	return nil
}

// func Close(ctx *WardenContext) error {
// 	err := ctx.db.Close()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
