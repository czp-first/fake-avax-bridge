package core

import (
	"github.com/czp-first/fake-avax-bridge/WardenConsumer/database"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

type ConsumerContext struct {
	db             *database.PgSQL
	pulsarCli      pulsar.Client
	bridgeSettings settings.BridgeSettingsInterface
}

func NewConsumerContext() (*ConsumerContext, error) {
	return &ConsumerContext{}, nil
}

func (ctx *ConsumerContext) initDb() {
	var err error
	ctx.db, err = database.NewPgSQL()
	if err != nil {
		log.Fatal(err)
	}
}

func (ctx *ConsumerContext) initPulsarCli() {
	ctx.pulsarCli = middleware.CreatePulsarClient()
}

func (ctx *ConsumerContext) initBridgeSettings() {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		log.Fatalf("fail initialize bridge settings: %v", err)
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()
	ctx.bridgeSettings = bridgeSettings
}

func (ctx *ConsumerContext) Init() {
	log.Infoln("Initializing DB...")
	ctx.initDb()
	log.Infoln("Successfully initialize DB")

	log.Infoln("Instancing pulsar client...")
	ctx.initPulsarCli()
	log.Infoln("Successfully initialize pulsar client")

	log.Infoln("Initializing bridgeSettings...")
	ctx.initBridgeSettings()
	log.Infoln("Successfully initialize bridgeSettings")
}

func (ctx *ConsumerContext) Close() {
	ctx.db.Close()
	ctx.pulsarCli.Close()
}
