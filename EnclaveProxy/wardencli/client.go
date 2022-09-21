package wardencli

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"
)

func GetOnboardTxn(address string, in *pb.GetWardenOnboardReq) *pb.GetWardenOnboardResp {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
