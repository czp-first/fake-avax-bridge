package main

import (
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

func main() {
	txnString := ""
	rawTxnBytes, err := hex.DecodeString(txnString)
	if err != nil {
		log.Fatal(err)
	}

	txn := &types.Transaction{}
	rlp.DecodeBytes(rawTxnBytes, &txn)

	err = txn.UnmarshalBinary(rawTxnBytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%+v", txn)
	client, err := ethclient.Dial("https://testnet-http.dxchain.com")
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), txn)
	if err != nil {
		log.Fatalf("send err: %v", err)
	}

	log.Infof("onboard txn hash: %v", txn.Hash().String())
}
