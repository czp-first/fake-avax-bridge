package core

import (
	"fmt"
	"net"
	"os"

	"github.com/czp-first/fake-avax-bridge/WardenServer/chains"
	"github.com/czp-first/fake-avax-bridge/WardenServer/credential"
	"github.com/czp-first/fake-avax-bridge/WardenServer/database"

	"github.com/czp-first/fake-avax-bridge/WardenServer/enclavecli"
	"github.com/czp-first/fake-avax-bridge/WardenServer/settings"
	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type WardenContext struct {
	pb.UnimplementedWardenServer
	db             *database.PgSQL
	bridgeSettings settings.BridgeSettingsInterface
	EthClient      *chains.EthClient
	DxClient       *chains.DxClient
	pulsarCli      pulsar.Client
	Enclave        enclavecli.EnclaveAPI
	credential     credential.CredentialInterface
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

	log.Infoln("Connecting eth client...")
	ctx.initEthClient()
	log.Infoln("Successfully connect eth client")

	log.Infoln("Connecting dx client...")
	ctx.initDxClient()
	log.Infoln("Successfully connect dx client")

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

func (ctx *WardenContext) initEthClient() {
	var err error
	ctx.EthClient, err = chains.NewEthClient()
	if err != nil {
		log.Fatalf("Fail connect eth client: %v\n", err)
	}
}

func (ctx *WardenContext) initDxClient() {
	var err error
	ctx.DxClient, err = chains.NewDxClient()
	if err != nil {
		log.Fatalf("Fail connect dx client : %v\n", err)
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
