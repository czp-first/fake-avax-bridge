package core

// func (ctx *WardenContext) seeDxchainBlock() {

// 	ctx.bridgeSettings.Lock.RLock()
// 	preSettings := ctx.bridgeSettings.Settings
// 	dxMinConfirmations := preSettings.Get("nonCritical").Get("minimumConfirmations").Get("dxchain").MustUint64()
// 	preDxLastestIndexedBlockNum := preSettings.Get("nonCritical").Get("networkViews").Get("dxchain").Get("lastIndexedBlock").MustUint64()
// 	preDxLastestSeenBlockNum := preSettings.Get("nonCritical").Get("networkViews").Get("dxchain").Get("lastSeenBlock").MustUint64()
// 	dxLastestBlockNum, err := ctx.DxClient.HttpClient.BlockNumber(context.Background())
// 	ctx.bridgeSettings.Lock.RUnlock()

// 	if err != nil {
// 		log.Fatalf("Get dxchain blockNumber err: %v\n", err)
// 	}
// 	log.Infof(
// 		"dxchain: dxMinConfirmations[%d], preDxLastestIndexedBlockNum[%d], preDxLastestSeenBlockNum[%d], dxLastestBlockNum[%d]",
// 		dxMinConfirmations, preDxLastestIndexedBlockNum, preDxLastestSeenBlockNum, dxLastestBlockNum,
// 	)

// 	chainlinkDxUsdFeedAddr := preSettings.Get("nonCritical").Get("chainlinkDxUsdFeedAddress").MustString()
// 	currentDxPrice, _ := chainlinkfeed.GetFeedData(common.HexToAddress(chainlinkDxUsdFeedAddr))
// 	log.Infof("Current dx price: %v\n", currentDxPrice)

// 	ctx.bridgeSettings.Lock.Lock()
// 	preSettings.SetPath([]string{"nonCritical", "currentDxPrice"}, currentDxPrice.Int64())

// 	if (dxLastestBlockNum - preDxLastestSeenBlockNum) > dxMinConfirmations {
// 		readyDxSeenBlockNum := preDxLastestSeenBlockNum + 1
// 		log.Infof("Ready index dxchain block[%d]\n", readyDxSeenBlockNum)

// 		tokensAddr := ctx.bridgeSettings.GetTokenAddrs(ctx.bridgeSettings.Networks.DxChain)
// 		query, err := ctx.makeFilterQuery(tokensAddr, contracts.DxErc20MetaData.ABI, "Unwrap", big.NewInt(int64(readyDxSeenBlockNum)), big.NewInt(int64(readyDxSeenBlockNum)), [][]common.Hash{})
// 		if err != nil {
// 			log.Errorln(err)
// 			return
// 		}

// 		ctx.DxClient.SeeBlock(query, ctx.bridgeSettings, ctx.offboard)
// 		log.Infof("Finish index dxchain block[%d]\n", readyDxSeenBlockNum)
// 		preSettings.SetPath([]string{"nonCritical", "networkViews", "dxchain", "lastIndexedBlock"}, dxLastestBlockNum)
// 		preSettings.SetPath([]string{"nonCritical", "networkViews", "dxchain", "lastSeenBlock"}, readyDxSeenBlockNum)
// 	}

// 	ctx.bridgeSettings.UpdateSettings()
// 	ctx.bridgeSettings.Lock.Unlock()
// }

// func (ctx *WardenContext) SeeDxchainBlock() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			ctx.seeDxchainBlock()
// 		}
// 	}
// }

// func (ctx *WardenContext) ConsumeOffboardTxn() {
// 	for {
// 		select {
// 		case offboardTxn := <-ctx.offboard:
// 			handlerTx, err := ctx.db.GetDB().Begin()
// 			if err != nil {
// 				log.Errorf("DB error: %v\n", err)
// 				continue
// 			}

// 			switch offboardTxn.Type {
// 			// receive offboard txn from enclave
// 			case chains.EnclaveTell:
// 				ctx.db.Offboard(
// 					offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.OffboardTxnHash, offboardTxn.Nonce, offboardTxn.Batch, handlerTx,
// 				)

// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}

// 				err = handlerTx.Commit()
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}
// 				continue

// 			// warden self seen offboard txn
// 			case chains.WardenSee:
// 				isExist, err := ctx.db.IsOffboardTxnExist(offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.Batch)
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					continue
// 				}
// 				if isExist {
// 					log.Infoln("txn already exists")
// 					continue
// 				}

