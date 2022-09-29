package settings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type LocalBridgeSettingsFactory struct {
}

func (ls *LocalBridgeSettingsFactory) MakeSettings() BridgeSettingsInterface {
	return &LocalBridgeSettings{
		BridgeSettings: BridgeSettings{
			ChainToken2Name: make(map[*big.Int]map[common.Address]string),
			ChainToken2Addr: make(map[*big.Int]map[string]common.Address),
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

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/static/%s", os.Getenv("CloudServiceURL"), os.Getenv("BridgeSettingsFilename")), nil)
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

	settingsJson, err := simplejson.NewJson(body)
	if err != nil {
		log.Fatal(err)
	}
	lbs.BridgeSettings.SettingsJson = settingsJson

	var settings Settings

	err = json.Unmarshal(body, &settings)
	if err != nil {
		log.Fatal(err)
		return
	}
	lbs.BridgeSettings.Settings = settings
}

func (lbs *LocalBridgeSettings) ProduceUpdate(client pulsar.Client, isOnboard bool) {
	jsonSchema := pulsar.NewJSONSchema(middleware.SettingsSchemaDef, nil)
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:  os.Getenv("PulsarSettingsTopic"),
		Schema: jsonSchema,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	msgValue := lbs.OnboardSettingsJSON
	if !isOnboard {
		msgValue = lbs.OffboardSettingsJSON
	}
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Value: msgValue,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("publish %s message", os.Getenv("PulsarSettingsTopic"))
}

func (lbs *LocalBridgeSettings) ConsumeUpdate(body *middleware.SettingsJSON) {

	lbs.Get()

	for _, field := range body.Fields {
		switch field.Type {
		case "int":
			value := new(big.Int)
			value, ok := value.SetString(field.Value, 10)
			if !ok {
				log.Fatal("fail convert string to bigint")
			}
			lbs.BridgeSettings.SettingsJson.SetPath(field.Path, value)
		case "bool":
			if field.Value == "false" {
				lbs.BridgeSettings.SettingsJson.SetPath(field.Path, false)
			} else {
				lbs.BridgeSettings.SettingsJson.SetPath(field.Path, true)
			}
		case "string":
			lbs.BridgeSettings.SettingsJson.SetPath(field.Path, field.Value)
		}
	}

	bridgeSettingsClient := http.Client{
		Timeout: time.Second * 2,
	}

	s_bytes, err := lbs.BridgeSettings.SettingsJson.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/settings/%s", os.Getenv("CloudServiceURL"), os.Getenv("BridgeSettingsFilename")), bytes.NewBuffer(s_bytes))
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
