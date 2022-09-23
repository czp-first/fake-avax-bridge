package database

import "math/big"

type WardenOnboard struct {
	RowId          int64
	ChainId        *big.Int
	BlockNumber    uint64
	Amount         *big.Int
	RealAmount     *big.Int
	Contract       string
	Account        string
	GasPrice       int64
	BlockHash      string
	TxnHash        string
	TxnIndex       uint
	Nonce          uint64
	OnboardTxnHash string
	Batch          int64
}

type WardenOffboard struct {
	RowId           int64
	ChainId         *big.Int
	BlockNumber     uint64
	Amount          *big.Int
	RealAmount      *big.Int
	Contract        string
	Account         string
	GasPrice        int64
	BlockHash       string
	TxnHash         string
	TxIndex         uint
	Nonce           uint64
	OffboardTxnHash string
	Batch           int64
}

type EnclaveOnboard struct {
	OnboardTxnHash string
	Nonce          int64
}

type EnclaveOffboard struct {
	OffboardTxnHash string
	Nonce           int64
}
