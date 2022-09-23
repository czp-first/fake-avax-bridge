package core

import (
	"time"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/database"

	log "github.com/sirupsen/logrus"
)

// confirm onboard txn in dx chain
func (ctx *WardenContext) confirmOnboard(wardenOnboard *database.WardenOnboard) error {

	log.Infof("confirm onboard: txn: %+v", wardenOnboard)
	if wardenOnboard == nil {
		return nil
	}

	isOk, err := ctx.ToChainClient.ConfirmTxn(wardenOnboard.OnboardTxnHash)
	if err != nil {
		log.Errorf("confirm onboard: to chain client onboard err: %v", err)
		return err
	}

	log.Infof("confirm onboard: txn status: %v", isOk)

	if !isOk {
		return nil
	}

	handlerTx, err := ctx.db.GetDB().Begin()
	if err != nil {
		log.Errorf("confirm onboard: start DB error: %v", err)
		return err
	}

	err = ctx.db.DoneWardenOnboardByOnboardTxnHash(wardenOnboard.OnboardTxnHash, handlerTx)

	if err != nil {
		log.Errorf("confirm onboard: done warden onboard by onboard txn hash error: %v", err)
		handlerTx.Rollback()
		return err
	}

	err = handlerTx.Commit()
	if err != nil {
		log.Errorf("confirm onboard: pg commit error: %v", err)
		handlerTx.Rollback()
		return err
	}

	return nil
}

func (ctx *WardenContext) ConfirmOnboard() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		oldestPendingWardenOnboard, err := ctx.db.RetrieveOldestPendingWardenOnboard()
		if err != nil {
			log.Errorf("confirm onboard: retrieve oldest pending warden onboard error: %v", err)
			continue
		}
		if oldestPendingWardenOnboard == nil {
			log.Info("confirm onboard: no pending onboard txn")
			continue
		}
		err = ctx.confirmOnboard(oldestPendingWardenOnboard)
		if err != nil {
			log.Errorf("confirm onboard: error: %v", err)
			continue
		}
	}
}
