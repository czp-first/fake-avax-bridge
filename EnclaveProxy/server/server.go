package server

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/czp-first/fake-avax-bridge/EnclaveProxy/enclavepb"
	"github.com/czp-first/fake-avax-bridge/EnclaveProxy/proxy"
	"github.com/czp-first/fake-avax-bridge/EnclaveProxy/wardencli"

	"github.com/czp-first/fake-avax-bridge/WardenServer/wardenpb"
)

type server struct {
	pb.UnimplementedEnclaveServer
}

type EnclaveTxn struct {
	BlockHash string `json:"block_hash"`
	TxnHash   string `json:"txn_hash"`
	Batch     int64  `json:"batch"`
}

type EnclaveOnboardTxnBody struct {
	EnclaveTxn     EnclaveTxn `json:"txn"`
	Identification string     `json:"identification"`
}

type EnclaveOnboardTxnReq struct {
	Method string                `json:"method"`
	Body   EnclaveOnboardTxnBody `json:"body"`
}

type EnclaveOnboardTxnContent struct {
	Status  string   `json:"status"`
	Wardens []string `json:"wardens"`
}

type EnclaveOnboardTxnResp struct {
	Content EnclaveOnboardTxnContent `json:"content"`
}

type EnclaveSignTxnReq struct {
	Method string             `json:"method"`
	Body   EnclaveSignTxnBody `json:"body"`
}

type EnclaveSignTxnBody struct {
	IsEip1559    bool           `json:"is_eip1559"`
	WardenShares []*WardenShare `json:"warden_shares"`
	ChainId      uint64         `json:"chain_id"`
	ContractAddr string         `json:"contract_addr"`
	RealAmount   uint64         `json:"amount"`
	GasPrice     uint64         `json:"gas_price"`
	AccountAddr  string         `json:"account_addr"`
	Nonce        uint64         `json:"nonce"`
	OriginTxn    string         `json:"origin_txn"`
	Fee          uint64         `json:"fee"`
}

type WardenShare struct {
	Identification string `json:"identification"`
	EncryptShare   string `json:"encrypt_share"`
}

type EnclaveSignTxnResp struct {
	Content EnclaveSignTxnContent `json:"content"`
}

type EnclaveSignTxnContent struct {
	Txn       string `json:"txn"`
	Nonce     uint64 `json:"nonce"`
	GasPrice  uint64 `json:"gas_price"`
	IsEip1559 bool   `json:"is_eip1559"`
}

type EnclaveOffboardTxnReq struct {
	Method string                 `json:"method"`
	Body   EnclaveOffboardTxnBody `json:"body"`
}

type EnclaveOffboardTxnBody struct {
	EnclaveTxn     EnclaveTxn `json:"txn"`
	Identification string     `json:"identification"`
}

type EnclaveOffboardTxnResp struct {
	Content EnclaveOffboardTxnContent `json:"content"`
}

type EnclaveOffboardTxnContent struct {
	Status  string   `json:"status"`
	Wardens []string `json:"wardens"`
}

type EnclaveSignOffboardTxnReq struct {
	Method string                     `json:"method"`
	Body   EnclaveSignOffboardTxnBody `json:"body"`
}

type EnclaveSignOffboardTxnBody struct {
	IsEip1559    bool           `json:"is_eip1559"`
	WardenShares []*WardenShare `json:"warden_shares"`
	ChainId      uint64         `json:"chain_id"`
	ContractAddr string         `json:"contract_addr"`
	RealAmount   uint64         `json:"amount"`
	GasPrice     uint64         `json:"gas_price"`
	AccountAddr  string         `json:"account_addr"`
	Nonce        uint64         `json:"nonce"`
}

type EnclaveSignOffboardTxnResp struct {
	Content EnclaveSignOffboardTxnContent `json:"content"`
}

type EnclaveSignOffboardTxnContent struct {
	Txn       string `json:"txn"`
	Nonce     uint64 `json:"nonce"`
	GasPrice  uint64 `json:"gas_price"`
	IsEip1559 bool   `json:"is_eip1559"`
}

