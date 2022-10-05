package wardencli

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"
)

type WardenClient struct {
	wardenConn *grpc.ClientConn
}

func NewWardenAPI(wardenUrl string) (*WardenClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock()}
	conn, err := grpc.Dial(wardenUrl, opts...)
	if err != nil {
		return nil, err
	}
	warden := &WardenClient{
		wardenConn: conn,
	}
	return warden, nil
}

func GetOnboardTxn(address string, in *pb.GetWardenOnboardReq) *pb.GetWardenOnboardResp {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock()}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	log.Infof("connected %s", address)
	defer conn.Close()

	wardenClient := pb.NewWardenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := wardenClient.GetWardenOnboard(ctx, in)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return r
}

func Onboard(address string, in *pb.OnboardReq) *pb.Empty {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock()}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	log.Infof("connected %s", address)
	defer conn.Close()

	wardenClient := pb.NewWardenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := wardenClient.Onboard(ctx, in)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return r
}

func GetOffboardTxn(address string, in *pb.GetWardenOffboardReq) *pb.GetWardenOffboardResp {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock()}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	log.Infof("connected %s", address)
	defer conn.Close()

	wardenClient := pb.NewWardenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := wardenClient.GetWardenOffboard(ctx, in)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return r
}

func Offboard(address string, in *pb.OffboardReq) *pb.Empty {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock()}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	log.Infof("connected %s", address)
	defer conn.Close()

	wardenClient := pb.NewWardenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := wardenClient.Offboard(ctx, in)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return r
}
