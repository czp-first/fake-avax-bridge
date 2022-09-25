package dal

import (
	"database/sql"
	"math/big"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/sqldb"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DAL struct {
	*sqldb.Db
}

func NewDAL(driver, info string, poolSize int) (*DAL, error) {
	db, err := sqldb.NewDb(driver, info, poolSize)
	if err != nil {
		log.Errorf("fail with db init: %s, %s, %d, err: %+v", driver, info, poolSize, err)
		return nil, err
	}

	dal := &DAL{
		db,
	}

	return dal, nil
}

func (d *DAL) Close() {
	if d.Db != nil {
		d.Db.Close()
		d.Db = nil
	}
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Warnln("closeRows: error:", err)
	}
}

func (d *DAL) DB() *sqldb.Db {
	return d.Db
}

// onboard
// TODO: batch
func (d *DAL) GetWardenOnboardByHash(blockHash, txnHash string) (*sqldb.WardenOnboard, bool, error) {
	var amount, blockNumber, chainId int64
	var contract, account, status string
	var txnIndex int

	q := `SELECT chain_id, contract, account, amount, block_number, txn_index, status FROM warden_onboard where block_hash=$1 and txn_hash=$2`
	err := d.QueryRow(q, blockHash, txnHash).Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	found, err := sqldb.ChkQueryRow(err)
	if found {
		onboardTxn := &sqldb.WardenOnboard{
			ChainId:  big.NewInt(chainId),
			Amount:   big.NewInt(amount),
			Contract: contract,
			Account:  account,
		}
		return onboardTxn, true, nil
	}
	return nil, found, err
}

func (d *DAL) GetOldestPendingWardenOnboard() (*sqldb.WardenOnboard, bool, error) {
	var blockNumber, nonce uint64
	var amount, batch, chainId, rowId int64
	var contract, account, status, blockHash, txnHash, onboardTxnHash string
	var txnIndex uint

	q := `SELECT
			id, chain_id, contract, account, amount, block_number, txn_index,
			status, block_hash, txn_hash, nonce, onboard_txn_hash, batch
		FROM
			warden_onboard
		WHERE
			status=$1
		ORDER BY id
		LIMIT 1`
	err := d.Db.QueryRow(q, sqldb.PEDNING).Scan(&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &nonce, &onboardTxnHash, &batch)
	found, err := sqldb.ChkQueryRow(err)
	if found {
		w := &sqldb.WardenOnboard{
			RowId:          rowId,
			ChainId:        big.NewInt(chainId),
			BlockNumber:    blockNumber,
			Amount:         big.NewInt(amount),
			Contract:       contract,
			Account:        account,
			BlockHash:      blockHash,
			TxnHash:        txnHash,
			Nonce:          nonce,
			OnboardTxnHash: onboardTxnHash,
			TxnIndex:       txnIndex,
			Batch:          batch,
		}
		return w, true, nil
	}
	return nil, found, err

}

func (d *DAL) UpdateWardenOnboardStatusById(rowId int64, status string) error {
	q := `UPDATE warden_onboard SET status=$1 WHERE id=$2`
	res, err := d.Exec(q, status, rowId)
	return sqldb.ChkExec(res, err, 1, "UpdateWardenOnboardStatusById")
}

func (d *DAL) GetOldestInitWardenOnboard() (*sqldb.WardenOnboard, bool, error) {
	var blockNumber uint64
	var amount, batch, chainId, rowId int64
	var contract, account, status, blockHash, txnHash string
	var txnIndex int

	q := `SELECT
			id, chain_id, contract, account, amount, block_number,
			txn_index, status, block_hash, txn_hash, batch
		FROM
			warden_onboard
		WHERE
			status=$1
		ORDER BY
			block_number, txn_index
		LIMIT 1`
	err := d.Db.QueryRow(q, sqldb.INIT).Scan(
		&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &batch)

	found, err := sqldb.ChkQueryRow(err)
	if found {
		w := &sqldb.WardenOnboard{
			RowId:       rowId,
			ChainId:     big.NewInt(chainId),
			Amount:      big.NewInt(amount),
			Contract:    contract,
			Account:     account,
			BlockHash:   blockHash,
			TxnHash:     txnHash,
			Batch:       batch,
			BlockNumber: blockNumber,
		}
		return w, true, nil
	}

	return nil, found, err
}

func (d *DAL) DoneWardenOnboardById(rowId int64) error {
	q := `UPDATE warden_onboard SET status=$1 WHERE id=$2`
	res, err := d.Db.Exec(q, sqldb.DONE, rowId)
	return sqldb.ChkExec(res, err, 1, "DoneWardenOnboardById")
}

// offboard
func (d *DAL) GetWardenOffboardByHash(blockHash, txnHash string) (*sqldb.WardenOffboard, bool, error) {
	var amount, blockNumber, chainId int64
	var contract, account, status string
	var txnIndex int

	// TODO: maybe many rows
	q := `SELECT chain_id, contract, account, amount, block_number, txn_index, status
			FROM warden_offboard
			WHERE block_hash=$1 AND txn_hash=$2`
	err := d.Db.QueryRow(q, blockHash, txnHash).Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	found, err := sqldb.ChkQueryRow(err)
	if found {
		o := &sqldb.WardenOffboard{
			ChainId:  big.NewInt(chainId),
			Amount:   big.NewInt(amount),
			Contract: contract,
			Account:  account,
		}
		return o, true, nil
	}

	return nil, found, err
}
