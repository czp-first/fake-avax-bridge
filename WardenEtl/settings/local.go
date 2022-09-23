package settings

import (
	"WardenEtl/middleware"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"time"

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

func (lbs *LocalBridgeSettings) Update(client pulsar.Client, isOnboard bool) {
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
