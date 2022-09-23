package database

import "math/big"

type OnboardTxn struct {
	RowId          int64
	ChainId        int64
	BlockNumber    int64
	Amount         *big.Int
	RealAmount     *big.Int
	Contract       string
	Account        string
	GasPrice       int64
	BlockHash      string
	TxnHash        string
	TxIndex        int
	Nonce          int64
	OnboardTxnHash string
	Batch          int64
}

type OffboardTxn struct {
	RowId           int64
	ChainId         int64
	BlockNumber     int64
	Amount          int64
	RealAmount      int64
	Contract        string
	Account         string
	GasPrice        int64
	BlockHash       string
	TxnHash         string
	TxIndex         int
	Nonce           int64
	OffboardTxnHash string
	Batch           int64
}

type EnclaveOnboard struct {
	OnboardTxnHash string
	Nonce          uint64
}

type EnclaveOffboard struct {
	OffboardTxnHash string
	Nonce           uint64
}
