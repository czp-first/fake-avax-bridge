package enclavecli

import (
	"context"
	"time"

	pb "github.com/czp-first/fake-avax-bridge/EnclaveProxy/enclavepb"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OnboardTxn struct {
	BlockHash      string
	TxnHash        string
	Account        string
	Amount         int64
	Identification string
	Asset          string
}

type EnclaveAPI interface {
	Close()
	ReceiveOnboardTxn(txn *pb.OnboardTxn) *pb.Status
	ReceiveOffboardTxn(txn *pb.OffboardTxn) *pb.Status
}

type EnclaveClient struct {
	enclaveConn *grpc.ClientConn
}

func NewEnclaveAPI(enclaveUrl string) (*EnclaveClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()}
	conn, err := grpc.Dial(enclaveUrl, opts...)
	if err != nil {
		return nil, err
	}
	enclave := &EnclaveClient{
		enclaveConn: conn,
	}
	return enclave, nil
}

func (c *EnclaveClient) ReceiveOnboardTxn(txn *pb.OnboardTxn) *pb.Status {

	client := pb.NewEnclaveClient(c.enclaveConn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	r, err := client.ReceiveOnboardTxn(ctx, txn)
	if err != nil {
		log.Fatalf("Fail request enclave rpc ReceiveOnboardTxn: %s\n", err)
	}
	log.Infof("Request enclave rpc ReceiveOnboardTxn resp: %v", r)
	return r
}

func (c *EnclaveClient) ReceiveOffboardTxn(txn *pb.OffboardTxn) *pb.Status {

	client := pb.NewEnclaveClient(c.enclaveConn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	r, err := client.ReceiveOffboardTxn(ctx, txn)
	if err != nil {
		log.Fatalf("Fail request enclave rpc ReceiveOffboardTxn: %s\n", err)
	}
	log.Infof("Request enclave rpc ReceiveOffboardTxn resp: %v", r)
	return r
}

func (c *EnclaveClient) Close() {
	if err := c.enclaveConn.Close(); err != nil {
		log.Warnln("Fail close enclave conn:", err)
	}
}
