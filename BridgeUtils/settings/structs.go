package settings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// bridge setting file
type Settings struct {
	Critical    Critical    `json:"critical"`
	NonCritical NonCritical `json:"nonCritical"`
}

type Critical struct {
	Assets                      map[string]Asset
	Networks                    Networks
	WalletAddress               WalletAddress
	UseEip1559TransactionFormat bool `json:"useEip1559TransactionFormat"`
}

type Asset struct {
	Asset                  string
	AssetName              string
	ChainlinkFeedAddress   common.Address
	Denomination           int64
	NativeContractAddress  common.Address
	NativeNetwork          string
	OffboardFeeDollars     int64
	OnboardFeeDollars      int64
	TokenName              string
	WrappedContractAddress common.Address
	WrappedNetwork         string
}

type Networks struct {
	Dxchain  *big.Int
	Ethereum *big.Int
}

type WalletAddress struct {
	Dxchain  common.Address
	Ethereum common.Address
}

type NonCritical struct {
	ChainlinkDxUsdFeedAddress  common.Address `json:"chainlinkDxUsdFeedAddress"`
	ChainlinkEthUsdFeedAddress common.Address `json:"chainlinkEthUsdFeedAddress"`
	CurrentDxPrice             *big.Int       `json:"currentDxPrice"`
	CurrentEthPrice            *big.Int       `json:"currentEthPrice"`
	CurrentGasPrices           CurrentGasPrices
	MinimumConfirmations       MinimumConfirmations
	NetworkViews               NetworkViews
	ShareVersion               *big.Int
}

type CurrentGasPrices struct {
	Dxchain  GasPrice
	Ethereum GasPrice
}

type GasPrice struct {
	NextBaseFee  string
	SuggestedTip string
}

type MinimumConfirmations struct {
	Dxchain  uint64
	Ethereum uint64
}

type NetworkViews struct {
	Dxchain  NetworkView
	Ethereum NetworkView
}

type NetworkView struct {
	LastIndexedBlock uint64
	LastSeenBlock    uint64
	NodeVersion      string
}

// runtime
type AssetConfig struct {
	ChainlinkFeedAddress common.Address
	FeeDollars           int64
	Name                 string
	AnotherChainContract common.Address
}

type AssetsConfig struct {
	Assets                     map[common.Address]AssetConfig
	ChainlinkDxUsdFeedAddress  common.Address
	ChainlinkEthUsdFeedAddress common.Address
	CurrentEthPrice            *big.Int
	CurrentDxPrice             *big.Int
}