// 				offboard, err := ctx.db.SelectOffboard(offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.Batch)
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					continue
// 				}

// 				if offboard != nil {
// 					err = ctx.db.InsertCompleteOffboardTxn(
// 						offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.ContractAddress, offboardTxn.AccountAddress,
// 						offboard.OffboardTxnHash, offboardTxn.ChainId, offboardTxn.BlockNumber, offboard.Nonce, offboardTxn.Batch,
// 						offboardTxn.TxIndex, offboardTxn.Amount,
// 						handlerTx,
// 					)

// 					if err != nil {
// 						log.Errorf("DB error: %v\n", err)
// 						handlerTx.Rollback()
// 						continue
// 					}

// 					err = handlerTx.Commit()
// 					if err != nil {
// 						log.Errorf("DB error: %v\n", err)
// 						handlerTx.Rollback()
// 						continue
// 					}
// 					continue
// 				}

// 				err = ctx.db.InsertOffboardTxn(
// 					offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.ContractAddress, offboardTxn.AccountAddress,
// 					offboardTxn.ChainId, offboardTxn.BlockNumber, offboardTxn.Batch, offboardTxn.TxIndex, offboardTxn.Amount,
// 					handlerTx,
// 				)

// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}

// 				err = handlerTx.Commit()
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}

// 			// txn is always pending
// 			case chains.WardenTimeout:
// 				isExist, err := ctx.db.IsOffboardTxnExist(offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.Batch)
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					continue
// 				}
// 				if isExist {
// 					log.Infoln("txn already exists")
// 					continue
// 				}

// 				offboard, err := ctx.db.SelectOffboard(offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.Batch)
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					continue
// 				}

// 				if offboard != nil {
// 					err = ctx.db.InsertCompleteOffboardTxn(
// 						offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.ContractAddress, offboardTxn.AccountAddress,
// 						offboard.OffboardTxnHash, offboardTxn.ChainId, offboardTxn.BlockNumber, offboard.Nonce, offboardTxn.Batch,
// 						offboardTxn.TxIndex, offboardTxn.Amount,
// 						handlerTx,
// 					)

// 					if err != nil {
// 						log.Errorf("DB error: %v\n", err)
// 						handlerTx.Rollback()
// 						continue
// 					}

// 					err = handlerTx.Commit()
// 					if err != nil {
// 						log.Errorf("DB error: %v\n", err)
// 						handlerTx.Rollback()
// 						continue
// 					}
// 					continue
// 				}

// 				err = ctx.db.InsertOffboardTxn(
// 					offboardTxn.BlockHash, offboardTxn.TxnHash, offboardTxn.ContractAddress, offboardTxn.AccountAddress,
// 					offboardTxn.ChainId, offboardTxn.BlockNumber, offboardTxn.Batch, offboardTxn.TxIndex, offboardTxn.Amount,
// 					handlerTx,
// 				)

// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}

// 				err = handlerTx.Commit()
// 				if err != nil {
// 					log.Errorf("DB error: %v\n", err)
// 					handlerTx.Rollback()
// 					continue
// 				}
// 			}
// 		}
// 	}
// }

// func (ctx *WardenContext) confirmOffboard(offboardTxn *database.OffboardTxn) {

// 	log.Infof("Confirm pending offboard txn: %v\n", offboardTxn.OffboardTxnHash)

// 	if offboardTxn == nil {
// 		return
// 	}

// 	isOk, err := ctx.EthClient.ConfirmTxn(offboardTxn.OffboardTxnHash)
// 	if err != nil {
// 		log.Errorf("DB error: %v\n", err)
// 		return
// 	}
// 	log.Infof("Pending offboard txn confirm status: %v\n", isOk)

// 	if !isOk {
// 		return
// 	}

// 	handlerTx, err := ctx.db.GetDB().Begin()
// 	if err != nil {
// 		log.Errorf("DB error: %v\n", err)
// 		return
// 	}

// 	err = ctx.db.UpdateOffboardTxnByOffboardTxnHash(offboardTxn.OffboardTxnHash, handlerTx)

// 	if err != nil {
// 		log.Errorf("DB error: %v\n", err)
// 		handlerTx.Rollback()
// 		return
// 	}

// 	err = handlerTx.Commit()
// 	if err != nil {
// 		log.Errorf("DB error: %v\n", err)
// 		handlerTx.Rollback()
// 		return
// 	}

// }