func (s *server) ReceiveOnboardTxn(ctx context.Context, in *pb.OnboardTxn) (*pb.Status, error) {

	log.Infof("receive onboard txn from warden, blockHash[%v], txnHash[%v]\n", in.BlockHash, in.TxnHash)
	req := EnclaveOnboardTxnReq{
		Method: "onboardTxn",
		Body: EnclaveOnboardTxnBody{
			EnclaveTxn: EnclaveTxn{
				BlockHash: in.BlockHash,
				TxnHash:   in.TxnHash,
				Batch:     in.Batch,
			},
			Identification: in.Identification,
		},
	}

	var resp EnclaveOnboardTxnResp
	proxy.Req(&req, &resp)
	log.Infof("onboard txn status: %v\n", resp.Content.Status)

	if resp.Content.Status != "ready" {
		return &pb.Status{Status: resp.Content.Status}, nil
	}

	go readyOnboard(resp.Content.Wardens, in.BlockHash, in.TxnHash, in.Batch)
	return &pb.Status{Status: resp.Content.Status}, nil
}

func (s *server) ReceiveOffboardTxn(ctx context.Context, in *pb.OffboardTxn) (*pb.Status, error) {

	req := EnclaveOffboardTxnReq{
		Method: "offboardTxn",
		Body: EnclaveOffboardTxnBody{
			EnclaveTxn: EnclaveTxn{
				BlockHash: in.BlockHash,
				TxnHash:   in.TxnHash,
				Batch:     in.Batch,
			},
			Identification: in.Identification,
		},
	}

	var resp EnclaveOffboardTxnResp
	proxy.Req(&req, &resp)

	if resp.Content.Status != "ready" {
		return &pb.Status{Status: resp.Content.Status}, nil
	}

	go readyOffboard(resp.Content.Wardens, in.BlockHash, in.TxnHash, in.Batch)
	return &pb.Status{Status: resp.Content.Status}, nil
}

func readyOnboard(wardens []string, blockHash, txnHash string, batch int64) error {

	log.Infof("ready onboard txn: %s", txnHash)
	wardenConfFile, _ := os.OpenFile(os.Getenv("WardensConfPath"), os.O_RDONLY, 0644)
	defer wardenConfFile.Close()
	wardenMap := make(map[string]string)
	decoder := json.NewDecoder(wardenConfFile)

	err := decoder.Decode(&wardenMap)
	if err != nil {
		return err
	}

	in := wardenpb.GetWardenOnboardReq{
		BlockHash: blockHash,
		TxnHash:   txnHash,
	}

	var chainId, fee, gasPrice, nonce, realAmount uint64
	var account, contract string
	var isEip1559 bool
	wardenShares := make([]*WardenShare, 0)

	for index, identification := range wardens {

		onboardTxnResp := wardencli.GetOnboardTxn(wardenMap[identification], &in)

		// TODO: nonce 处理
		if index == 0 {
			chainId = onboardTxnResp.ChainId
			realAmount = onboardTxnResp.RealAmount
			account = onboardTxnResp.Account
			gasPrice = onboardTxnResp.GasPrice
			contract = onboardTxnResp.Contract
			nonce = onboardTxnResp.Nonce
			fee = onboardTxnResp.Fee
			isEip1559 = onboardTxnResp.IsEip1559
		} else {

			if isEip1559 != onboardTxnResp.IsEip1559 {
				log.Fatalf("isEip1559 diff: %v, %v", isEip1559, onboardTxnResp.IsEip1559)
			}

			if chainId != onboardTxnResp.ChainId {
				log.Fatalf("chainId diff: %v, %v", chainId, onboardTxnResp.ChainId)
			}

			if realAmount < onboardTxnResp.RealAmount {
				realAmount = onboardTxnResp.RealAmount
			}

			if account != onboardTxnResp.Account {
				log.Fatalf("account diff: %v, %v", account, onboardTxnResp.Account)
			}

			if gasPrice > onboardTxnResp.GasPrice {
				gasPrice = onboardTxnResp.GasPrice
			}

			if contract != onboardTxnResp.Contract {
				log.Fatalf("contract diff: %v, %v", contract, onboardTxnResp.Contract)
			}

			if fee > onboardTxnResp.Fee {
				fee = onboardTxnResp.Fee
			}
		}

		wardenShares = append(wardenShares, &WardenShare{
			Identification: identification,
			EncryptShare:   onboardTxnResp.Share,
		})

	}

	req := EnclaveSignTxnReq{
		Method: "signOnboardTxn",
		Body: EnclaveSignTxnBody{
			IsEip1559:    isEip1559,
			WardenShares: wardenShares,
			ChainId:      chainId,
			ContractAddr: contract,
			RealAmount:   realAmount,
			GasPrice:     gasPrice,
			AccountAddr:  account,
			Nonce:        nonce,
			OriginTxn:    txnHash,
			Fee:          fee,
		},
	}

	var resp EnclaveSignTxnResp
	proxy.Req(&req, &resp)

	rawTxnBytes, err := hex.DecodeString(resp.Content.Txn)
	if err != nil {
		log.Fatal(err)
	}
	txn := &types.Transaction{}
	rlp.DecodeBytes(rawTxnBytes, &txn)

	err = txn.UnmarshalBinary(rawTxnBytes)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(os.Getenv("DxChainHttps"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), txn)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("onboard txn hash: %v", txn.Hash().String())

	for _, addr := range wardenMap {
		wardencli.Onboard(addr, &wardenpb.OnboardReq{
			BlockHash:      blockHash,
			TxnHash:        txnHash,
			OnboardTxnHash: txn.Hash().String(),
			Nonce:          resp.Content.Nonce,
			Batch:          batch,
		})
	}

	return nil
}

