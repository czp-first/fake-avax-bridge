package settings

import (
	"github.com/ethereum/go-ethereum/common"
)

type S3BridgeSettingsFactory struct {
}

func (s *S3BridgeSettingsFactory) MakeSettings() BridgeSettingsInterface {
	return &S3BridgeSettings{
		BridgeSettings: BridgeSettings{
			ChainToken2Name: make(map[uint64]map[common.Address]string),
			ChainToken2Addr: make(map[uint64]map[string]common.Address),
		},
	}
}

type S3BridgeSettings struct {
	BridgeSettings
}

func (s *S3BridgeSettings) InitSettings() {

}

func (s *S3BridgeSettings) Update() {

}