// func (ctx *WardenContext) monitorOffboard() {
// 	ctx.bridgeSettings.Lock.RLock()
// 	settings := ctx.bridgeSettings.Settings
// 	lastSeenBlockNum := settings.Get("nonCritical").Get("networkViews").Get("dxchain").Get("lastSeenBlock").MustInt64()
// 	ctx.bridgeSettings.Lock.RUnlock()

// 	identification, err := ioutil.ReadFile(os.Getenv("IdentificationFilePath"))
// 	if err != nil {
// 		log.Errorf("Fail read identification file: %v\n", err)
// 	}

// 	// select the oldest pending offboard txn
// 	pendingOffboardTxn, err := ctx.db.SelectPendingOffboardTxn()
// 	if err != nil {
// 		log.Errorf("DB error: %v\n", err)
// 		return
// 	}

// 	// compare blockNumber of the oldest pending onboard txn with lastSeenBlockNum
// 	if pendingOffboardTxn != nil && lastSeenBlockNum-pendingOffboardTxn.BlockNumber > 200 {
// 		isOk, err := ctx.EthClient.ConfirmTxn(pendingOffboardTxn.OffboardTxnHash)
// 		if err != nil {
// 			log.Errorf("Confirm error: %v\n", err)
// 			return
// 		}
// 		if isOk {
// 			ctx.confirmOffboard(pendingOffboardTxn)
// 			return
// 		}

// 		// txn timeout
// 		handlerTx, err := ctx.db.GetDB().Begin()
// 		if err != nil {
// 			log.Errorf("DB error :%v\n", err)
// 			return
// 		}

// 		err = ctx.db.UpdateInitOffboardTxn(pendingOffboardTxn.RowId, "timeout", handlerTx)
// 		if err != nil {
// 			log.Errorf("DB error :%v\n", err)
// 			handlerTx.Rollback()
// 			return
// 		}

// 		err = handlerTx.Commit()
// 		if err != nil {
// 			log.Errorf("DB error :%v\n", err)
// 			handlerTx.Rollback()
// 			return
// 		}

// 		ctx.offboard <- &chains.OffboardTxn{
// 			Type:            chains.WardenTimeout,
// 			BlockHash:       pendingOffboardTxn.BlockHash,
// 			TxnHash:         pendingOffboardTxn.TxnHash,
// 			ContractAddress: pendingOffboardTxn.Contract,
// 			AccountAddress:  pendingOffboardTxn.Account,
// 			ChainId:         pendingOffboardTxn.ChainId,
// 			BlockNumber:     pendingOffboardTxn.BlockNumber,
// 			TxIndex:         pendingOffboardTxn.TxIndex,
// 			Amount:          pendingOffboardTxn.Amount,
// 			Batch:           pendingOffboardTxn.Batch + 1,
// 		}
// 		return
// 	}

// 	initOffboardTxn := ctx.db.SelectInitOffboardTxn()

// 	if initOffboardTxn == nil {
// 		log.Infoln("No init offboard txn")
// 		return
// 	}

// 	// enclave ReceiveOffboardTxn
// 	resp := ctx.Enclave.ReceiveOffboardTxn(
// 		&enclavepb.OffboardTxn{
// 			Batch:          initOffboardTxn.Batch,
// 			BlockHash:      initOffboardTxn.BlockHash,
// 			TxnHash:        initOffboardTxn.TxnHash,
// 			Identification: string(identification),
// 		},
// 	)
// 	log.Infof("Enclave resp offboard txn status: %v\n", resp.Status)

// 	handlerTx, err := ctx.db.GetDB().Begin()
// 	if err != nil {
// 		log.Errorf("DB error :%v\n", err)
// 		return
// 	}

// 	err = ctx.db.UpdateInitOffboardTxn(initOffboardTxn.RowId, resp.Status, handlerTx)
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

// }

// func (ctx *WardenContext) MonitorOffboard() {
// 	ticker := time.NewTicker(10 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			ctx.monitorOffboard()
// 		}
// 	}
// }

// func (ctx *WardenContext) ConfirmOffboard() {
// 	ticker := time.NewTicker(10 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			offboardTxn, err := ctx.db.SelectPendingOffboardTxn()
// 			if err != nil {
// 				log.Errorf("DB error: %v\n", err)
// 				continue
// 			}

// 			if offboardTxn == nil {
// 				continue
// 			}
// 			ctx.confirmOffboard(offboardTxn)
// 		}
// 	}
// }
