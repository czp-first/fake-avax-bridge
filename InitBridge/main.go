package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"

	"github.com/czp-first/fake-avax-bridge/InitBridge/proxy"
	pb "github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"
)

type Warden struct {
	Identification string `json:"identification"`
	Type           string `json:"type"`
	Credential     string `json:"credential"`
	Url            string `json:"url"`
}

type StoreReq struct {
	Method string    `json:"method"`
	Body   StoreBody `json:"body"`
}

type StoreBody struct {
	ShareVersion int       `json:"share_version"`
	Threshold    int       `json:"threshold"`
	Wardens      []*Warden `json:"wardens"`
	FromChainId  int       `json:"from_chain_id"`
	ToChainId    int       `json:"to_chain_id"`
}

type StoreResp struct {
	Content StoreContent `json:"content"`
}

type StoreContent struct {
	EncryptShares      []Share `json:"encrypt_shares"`
	FromAccountAddress string  `json:"from_account_address"`
	ToAccountAddress   string  `json:"to_account_address"`
}

type Share struct {
	Identification string `json:"identification"`
	Share          string `json:"share"`
}

var (
	envFile string
)

func init() {
	flag.StringVar(&envFile, "e", "", "env file")
}

func main() {

	var err error
	flag.Parse()
	if envFile != "" {
		log.Info("Initializing env...")
		err = godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Fail initialize env: %v", err)
			return
		}
		log.Info("Successfully initialize env")
	}

	// Set up a connection to the server.
	wardensRPC := make([]string, 0)
	err = json.Unmarshal([]byte(os.Getenv("WardensRPC")), &wardensRPC)
	if err != nil {
		log.Fatalf("fail parse wardes rpc from env file")
		return
	}

	wardens := make([]*Warden, 0)
	identification2RPC := make(map[string]string, 0)
	for _, wardenRPC := range wardensRPC {
		log.Infof("initializing warden: %s", wardenRPC)
		conn, err := grpc.Dial(wardenRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewWardenClient(conn)

		identification := uuid.NewString()
		identification2RPC[identification] = wardenRPC

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetCredential(ctx, &pb.GetCredentialReq{Identification: identification})
		if err != nil {
			log.Fatalf("could not get credential: %v", err)
		}
		log.Infof("Type: %s", r.GetType())
		log.Infof("Credential: %s", r.GetCredential())

		wardens = append(wardens, &Warden{
			Identification: identification,
			Type:           r.GetType(),
			Credential:     r.GetCredential(),
			Url:            wardenRPC,
		})
		log.Infof("successfully initializing warden: %s", wardenRPC)
	}

	log.Infof("initializing enclave")
	threshold, err := strconv.Atoi(os.Getenv("Threshold"))
	if err != nil {
		log.Fatalf("fail parse threshold from env file: %v", err)
	}

	fromChainId, err := strconv.Atoi(os.Getenv("FromChainId"))
	if err != nil {
		log.Fatalf("fail parse from chain id from env file: %v", err)
	}

	toChainId, err := strconv.Atoi(os.Getenv("ToChainId"))
	if err != nil {
		log.Fatalf("fail parse to chain id from env file: %v", err)
	}
	var resp StoreResp
	req := StoreReq{
		Method: "storeCredential",
		Body: StoreBody{
			// TODO: version
			ShareVersion: 0,
			Threshold:    threshold,
			Wardens:      wardens,
			FromChainId:  fromChainId,
			ToChainId:    toChainId,
		},
	}
	proxy.Req(&req, &resp)
	log.Infof("%+v", resp)
	log.Infof("successfully initialize enclave")

	for _, share := range resp.Content.EncryptShares {
		identification := share.Identification
		wardenRPC := identification2RPC[identification]
		log.Infof("save warden share: %v", wardenRPC)
		conn, err := grpc.Dial(wardenRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewWardenClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = c.SaveShare(ctx, &pb.SaveShareReq{
			Share:                  share.Share,
			OnboardAccountAddress:  resp.Content.FromAccountAddress,
			OffboardAccountAddress: resp.Content.ToAccountAddress,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Info("successfully save wardens share")
}
