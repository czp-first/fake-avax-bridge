package core

import (
	"fmt"
	"net"
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/chain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/czp-first/fake-avax-bridge/EnclaveProxy/enclavepb"
)

type EnclaveProxyContext struct {
	pb.UnimplementedEnclaveServer
	FromChainClient *chain.ChainClient
	ToChainClient   *chain.ChainClient
}

func NewEnclaveProxyContext() *EnclaveProxyContext {
	return &EnclaveProxyContext{}
}

func (ctx *EnclaveProxyContext) Init() error {
	var err error
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
	return nil
}

func (ctx *EnclaveProxyContext) initFromChainClient() error {
	var err error
	ctx.FromChainClient, err = chain.NewChainClient(os.Getenv("FromChainHttps"), os.Getenv("FromChainWss"))
	if err != nil {
		return err
	}
	return nil
}

func (ctx *EnclaveProxyContext) initToChainClient() error {
	var err error
	ctx.ToChainClient, err = chain.NewChainClient(os.Getenv("ToChainHttps"), os.Getenv("ToChainWss"))
	if err != nil {
		return err
	}
	return nil
}

func NewServer(ctx *EnclaveProxyContext) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("EnclaveProxyRPCPort")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterEnclaveServer(s, ctx)
	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
