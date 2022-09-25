package core

import (
	"time"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/sqldb"

	log "github.com/sirupsen/logrus"
)

// confirm onboard txn in dx chain
func (ctx *WardenContext) confirmOnboard(wardenOnboard *sqldb.WardenOnboard) error {

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

	dbErr := ctx.db.DoneWardenOnboardById(wardenOnboard.RowId)

	if dbErr != nil {
		log.Errorf("confirm onboard: done warden onboard by id, error:%v", dbErr)
		return err
	}

	return nil
}

func (ctx *WardenContext) ConfirmOnboard() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		oldestPendingWardenOnboard, found, err := ctx.db.GetOldestPendingWardenOnboard()
		if err != nil {
			log.Errorf("confirm onboard: retrieve oldest pending warden onboard error: %v", err)
			continue
		}
		if !found {
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
