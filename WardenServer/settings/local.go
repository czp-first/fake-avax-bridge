package settings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"

	log "github.com/sirupsen/logrus"
)

type LocalBridgeSettingsFactory struct {
}

func (ls *LocalBridgeSettingsFactory) MakeSettings() BridgeSettingsInterface {
	return &LocalBridgeSettings{
		BridgeSettings: BridgeSettings{
			ChainToken2Name: make(map[uint64]map[common.Address]string),
			ChainToken2Addr: make(map[uint64]map[string]common.Address),
		},
	}
}

type LocalBridgeSettings struct {
	BridgeSettings
}

func (lbs *LocalBridgeSettings) InitSettings() {

	settingsClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("BridgeSettingsFileURL"), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := settingsClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var settings Settings

	err = json.Unmarshal(body, &settings)
	if err != nil {
		log.Fatal(err)
		return
	}
	lbs.BridgeSettings.Settings = settings
}

func (lbs *LocalBridgeSettings) Update() {

}
