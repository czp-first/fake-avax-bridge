package core

import (
	"fmt"
	"os"

	"github.com/czp-first/fake-avax-bridge/WardenConsumer/dal"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/sqldb"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

type ConsumerContext struct {
	db             *dal.DAL
	pulsarCli      pulsar.Client
	bridgeSettings settings.BridgeSettingsInterface
}

func NewConsumerContext() (*ConsumerContext, error) {
	return &ConsumerContext{}, nil
}

func (ctx *ConsumerContext) initDb() error {
	var err error
	// init db
	ctx.db, err = dal.NewDAL(
		sqldb.PostgresDriver,
		fmt.Sprint(
			sqldb.PostgresFmt,
			os.Getenv("PgUser"),
			os.Getenv("PgPassword"),
			fmt.Sprintf("%s:%s", os.Getenv("PgHost"), os.Getenv("PgPort")), os.Getenv("PgDb")),
		sqldb.DefaultPostgresPoolSize)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *ConsumerContext) initPulsarCli() error {
	puslarCli, err := middleware.CreatePulsarClient()
	if err != nil {
		return err
	}
	ctx.pulsarCli = puslarCli
	return nil
}

func (ctx *ConsumerContext) initBridgeSettings() error {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		return err
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()
	bridgeSettings.InitSettings()
	bridgeSettings.Get()

	ctx.bridgeSettings = bridgeSettings
	return nil
}

func (ctx *ConsumerContext) Init() error {
	var err error

	log.Infoln("Initializing DB...")
	err = ctx.initDb()
	if err != nil {
		log.Errorf("fail connect db, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize DB")

	log.Infoln("Instancing pulsar client...")
	err = ctx.initPulsarCli()
	if err != nil {
		log.Errorf("fail create pulsar client, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize pulsar client")

	log.Infoln("Initializing bridgeSettings...")
	err = ctx.initBridgeSettings()
	if err != nil {
		log.Errorf("fail init bridge settings, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize bridgeSettings")
	return nil
}

func (ctx *ConsumerContext) Close() {
	ctx.db.Close()
	ctx.pulsarCli.Close()
}
