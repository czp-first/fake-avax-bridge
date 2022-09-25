package core

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/czp-first/fake-avax-bridge/EnclaveProxy/enclavepb"

	log "github.com/sirupsen/logrus"
)

func (ctx *WardenContext) monitorWardenOnboard() error {

	// ctx.bridgeSettings.Lock.RLock()
	// settings := ctx.bridgeSettings.Settings
	// lastSeenBlockNum := settings.Get("nonCritical").Get("networkViews").Get("ethereum").Get("lastSeenBlock").MustInt64()
	// ctx.bridgeSettings.Lock.RUnlock()

	// settings := ctx.bridgeSettings.GetSettings()
	// lastSeenBlockNum := settings.NonCritical.NetworkViews.Ethereum.LastSeenBlock

	// // select the oldest pending warden onboard
	// oldestPendingWardenOnboard, err := ctx.db.GetOldestPendingWardenOnboard()
	// if err != nil {
	// 	log.Errorf("DB error: %v\n", err)
	// 	return
	// }

	// compare blockNumber of the oldest pending onboard txn with lastSeenBlockNum
	// if oldestPendingWardenOnboard != nil && lastSeenBlockNum-oldestPendingWardenOnboard.BlockNumber > 200 {
	// 	log.Infof("exist timeout warden onboard: %+v", oldestPendingWardenOnboard)
	// 	isOk, err := ctx.DxClient.ConfirmTxn(oldestPendingWardenOnboard.OnboardTxnHash)
	// 	if err != nil {
	// 		log.Errorf("monitor: Confirm error: %v\n", err)
	// 		return
	// 	}
	// 	if isOk {
	// 		ctx.confirmOnboard(oldestPendingWardenOnboard)
	// 		log.Info("timeout warden onboard already confirm")
	// 		return
	// 	}

	// 	// txn timeout
	// 	// TODO: table onboard status
	// 	handlerTx, err := ctx.db.GetDB().Begin()
	// 	if err != nil {
	// 		log.Errorf("DB error :%v\n", err)
	// 		return
	// 	}
	// 	err = ctx.db.UpdateWardenOnboardStatusById(oldestPendingWardenOnboard.RowId, "timeout", handlerTx)
	// 	if err != nil {
	// 		log.Errorf("DB error :%v\n", err)
	// 		handlerTx.Rollback()
	// 		return
	// 	}

	// 	err = handlerTx.Commit()
	// 	if err != nil {
	// 		log.Errorf("DB error :%v\n", err)
	// 		handlerTx.Rollback()
	// 		return
	// 	}

	// 	jsonSchema := pulsar.NewJSONSchema(onboardTxnSchemaDef, nil)

	// 	onboardTxnProducer, err := ctx.pulsarCli.CreateProducer(pulsar.ProducerOptions{
	// 		Topic:  os.Getenv("PulsarTopic"),
	// 		Schema: jsonSchema,
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("Could not instance onboardTxn producer: %v", err)
	// 	}
	// 	defer onboardTxnProducer.Close()
	// 	_, err = onboardTxnProducer.Send(context.Background(), &pulsar.ProducerMessage{
	// 		Value: &OnboardTxnJSON{
	// 			Type:            WardenTimeout,
	// 			BlockHash:       oldestPendingWardenOnboard.BlockHash,
	// 			TxnHash:         oldestPendingWardenOnboard.TxnHash,
	// 			ContractAddress: oldestPendingWardenOnboard.Contract,
	// 			AccountAddress:  oldestPendingWardenOnboard.Account,
	// 			ChainId:         oldestPendingWardenOnboard.ChainId,
	// 			BlockNumber:     oldestPendingWardenOnboard.BlockNumber,
	// 			TxnIndex:        oldestPendingWardenOnboard.TxnIndex,
	// 			Amount:          oldestPendingWardenOnboard.Amount,
	// 			Batch:           oldestPendingWardenOnboard.Batch + 1,
	// 		},
	// 	})

	// 	if err != nil {
	// 		log.Fatalf("onboard: publish message err: %v", err)
	// 	}
	// 	log.Infoln("onboard: published message")

	// 	return
	// }

	oldestInitWardenOnboard, found, err := ctx.db.GetOldestInitWardenOnboard()
	if err != nil {
		log.Errorf("monitor onboard: retrieve oldest init warden onboard: %v", err)
		return err
	}
	if !found {
		log.Infoln("monitor onboard: No init onboard txn")
		return nil
	}
	log.Infof("monitor onboard: ready onboard: %+v", oldestInitWardenOnboard)

	identification, err := ioutil.ReadFile(os.Getenv("IdentificationFilePath"))
	if err != nil {
		log.Errorf("monitor onboard: Fail read identification file: %v", err)
		return err
	}
	resp := ctx.Enclave.ReceiveOnboardTxn(
		&enclavepb.OnboardTxn{
			Batch:          oldestInitWardenOnboard.Batch,
			BlockHash:      oldestInitWardenOnboard.BlockHash,
			TxnHash:        oldestInitWardenOnboard.TxnHash,
			Identification: string(identification),
		},
	)
	log.Infof("monitor onboard: Enclave resp onboard txn status: %v", resp.Status)

	err = ctx.db.UpdateWardenOnboardStatusById(oldestInitWardenOnboard.RowId, resp.Status)
	if err != nil {
		log.Errorf("monitor onboard: update warden onboard status by id error: %v", err)
		return err
	}

	return nil
}

func (ctx *WardenContext) MonitorOnboard() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := ctx.monitorWardenOnboard()
		if err != nil {
			log.Errorf("monitor onboard: %v", err)
			continue
		}
	}
}
