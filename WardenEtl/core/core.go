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
func (ctx *WardenContext) initPulsarCli() {
	ctx.pulsarCli = middleware.CreatePulsarClient()
}

// connect from chain
func (ctx *WardenContext) initFromChainClient() {
	var err error
	ctx.FromChainClient, err = chain.NewChainClient(os.Getenv("FromChainHttps"), os.Getenv("FromChainWss"))
	if err != nil {
		log.Fatalf("Fail connect from chain client: %v", err)
	}
}

// connect to chain
func (ctx *WardenContext) initToChainClient() {
	var err error
	ctx.ToChainClient, err = chain.NewChainClient(os.Getenv("ToChainHttps"), os.Getenv("ToChainWss"))
	if err != nil {
		log.Fatalf("Fail connect to chain client : %v", err)
	}
}

// init bridge settings
func (ctx *WardenContext) initBridgeSettings() {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		log.Fatalf("Fail initialize bridge settings: %v", err)
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()

	ctx.bridgeSettings = bridgeSettings
}

func (ctx *WardenContext) Init() {

	log.Infoln("Initializing bridgeSettings...")
	ctx.initBridgeSettings()
	log.Infoln("Successfully initialize bridgeSettings")

	log.Infoln("Connecting from chain client...")
	ctx.initFromChainClient()
	log.Infoln("Successfully connect from chain client")

	log.Infoln("Connecting to chain client...")
	ctx.initToChainClient()
	log.Infoln("Successfully connect to chain client")

	log.Infoln("Instancing pulsar client...")
	ctx.initPulsarCli()
	log.Infoln("Successfully initialize pulsar client")
}

// func Close(ctx *WardenContext) error {
// 	err := ctx.db.Close()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
