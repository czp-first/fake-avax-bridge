package settings

import (
	"fmt"
	"math/big"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/ethereum/go-ethereum/common"
)

type BridgeSettingsFactory interface {
	MakeSettings() BridgeSettingsInterface
}

type BridgeSettingsInterface interface {
	InitSettings()
	Get() error

	ProduceUpdate(client pulsar.Client, isOnboard bool)
	ConsumeUpdate(body *middleware.SettingsJSON)

	GetSettings() Settings
	GetTokenAddrs(chainId *big.Int) []common.Address
	GetOnboardAssetsConfig() *AssetsConfig
	IsUseEip1559TransactionFormat() bool
	GetCurrentGasPrices() CurrentGasPrices

	NewOnboardSettingsJSON()
	AppendOnboardSettingsJSON(field *middleware.SettingsField)
	NewOffboardSettingsJSON()
	AppendOffboardSettingsJSON(field *middleware.SettingsField)
}

func GetBridgeSettingsFactory() (BridgeSettingsFactory, error) {
	bridgeSettingsType := os.Getenv("BridgeSettingsType")
	if bridgeSettingsType == "local" {
		return &LocalBridgeSettingsFactory{}, nil
	}
	if bridgeSettingsType == "s3" {
		return &S3BridgeSettingsFactory{}, nil
	}
	return nil, fmt.Errorf("unknown bridge settings type: %s", bridgeSettingsType)

}

type BridgeSettings struct {
	SettingsJson         *simplejson.Json
	Settings             Settings
	ChainToken2Name      map[*big.Int]map[common.Address]string
	ChainToken2Addr      map[*big.Int]map[string]common.Address
	OnboardAssetsConfig  AssetsConfig
	OffboardAssetsConfig AssetsConfig
	OnboardSettingsJSON  *middleware.SettingsJSON
	OffboardSettingsJSON *middleware.SettingsJSON
}

func (bs *BridgeSettings) NewOnboardSettingsJSON() {
	bs.OnboardSettingsJSON = &middleware.SettingsJSON{Fields: make([]*middleware.SettingsField, 0)}
}

func (bs *BridgeSettings) AppendOnboardSettingsJSON(field *middleware.SettingsField) {
	bs.OnboardSettingsJSON.Fields = append(bs.OnboardSettingsJSON.Fields, field)
}

func (bs *BridgeSettings) NewOffboardSettingsJSON() {
	bs.OffboardSettingsJSON = &middleware.SettingsJSON{Fields: make([]*middleware.SettingsField, 0)}
}

func (bs *BridgeSettings) AppendOffboardSettingsJSON(field *middleware.SettingsField) {
	bs.OffboardSettingsJSON.Fields = append(bs.OffboardSettingsJSON.Fields, field)
}

func (bs *BridgeSettings) Get() error {
	var settings = bs.Settings

	onboardAssetsConfig := AssetsConfig{
		Assets:                     make(map[common.Address]AssetConfig),
		ChainlinkDxUsdFeedAddress:  bs.Settings.NonCritical.ChainlinkDxUsdFeedAddress,
		ChainlinkEthUsdFeedAddress: bs.Settings.NonCritical.ChainlinkEthUsdFeedAddress,
		CurrentDxPrice:             bs.Settings.NonCritical.CurrentDxPrice,
		CurrentEthPrice:            bs.Settings.NonCritical.CurrentEthPrice,
	}

	offboardAssetsConfig := AssetsConfig{
		Assets:                     make(map[common.Address]AssetConfig),
		ChainlinkDxUsdFeedAddress:  bs.Settings.NonCritical.ChainlinkDxUsdFeedAddress,
		ChainlinkEthUsdFeedAddress: bs.Settings.NonCritical.ChainlinkEthUsdFeedAddress,
		CurrentDxPrice:             bs.Settings.NonCritical.CurrentDxPrice,
		CurrentEthPrice:            bs.Settings.NonCritical.CurrentEthPrice,
	}

	for assetName, asset := range settings.Critical.Assets {
		onboardAssetConfig := AssetConfig{
			Name:                 assetName,
			ChainlinkFeedAddress: asset.ChainlinkFeedAddress,
			FeeDollars:           asset.OnboardFeeDollars,
			AnotherChainContract: asset.WrappedContractAddress,
		}
		onboardAssetsConfig.Assets[asset.NativeContractAddress] = onboardAssetConfig

		offboardAssetConfig := AssetConfig{
			Name:                 assetName,
			ChainlinkFeedAddress: asset.ChainlinkFeedAddress,
			FeeDollars:           asset.OffboardFeeDollars,
			AnotherChainContract: asset.NativeContractAddress,
		}
		offboardAssetsConfig.Assets[asset.WrappedContractAddress] = offboardAssetConfig

		nativeChainId := bs.Settings.Critical.Networks.Ethereum
		nativeSubChainTokenNameMap, isFound := bs.ChainToken2Name[nativeChainId]
		if isFound {
			nativeSubChainTokenNameMap[asset.NativeContractAddress] = assetName
		} else {
			bs.ChainToken2Name[nativeChainId] = make(map[common.Address]string)
			bs.ChainToken2Name[nativeChainId][asset.NativeContractAddress] = assetName
		}

		wrappedChainId := bs.Settings.Critical.Networks.Dxchain
		wrappedSubChainTokenNameMap, isFound := bs.ChainToken2Name[wrappedChainId]
		if isFound {
			wrappedSubChainTokenNameMap[asset.WrappedContractAddress] = assetName
		} else {
			bs.ChainToken2Name[wrappedChainId] = make(map[common.Address]string)
			bs.ChainToken2Name[wrappedChainId][asset.WrappedContractAddress] = assetName
		}

		nativeSubChainTokenAddrMap, isFound := bs.ChainToken2Addr[nativeChainId]
		if !isFound {
			nativeSubChainTokenAddrMap = make(map[string]common.Address)
			bs.ChainToken2Addr[nativeChainId] = nativeSubChainTokenAddrMap
		}
		nativeSubChainTokenAddrMap[assetName] = asset.NativeContractAddress

		wrappedSubChainTokenAddrMap, isFound := bs.ChainToken2Addr[wrappedChainId]
		if !isFound {
			wrappedSubChainTokenAddrMap = make(map[string]common.Address)
			bs.ChainToken2Addr[wrappedChainId] = wrappedSubChainTokenAddrMap
		}
		wrappedSubChainTokenAddrMap[assetName] = asset.WrappedContractAddress
	}

	bs.Settings = settings
	bs.OnboardAssetsConfig = onboardAssetsConfig
	bs.OffboardAssetsConfig = offboardAssetsConfig
	return nil
}

func (bs *BridgeSettings) GetSettings() Settings {
	return bs.Settings
}

func (bs *BridgeSettings) GetTokenAddrs(chainId *big.Int) []common.Address {
	tokenMap := bs.ChainToken2Name[chainId]
	tokensAddr := make([]common.Address, 0, len(tokenMap))
	for tokenAddr := range tokenMap {
		tokensAddr = append(tokensAddr, tokenAddr)
	}
	return tokensAddr
}

func (bs *BridgeSettings) GetOnboardAssetsConfig() *AssetsConfig {
	return &bs.OnboardAssetsConfig
}

func (bs *BridgeSettings) IsUseEip1559TransactionFormat() bool {
	return bs.Settings.Critical.UseEip1559TransactionFormat
}

func (bs *BridgeSettings) GetCurrentGasPrices() CurrentGasPrices {
	return bs.Settings.NonCritical.CurrentGasPrices
}