func readyOffboard(wardens []string, blockHash, txnHash string, batch int64) error {
	wardenConfFile, _ := os.OpenFile(os.Getenv("WardensConfPath"), os.O_RDONLY, 0644)
	defer wardenConfFile.Close()
	wardenMap := make(map[string]string)
	decoder := json.NewDecoder(wardenConfFile)

	err := decoder.Decode(&wardenMap)
	if err != nil {
		return err
	}

	in := wardenpb.GetWardenOffboardReq{
		BlockHash: blockHash,
		TxnHash:   txnHash,
	}

	var chainId, realAmount, gasPrice, nonce uint64
	var account, contract string
	wardenShares := make([]*WardenShare, 0)

	for index, identification := range wardens {

		offboardTxnResp := wardencli.GetOffboardTxn(wardenMap[identification], &in)

		// TODO: nonce 处理
		if index == 0 {
			chainId = offboardTxnResp.ChainId
			realAmount = offboardTxnResp.RealAmount
			account = offboardTxnResp.Account
			gasPrice = offboardTxnResp.GasPrice
			contract = offboardTxnResp.Contract
			nonce = offboardTxnResp.Nonce
		} else {
			if chainId != offboardTxnResp.ChainId {
				log.Panicf("chainId diff: %v, %v", chainId, offboardTxnResp.ChainId)
				return nil
			}

			if realAmount < offboardTxnResp.RealAmount {
				realAmount = offboardTxnResp.RealAmount
			}

			if account != offboardTxnResp.Account {
				log.Panicf("account diff: %v, %v", account, offboardTxnResp.Account)
				return nil
			}

			if gasPrice > offboardTxnResp.GasPrice {
				gasPrice = offboardTxnResp.GasPrice
			}

			if contract != offboardTxnResp.Contract {
				log.Panicf("contract diff: %v, %v", contract, offboardTxnResp.Contract)
				return nil
			}
		}

		wardenShares = append(wardenShares, &WardenShare{
			Identification: identification,
			EncryptShare:   offboardTxnResp.Share,
		})
	}

	req := EnclaveSignOffboardTxnReq{
		Method: "signOffboardTxn",
		Body: EnclaveSignOffboardTxnBody{
			WardenShares: wardenShares,
			ChainId:      chainId,
			ContractAddr: contract,
			RealAmount:   realAmount,
			GasPrice:     gasPrice,
			AccountAddr:  account,
			Nonce:        nonce,
		},
	}

	var resp EnclaveSignOffboardTxnResp
	proxy.Req(&req, &resp)

	rawTxnBytes, err := hex.DecodeString(resp.Content.Txn)
	if err != nil {
		log.Fatal(err)
	}
	txn := new(types.Transaction)
	rlp.DecodeBytes(rawTxnBytes, &txn)

	client, err := ethclient.Dial(os.Getenv("ETHHttps"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), txn)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("offboard txn hash: %v", txn.Hash().String())

	for _, addr := range wardenMap {
		wardencli.Offboard(addr, &wardenpb.OffboardReq{
			BlockHash:       blockHash,
			TxnHash:         txnHash,
			OffboardTxnHash: txn.Hash().String(),
			Nonce:           resp.Content.Nonce,
			Batch:           batch,
		})
	}

	return nil
}

func NewServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("RPCPort")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterEnclaveServer(s, &server{})
	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
