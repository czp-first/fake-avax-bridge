package core

import (
	"fmt"
	"net"
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/settings"
	"github.com/czp-first/fake-avax-bridge/WardenServer/credential"
	"github.com/czp-first/fake-avax-bridge/WardenServer/database"
	"github.com/czp-first/fake-avax-bridge/WardenServer/enclavecli"
	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type WardenContext struct {
	pb.UnimplementedWardenServer
	db              *database.PgSQL
	bridgeSettings  settings.BridgeSettingsInterface
	FromChainClient *chain.ChainClient
	ToChainClient   *chain.ChainClient
	pulsarCli       pulsar.Client
	Enclave         enclavecli.EnclaveAPI
	credential      credential.CredentialInterface
}

func NewWardenContext() (*WardenContext, error) {
	return &WardenContext{}, nil
}

func (ctx *WardenContext) Init() {
	// init db
	log.Infoln("Initializing DB...")
	ctx.initDb()
	log.Infoln("Successfully initialize DB")

	// init bridgeSettings
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

	log.Infoln("Instancing enclave proxy client...")
	ctx.initEnclaveClient()
	log.Infoln("Successfully instance enclave proxy client")

	log.Infoln("Initializing credential...")
	ctx.initCredential()
	log.Infoln("Successfully initialize credential")
}

func (ctx *WardenContext) initDb() {
	var err error
	// init db
	ctx.db, err = database.NewPgSQL()
	if err != nil {
		log.Fatalf("Fail initialize DB")
	}
}

func (ctx *WardenContext) initFromChainClient() {
	var err error
	ctx.FromChainClient, err = chain.NewChainClient(os.Getenv("FromChainHttps"), os.Getenv("FromChainWss"))
	if err != nil {
		log.Fatalf("Fail connect from chain client: %v", err)
	}
}

func (ctx *WardenContext) initToChainClient() {
	var err error
	ctx.ToChainClient, err = chain.NewChainClient(os.Getenv("ToChainHttps"), os.Getenv("ToChainWss"))
	if err != nil {
		log.Fatalf("Fail connect to chain client : %v", err)
	}
}

func (ctx *WardenContext) initBridgeSettings() {
	bridgeSettingsFactory, err := settings.GetBridgeSettingsFactory()
	if err != nil {
		log.Fatalf("Fail initialize bridge settings: %v", err)
	}
	bridgeSettings := bridgeSettingsFactory.MakeSettings()
	bridgeSettings.InitSettings()
	bridgeSettings.Get()

	ctx.bridgeSettings = bridgeSettings
}

func (ctx *WardenContext) initPulsarCli() {
	pulsarCli, err := createClient()
	if err != nil {
		log.Fatalf("fail create pulsar client: %v", err)
	}
	ctx.pulsarCli = pulsarCli
}

func (ctx *WardenContext) initEnclaveClient() {
	var err error
	ctx.Enclave, err = enclavecli.NewEnclaveAPI(os.Getenv("EnclaveRPC"))
	if err != nil {
		log.Fatalln("Fail to connect enclave server:", err)
	}
}

func (ctx *WardenContext) initCredential() {
	credentialFactory, err := credential.GetCredential()
	if err != nil {
		log.Fatalf("fail initialize credential: %v", err)
	}
	credential := credentialFactory.MakeCredential()
	ctx.credential = credential

}

func NewServer(ctx *WardenContext) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("Port")))
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
