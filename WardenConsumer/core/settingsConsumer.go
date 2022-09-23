package core

import (
	"context"
	"os"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
)

func (ctx *ConsumerContext) ConsumeSettings() {
	jsonSchema := pulsar.NewJSONSchema(middleware.SettingsSchemaDef, nil)

	consumer, err := ctx.pulsarCli.Subscribe(pulsar.ConsumerOptions{
		Topic:                       os.Getenv("PulsarSettingsTopic"),
		SubscriptionName:            os.Getenv("PulsarSettingsSubscriptionName"),
		Schema:                      jsonSchema,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
		Type:                        pulsar.Shared,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for {
		var settings middleware.SettingsJSON
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal("could not receive msg:", err)
		}
		err = msg.GetSchemaValue(&settings)
		if err != nil {
			log.Fatal("could not get msg:", err)
		}
		log.Infof("Received message msgId: %#v -- content: '%s'", msg.ID(), string(msg.Payload()))

		ctx.bridgeSettings.ConsumeUpdate(&settings)
		consumer.Ack(msg)
	}
}
