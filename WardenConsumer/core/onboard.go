package core

import (
	"context"
	"os"

	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/middleware"
)

func (ctx *ConsumerContext) ConsumeOnboardTxn() {
	jsonSchema := pulsar.NewJSONSchema(middleware.OnboardTxnSchemaDef, nil)

	consumer, err := ctx.pulsarCli.Subscribe(pulsar.ConsumerOptions{
		Topic:                       os.Getenv("PulsarOnboardTopic"),
		SubscriptionName:            os.Getenv("PulsarOnboardSubscriptionName"),
		Schema:                      jsonSchema,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
		Type:                        pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for {
		var onboardTxn middleware.OnboardTxnJSON
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal("could not receive msg:", err)
		}
		err = msg.GetSchemaValue(&onboardTxn)
		if err != nil {
			log.Fatal("could not get msg:", err)
		}

		log.Infof("Received message msgId: %#v -- content: '%s'", msg.ID(), string(msg.Payload()))

		handlerTx, err := ctx.db.GetDB().Begin()
		if err != nil {
			log.Fatalf("DB error: %v", err)
		}

		switch onboardTxn.Type {
		// receive onboard txn from enclave
		case middleware.Enclave:
			err = ctx.db.EnclaveOnboard(
				onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.OnboardTxnHash, onboardTxn.Nonce, onboardTxn.Batch, handlerTx,
			)
			if err != nil {
				handlerTx.Rollback()
				log.Fatalf("DB error: %v", err)
			}
			err = handlerTx.Commit()
			if err != nil {
				handlerTx.Rollback()
				log.Fatalf("DB error: %v", err)
			}
			consumer.Ack(msg)
			continue
		// receive onboard txn from warden
		case middleware.Warden:
			isExist, err := ctx.db.IsWardenOnboardExist(onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.Batch)
			if err != nil {
				log.Fatalf("DB error: %v", err)
			}
			if isExist {
				log.Infoln("this onboard txn from warden already exists in warden")
				consumer.Ack(msg)
				continue
			}

			enclaveOnboard, err := ctx.db.SelectEnclaveOnboard(onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.Batch)
			if err != nil {
				log.Fatalf("DB error: %v", err)
			}
			if enclaveOnboard != nil {
				err = ctx.db.InsertPendingWardenOnboard(
					onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.ContractAddress, onboardTxn.AccountAddress,
					enclaveOnboard.OnboardTxnHash, onboardTxn.BlockNumber, enclaveOnboard.Nonce, onboardTxn.Batch,
					onboardTxn.TxnIndex, onboardTxn.ChainId, onboardTxn.Amount,
					handlerTx,
				)
				if err != nil {
					handlerTx.Rollback()
					log.Fatalf("DB error: %v", err)
				}
				err = handlerTx.Commit()
				if err != nil {
					handlerTx.Rollback()
					log.Fatalf("DB error: %v", err)
				}
				consumer.Ack(msg)
				continue
			}

			err = ctx.db.InsertWardenOnboard(
				onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.ContractAddress, onboardTxn.AccountAddress,
				onboardTxn.ChainId, onboardTxn.Amount, onboardTxn.BlockNumber, onboardTxn.Batch, onboardTxn.TxnIndex,
				handlerTx,
			)
			if err != nil {
				handlerTx.Rollback()
				log.Fatal("DB error: %v", err)
			}

			err = handlerTx.Commit()
			if err != nil {
				handlerTx.Rollback()
				log.Fatalf("DB error: %v", err)
			}
			consumer.Ack(msg)
			continue

		// txn is always pending
		case middleware.WardenTimeout:
			isExists, err := ctx.db.IsWardenOnboardExist(onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.Batch)
			if err != nil {
				log.Fatalf("DB error: %v", err)
			}
			if isExists {
				log.Infoln("txn already exists")
				consumer.Ack(msg)
				continue
			}

			enclaveOnboard, err := ctx.db.SelectEnclaveOnboard(onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.Batch)
			if err != nil {
				log.Fatalf("DB error: %v", err)
			}
			if enclaveOnboard != nil {
				err = ctx.db.InsertPendingWardenOnboard(
					onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.ContractAddress, onboardTxn.AccountAddress, enclaveOnboard.OnboardTxnHash,
					onboardTxn.BlockNumber, enclaveOnboard.Nonce, onboardTxn.Batch, onboardTxn.TxnIndex,
					onboardTxn.ChainId, onboardTxn.Amount,
					handlerTx,
				)
				if err != nil {
					handlerTx.Rollback()
					log.Fatalf("DB error: %v", err)
				}

				err = handlerTx.Commit()
				if err != nil {
					handlerTx.Rollback()
					log.Fatalf("DB error: %v", err)
				}
				consumer.Ack(msg)
				continue
			}

			err = ctx.db.InsertWardenOnboard(
				onboardTxn.BlockHash, onboardTxn.TxnHash, onboardTxn.ContractAddress, onboardTxn.AccountAddress,
				onboardTxn.ChainId, onboardTxn.Amount, onboardTxn.BlockNumber, onboardTxn.Batch, onboardTxn.TxnIndex,
				handlerTx,
			)
			if err != nil {
				handlerTx.Rollback()
				log.Fatalf("DB error: %v", err)
			}

			err = handlerTx.Commit()
			if err != nil {
				handlerTx.Rollback()
				log.Fatalf("DB error: %v", err)
			}
			consumer.Ack(msg)
			continue
		}

	}

	// if err := consumer.Unsubscribe(); err != nil {
	// 	log.Fatal(err)
	// }
}
