package settings

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/bitly/go-simplejson"

	"WardenConsumer/middleware"
)

type LocalBridgeSettingsFactory struct {
}

func (lbsf *LocalBridgeSettingsFactory) MakeBridgeSettings() BridgeSettingsInterface {
	return &LocalBridgeSettings{}
}

type LocalBridgeSettings struct {
	BridgeSettings
}

func (lbs *LocalBridgeSettings) Get() {
	bridgeSettingsClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("BridgeSettingsFileURL"), nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := bridgeSettingsClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bridgeSettings, err := simplejson.NewJson(body)
	if err != nil {
		log.Fatal(err)
	}

	lbs.BridgeSettings.Settings = bridgeSettings
}

func (lbs *LocalBridgeSettings) Update(body *middleware.SettingsJSON) {

	lbs.Get()

	for _, field := range body.Fields {
		switch field.Type {
		case "int":
			value := new(big.Int)
			value, ok := value.SetString(field.Value, 10)
			if !ok {
				log.Fatal("fail convert string to bigint")
			}
			lbs.BridgeSettings.Settings.SetPath(field.Path, value)
		case "bool":
			if field.Value == "false" {
				lbs.BridgeSettings.Settings.SetPath(field.Path, false)
			} else {
				lbs.BridgeSettings.Settings.SetPath(field.Path, true)
			}
		case "string":
			lbs.BridgeSettings.Settings.SetPath(field.Path, field.Value)
		}
	}

	bridgeSettingsClient := http.Client{
		Timeout: time.Second * 2,
	}

	s_bytes, err := lbs.BridgeSettings.Settings.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest(http.MethodPut, os.Getenv("BridgeSettingsFileURL"), bytes.NewBuffer(s_bytes))
	if err != nil {
		log.Fatal(err)
	}
	response, err := bridgeSettingsClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}
}
