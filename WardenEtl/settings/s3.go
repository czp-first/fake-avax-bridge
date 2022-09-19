package settings

import (
	"math/big"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/ethereum/go-ethereum/common"
)

type S3BridgeSettingsFactory struct {
}

func (s *S3BridgeSettingsFactory) MakeSettings() BridgeSettingsInterface {
	return &S3BridgeSettings{
		BridgeSettings: BridgeSettings{
			ChainToken2Name: make(map[*big.Int]map[common.Address]string),
			ChainToken2Addr: make(map[*big.Int]map[string]common.Address),
		},
	}
}

type S3BridgeSettings struct {
	BridgeSettings
}

func (s *S3BridgeSettings) InitSettings() {

}

func (s *S3BridgeSettings) Update(client pulsar.Client, isOnboard bool) {

}
