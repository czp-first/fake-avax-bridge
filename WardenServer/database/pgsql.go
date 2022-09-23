package database

import (
	"database/sql"
	"fmt"
	"math/big"
	"os"

	"github.com/czp-first/fake-avax-bridge/BridgeUtils/database"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type PgSQL struct {
	Db *sql.DB
}

func initDB(driverName, constr string) (*sql.DB, error) {
	dbCon, err := sql.Open(driverName, constr)
	if err != nil {
		return nil, err
	}

	err = dbCon.Ping()
	if err != nil {
		return nil, err
	}

	return dbCon, nil
}

func NewPgSQL() (*PgSQL, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PgHost"), os.Getenv("PgPort"), os.Getenv("PgUser"), os.Getenv("PgPassword"), os.Getenv("PgDb"),
	)

	dbCon, err := initDB("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return &PgSQL{Db: dbCon}, nil
}

func (con *PgSQL) Close() error {
	return con.Db.Close()
}

func (con *PgSQL) GetDB() *sql.DB {
	return con.Db
}

// onboard
func (con *PgSQL) SelectWardenOnboard(blockHash, txnHash string) *database.WardenOnboard {
	var amount, blockNumber, chainId int64
	var contract, account, status string
	var txnIndex int

	row := con.Db.QueryRow(
		`select chain_id, contract, account, amount, block_number, txn_index, status from warden_onboard where block_hash=$1 and txn_hash=$2`,
		blockHash, txnHash,
	)

	err := row.Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	if err != nil {
		log.Errorf("Fail select warden onboard: %v", err)
	}

	onboardTxn := &database.WardenOnboard{
		ChainId:  big.NewInt(chainId),
		Amount:   big.NewInt(amount),
		Contract: contract,
		Account:  account,
	}
	return onboardTxn
}

func (con *PgSQL) RetrieveOldestPendingWardenOnboard() (*database.WardenOnboard, error) {

	var blockNumber, nonce uint64
	var amount, batch, chainId, rowId int64
	var contract, account, status, blockHash, txnHash, onboardTxnHash string
	var txnIndex uint

	row := con.Db.QueryRow(`select
			id, chain_id, contract, account, amount, block_number, txn_index,
			status, block_hash, txn_hash, nonce, onboard_txn_hash, batch
		from
			warden_onboard
		where
			status=$1
		order by id
		limit 1`, "pending")
	err := row.Scan(&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &nonce, &onboardTxnHash, &batch)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("confirm onboard: Fail select oldest pending warden onboard: %v", err)
		return nil, err
	}
	return &database.WardenOnboard{
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
	}, nil

}

func (con *PgSQL) UpdateWardenOnboardStatusById(rowId int64, status string, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`update warden_onboard set status=$1 where id=$2`)
	if err != nil {
		log.Errorf("Fail update init warden onboard: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, rowId)
	if err != nil {
		log.Errorf("Fail update init warden onboard: %v", err)
		return err
	}
	return nil
}

func (con *PgSQL) RetrieveOldestInitWardenOnboard() (*database.WardenOnboard, error) {
	var blockNumber uint64
	var amount, batch, chainId, rowId int64
	var contract, account, status, blockHash, txnHash string
	var txnIndex int

	row := con.Db.QueryRow(`select
			id, chain_id, contract, account, amount, block_number,
			txn_index, status, block_hash, txn_hash, batch
		from
			warden_onboard
		where
			status=$1
		order by
			block_number, txn_index
		limit 1`, "init",
	)

	err := row.Scan(
		&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &batch,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Fail select oldest init warden onboard: %v\n", err)
		return nil, err
	}

	return &database.WardenOnboard{
		RowId:       rowId,
		ChainId:     big.NewInt(chainId),
		Amount:      big.NewInt(amount),
		Contract:    contract,
		Account:     account,
		BlockHash:   blockHash,
		TxnHash:     txnHash,
		Batch:       batch,
		BlockNumber: blockNumber,
	}, nil
}

func (con *PgSQL) DoneWardenOnboardByOnboardTxnHash(onboardTxnHash string, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`update warden_onboard set status=$1 where onboard_txn_hash=$2`)
	if err != nil {
		log.Errorf("Fail prepare update warden onboard by onboard_txn_hash: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("done", onboardTxnHash)
	if err != nil {
		return err
	}
	return nil
}

// offboard
func (con *PgSQL) SelectWardenOffboard(blockHash, txnHash string) *database.WardenOffboard {
	var amount, blockNumber, chainId int64
	var contract, account, status string
	var txnIndex int

	row := con.Db.QueryRow(`
		select chain_id, contract, account, amount, block_number, txn_index, status
		from warden_offboard
		where block_hash=$1 and txn_hash=$2
	`, blockHash, txnHash,
	)

	err := row.Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	if err != nil {
		log.Fatalf("Fail select warden offboard : %v\n", err)
	}

	offboardTxn := &database.WardenOffboard{
		ChainId:  big.NewInt(chainId),
		Amount:   big.NewInt(amount),
		Contract: contract,
		Account:  account,
	}
	return offboardTxn
}
