package core

import (
	"WardenConsumer/database"
	"WardenConsumer/middleware"
	"WardenConsumer/settings"

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
	ctx.pulsarCli = middleware.CreateClient()
}

func (ctx *ConsumerContext) initBridgeSettings() {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		log.Fatalf("fail initialize bridge settings: %v", err)
	}
	bridgeSettings := bridgeSettingsFactory.MakeBridgeSettings()
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
