package core

import (
	"fmt"
	"net"
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/sqldb"
	"github.com/czp-first/fake-avax-bridge/WardenServer/credential"
	"github.com/czp-first/fake-avax-bridge/WardenServer/dal"
	"github.com/czp-first/fake-avax-bridge/WardenServer/enclavecli"
	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type WardenContext struct {
	pb.UnimplementedWardenServer
	db              *dal.DAL
	bridgeSettings  settings.BridgeSettingsInterface
	FromChainClient *chain.ChainClient
	ToChainClient   *chain.ChainClient
	pulsarCli       pulsar.Client
	Enclave         enclavecli.EnclaveAPI
	credential      credential.CredentialInterface
}

func NewWardenContext() *WardenContext {
	return &WardenContext{}
}

func (ctx *WardenContext) Init() error {
	// init db
	log.Infoln("Initializing DB...")
	err := ctx.initDb()
	if err != nil {
		log.Errorf("fail connect db, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize DB")

	// init bridgeSettings
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

	log.Infoln("Instancing enclave proxy client...")
	err = ctx.initEnclaveClient()
	if err != nil {
		log.Errorf("fail create enclave client, err:%v", err)
		return err
	}
	log.Infoln("Successfully instance enclave proxy client")

	log.Infoln("Initializing credential...")
	err = ctx.initCredential()
	if err != nil {
		log.Errorf("fail load credential, err:%v", err)
		return err
	}
	log.Infoln("Successfully initialize credential")
	return nil
}

func (ctx *WardenContext) initDb() error {
	var err error
	// init db
	ctx.db, err = dal.NewDAL(sqldb.PostgresDriver, fmt.Sprintf(sqldb.PostgresFmt, os.Getenv("PgUser"), os.Getenv("PgPassword"), fmt.Sprintf("%s:%s", os.Getenv("PgHost"), os.Getenv("PgPort")), os.Getenv("PgDb")), sqldb.DefaultPostgresPoolSize)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *WardenContext) initFromChainClient() error {
	var err error
	ctx.FromChainClient, err = chain.NewChainClient(os.Getenv("FromChainHttps"), os.Getenv("FromChainWss"))
	if err != nil {
		return err
	}
	return nil
}

func (ctx *WardenContext) initToChainClient() error {
	var err error
	ctx.ToChainClient, err = chain.NewChainClient(os.Getenv("ToChainHttps"), os.Getenv("ToChainWss"))
	if err != nil {
		return err
	}
	return nil
}

func (ctx *WardenContext) initBridgeSettings() error {
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

func (ctx *WardenContext) initPulsarCli() error {
	pulsarCli, err := middleware.CreatePulsarClient(os.Getenv("PulsarURL"))
	if err != nil {
		return err
	}
	ctx.pulsarCli = pulsarCli
	return nil
}

func (ctx *WardenContext) initEnclaveClient() error {
	var err error
	ctx.Enclave, err = enclavecli.NewEnclaveAPI(os.Getenv("EnclaveProxyRPC"))
	if err != nil {
		return err
	}
	return nil
}

func (ctx *WardenContext) initCredential() error {
	credentialFactory, err := credential.GetCredential()
	if err != nil {
		return err
	}
	credential := credentialFactory.MakeCredential()
	ctx.credential = credential
	return nil
}

func NewServer(ctx *WardenContext) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("WardenRPCPort")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	log.Infof("Warden RPC listening at %v", lis.Addr())
	pb.RegisterWardenServer(s, ctx)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
