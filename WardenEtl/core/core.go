package core

import (
	"WardenEtl/chains"
	"WardenEtl/middleware"
	"WardenEtl/settings"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

type WardenContext struct {
	pulsarCli      pulsar.Client
	bridgeSettings settings.BridgeSettingsInterface
	EthClient      *chains.EthClient
	DxClient       *chains.DxClient
}

func NewWardenContext() (*WardenContext, error) {
	return &WardenContext{}, nil
}

// create pulsar client
func (ctx *WardenContext) initPulsarCli() {
	ctx.pulsarCli = middleware.CreatePulsarClient()
}

// connect eth chain
func (ctx *WardenContext) initEthClient() {
	var err error
	ctx.EthClient, err = chains.NewEthClient()
	if err != nil {
		log.Fatalf("Fail connect eth client: %v\n", err)
	}
}

// connect dx chain
func (ctx *WardenContext) initDxClient() {
	var err error
	ctx.DxClient, err = chains.NewDxClient()
	if err != nil {
		log.Fatalf("Fail connect dx client : %v\n", err)
	}
}

// init bridge settings
func (ctx *WardenContext) initBridgeSettings() {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		log.Fatalf("Fail initialize bridge settings: %v\n", err)
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()

	ctx.bridgeSettings = bridgeSettings
}

func (ctx *WardenContext) Init() {

	log.Infoln("Initializing bridgeSettings...")
	ctx.initBridgeSettings()
	log.Infoln("Successfully initialize bridgeSettings")

	log.Infoln("Connecting eth client...")
	ctx.initEthClient()
	log.Infoln("Successfully connect eth client")

	log.Infoln("Connecting dx client...")
	ctx.initDxClient()
	log.Infoln("Successfully connect dx client")

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
