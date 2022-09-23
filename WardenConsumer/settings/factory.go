package settings

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"

	"WardenConsumer/middleware"
)

type BridgeSettingsFactory interface {
	MakeBridgeSettings() BridgeSettingsInterface
}

type BridgeSettingsInterface interface {
	Get()
	Update(body *middleware.SettingsJSON)
}

func GetBridgeSettingsFactory() (BridgeSettingsFactory, error) {
	bridgeSettingsType := os.Getenv("BridgeSettingsType")
	if bridgeSettingsType == "local" {
		return &LocalBridgeSettingsFactory{}, nil
	}

	return nil, fmt.Errorf("unknown BridgeSettingsType: %s", bridgeSettingsType)
}

type BridgeSettings struct {
	Settings *simplejson.Json
}
