package database

import (
	"database/sql"
	"fmt"
	"math/big"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type PgSQL struct {
	db *sql.DB
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

	return &PgSQL{db: dbCon}, nil
}

func (con *PgSQL) Close() error {
	return con.db.Close()
}

func (con *PgSQL) GetDB() *sql.DB {
	return con.db
}

func (con *PgSQL) SelectOnboard(blockHash, txnHash string, batch int64) (*Onboard, error) {
	var onboardTxnHash string
	var nonce uint64

	row := con.db.QueryRow(
		`select onboard_txn_hash, nonce from onboard where block_hash=$1 and txn_hash=$2 and batch=$3 and status=$4`,
		blockHash, txnHash, batch, "normal")

	err := row.Scan(&onboardTxnHash, &nonce)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Fail select onboard: %v\n", err)
		return nil, err
	}
	return &Onboard{OnboardTxnHash: onboardTxnHash, Nonce: nonce}, nil
}

func (con *PgSQL) IsOnboardTxnExist(blockHash, txnHash string, batch int64) (bool, error) {

	var isExist int64
	row := con.db.QueryRow(
		`select 1 from onboard_txn where block_hash=$1 and txn_hash=$2 and batch=$3 and status<>$4
	`, blockHash, txnHash, batch, "timeout",
	)

	err := row.Scan(&isExist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (con *PgSQL) InsertCompleteOnboardTxn(blockHash, txnHash, contract, account, onboardTxnHash string, chainId, blockNumber, nonce, batch int64, txnIndex int, amount int64, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`
		insert into onboard_txn(
			block_hash, txn_hash, chain_id, contract, account, amount, block_number, txn_index, status, onboard_txn_hash, nonce, batch
		)values(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)
	`)
	if err != nil {
		log.Errorf("Fail insert complete onboard_txn: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blockHash, txnHash, chainId, contract, account, amount, blockNumber, txnIndex, "pending", onboardTxnHash, nonce, batch)
	if err != nil {
		log.Errorf("Fail insert complete onboard_txn: %v\n", err)
		return err
	}
	return nil
}

func (con *PgSQL) InsertOnboardTxn(blockHash, txnHash, contract, account string, chainId, blockNumber, batch int64, txnIndex int, amount int64, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(
		`
		insert into onboard_txn(
			block_hash, txn_hash, chain_id, contract, account, amount, block_number, txn_index, status, batch
		)values(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`)
	if err != nil {
		log.Errorf("Fail insert onboard_txn: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blockHash, txnHash, chainId, contract, account, amount, blockNumber, txnIndex, "init", batch)
	if err != nil {
		log.Errorf("Fail insert onboard_txn: %v\n", err)
		return err
	}
	return nil
}

func (con *PgSQL) SelectTxn(blockHash, txnHash string) *WardenOnboard {
	var amount, blockNumber, chainId *big.Int
	var contract, account, status string
	var txnIndex int

	row := con.db.QueryRow(
		`select chain_id, contract, account, amount, block_number, txn_index, status from onboard_txn where block_hash=$1 and txn_hash=$2`,
		blockHash, txnHash,
	)

	err := row.Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	if err != nil {
		log.Errorf("Fail select onboard_txn: %v\n", err)
	}

	onboardTxn := &WardenOnboard{
		ChainId:  chainId,
		Amount:   amount,
		Contract: contract,
		Account:  account,
	}
	return onboardTxn
}

func (con *PgSQL) RetrieveOldestPendingWardenOnboard() (*WardenOnboard, error) {

	var amount, chainId *big.Int
	var blockNumber, nonce uint64
	var rowId, batch int64
	var contract, account, status, blockHash, txnHash, onboardTxnHash string
	var txnIndex uint

	row := con.db.QueryRow(`
		select id, chain_id, contract, account, amount, block_number, txn_index, status, block_hash, txn_hash, nonce, onboard_txn_hash, batch from onboard_txn where status=$1 order by id limit 1
	`, "pending")
	err := row.Scan(&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &nonce, &onboardTxnHash, &batch)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Fail select pending onboard_txn: %v\n", err)
		return nil, err
	}
	return &WardenOnboard{
		RowId:          rowId,
		ChainId:        chainId,
		BlockNumber:    blockNumber,
		Amount:         amount,
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

func (con *PgSQL) SelectPendingOnboardTxnCount() (int, error) {
	var count int

	row := con.db.QueryRow(
		`select count(1) from onboard_txn where status=$1`, "pending",
	)

	err := row.Scan(&count)
	if err != nil {
		log.Errorf("Fail select pending onboard_txn count: %v\n", err)
		return 0, err
	}

	return count, nil
}

func (con *PgSQL) Onboard(blockHash, txnHash, onboardTxnHash string, nonce, batch int64, handler *sql.Tx) error {
	stmt1, err := handler.Prepare(
		`
		insert into onboard(
			block_hash, txn_hash, onboard_txn_hash, nonce, batch, status
		)values(
			$1, $2, $3, $4, $5, $6
		)
	`)
	if err != nil {
		log.Errorf("Fail insert onboard: %v\n", err)
		return err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(blockHash, txnHash, onboardTxnHash, nonce, batch, "normal")
	if err != nil {
		log.Errorf("Fail insert onboard: %v\n", err)
		return err
	}

	stmt2, err := handler.Prepare(
		`update onboard_txn set onboard_txn_hash=$1, status=$2, nonce=$3 where block_hash=$4 and txn_hash=$5 and batch=$6 and status<>$7
	`)
	if err != nil {
		log.Errorf("Fail update onboard_txn: %v\n", err)
		return err
	}
	defer stmt1.Close()

	_, err = stmt2.Exec(onboardTxnHash, "pending", nonce, blockHash, txnHash, batch, "timeout")
	if err != nil {
		log.Errorf("Fail update onboard_txn: %v\n", err)
		return err
	}
	return nil
}

// offboard

func (con *PgSQL) SelectOffboard(blockHash, txnHash string, batch int64) (*Offboard, error) {
	var offboardTxnHash string
	var nonce uint64

	row := con.db.QueryRow(
		`select offboard_txn_hash, nonce from offboard where block_hash=$1 and txn_hash=$2 and batch=$3 and status=$4`,
		blockHash, txnHash, batch, "normal",
	)

	err := row.Scan(&offboardTxnHash, &nonce)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Fail select offboard: %v\n", err)
		return nil, err
	}
	return &Offboard{OffboardTxnHash: offboardTxnHash, Nonce: nonce}, nil
}

func (con *PgSQL) IsOffboardTxnExist(blockHash, txnHash string, batch int64) (bool, error) {
	var isExist int64
	row := con.db.QueryRow(
		`select 1 from onboard_txn where block_hash=$1 and txn_hash=$2 and batch=$3 and status<>$4
	`, blockHash, txnHash, batch, "timeout",
	)

	err := row.Scan(&isExist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (con *PgSQL) InsertCompleteOffboardTxn(blockHash, txnHash, contract, account, offboardTxnHash string, chainId, blockNumber, nonce, batch int64, txnIndex int, amount int64, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`
		insert into offboard_txn(
			block_hash, txn_hash, chain_id, contract, account, amount, block_number, txn_index, status, offboard_txn_hash, nonce, batch
		)values(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)
	`)
	if err != nil {
		log.Errorf("Fail insert complete offboard_txn: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blockHash, txnHash, chainId, contract, account, amount, blockNumber, txnIndex, "pending", offboardTxnHash, nonce, batch)
	if err != nil {
		log.Errorf("Fail insert complete offboard_txn: %v\n", err)
		return err
	}
	return nil
}

func (con *PgSQL) InsertOffboardTxn(blockHash, txnHash, contract, account string, chainId, blockNumber, batch int64, txnIndex int, amount int64, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(
		`insert into offboard_txn(
			block_hash, txn_hash, chain_id, contract, account, amount, block_number, txn_index, status, batch
		)values(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`)
	if err != nil {
		log.Errorf("Fail insert offboard_txn: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blockHash, txnHash, chainId, contract, account, amount, blockNumber, txnIndex, "init", batch)
	if err != nil {
		log.Errorf("Fail insert offboard_txn: %v\n", err)
		return err
	}

	return nil
}

func (con *PgSQL) SelectInitOffboardTxn() *WardenOffboard {
	var chainId, amount *big.Int
	var rowId, blockNumber, batch int64
	var contract, account, status, blockHash, txnHash string
	var txnIndex int

	row := con.db.QueryRow(
		`
		select id, chain_id, contract, account, amount, block_number, txn_index, status, block_hash, txn_hash, batch
		from offboard_txn where status=$1 order by id limit 1
	`, "init",
	)

	err := row.Scan(
		&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &batch,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Errorf("Fail select init offboard txn: %v\n", err)
		return nil
	}

	offboardTxn := &WardenOffboard{
		RowId:     rowId,
		ChainId:   chainId,
		Amount:    amount,
		Contract:  contract,
		Account:   account,
		BlockHash: blockHash,
		TxnHash:   txnHash,
		Batch:     batch,
	}
	return offboardTxn
}

func (con *PgSQL) SelectPendingOffboardTxn() (*WardenOffboard, error) {

	var amount, chainId *big.Int
	var blockNumber, nonce uint64
	var rowId, batch int64
	var contract, account, status, blockHash, txnHash, offboardTxnHash string
	var txnIndex int

	row := con.db.QueryRow(`
		select id, chain_id, contract, account, amount, block_number, txn_index, status, block_hash, txn_hash, nonce, offboard_txn_hash, batch from offboard_txn where status=$1 order by id limit 1
	`, "pending")
	err := row.Scan(&rowId, &chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status,
		&blockHash, &txnHash, &nonce, &offboardTxnHash, &batch)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Fail select pending offboard_txn: %v\n", err)
		return nil, err
	}
	return &WardenOffboard{
		RowId:           rowId,
		ChainId:         chainId,
		BlockNumber:     blockNumber,
		Amount:          amount,
		Contract:        contract,
		Account:         account,
		BlockHash:       blockHash,
		TxnHash:         txnHash,
		Nonce:           nonce,
		OffboardTxnHash: offboardTxnHash,
		TxIndex:         txnIndex,
		Batch:           batch,
	}, nil

}

func (con *PgSQL) SelectPendingOffboardTxnCount() (int, error) {
	var count int

	row := con.db.QueryRow(
		`select count(1) from offboard_txn where status=$1`, "pending",
	)

	err := row.Scan(&count)
	if err != nil {
		log.Errorf("Fail select pending offboard_txn: %v\n", err)
		return 0, err
	}

	return count, nil
}

func (con *PgSQL) SelectOffboardTxn(blockHash, txnHash string) *WardenOffboard {
	var amount, chainId *big.Int
	var blockNumber uint64
	var contract, account, status string
	var txnIndex int

	row := con.db.QueryRow(
		`
		select chain_id, contract, account, amount, block_number, txn_index, status from offboard_txn where block_hash=$1 and txn_hash=$2
	`, blockHash, txnHash,
	)

	err := row.Scan(&chainId, &contract, &account, &amount, &blockNumber, &txnIndex, &status)
	if err != nil {
		log.Errorf("Fail select offboard_txn :%v\n", err)
	}

	offboardTxn := &WardenOffboard{
		ChainId:  chainId,
		Amount:   amount,
		Contract: contract,
		Account:  account,
	}
	return offboardTxn
}

func (con *PgSQL) UpdateInitOffboardTxn(rowId int64, status string, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`
		update offboard_txn set status=$1 where id=$2
	`)
	if err != nil {
		log.Errorf("Fail update offboard_txn: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, rowId)
	if err != nil {
		log.Errorf("Fail update offboard_txn: %v\n", err)
		return err
	}
	return nil
}

func (con *PgSQL) UpdateOffboardTxnByOffboardTxnHash(offboardTxnHash string, handlerTx *sql.Tx) error {
	stmt, err := handlerTx.Prepare(`
		update offboard_txn set status=$1 where offboard_txn_hash=$2
	`)
	if err != nil {
		log.Errorf("Fail update offboard_txn by offboard_txn_hash: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("done", offboardTxnHash)
	if err != nil {
		log.Errorf("Fail update offboard_txn by offboard_txn_hash: %v\n", err)
		return err
	}
	return nil
}

func (con *PgSQL) Offboard(blockHash, txnHash, offboardTxnHash string, nonce, batch int64, handlerTx *sql.Tx) error {
	stmt1, err := handlerTx.Prepare(
		`insert into offboard(
			block_hash, txn_hash, offboard_txn_hash, nonce, batch, status
		)values(
			$1, $2, $3, $4, $5, $6
		)
	`)

	if err != nil {
		log.Errorf("Fail insert offboard: %v\n", err)
		return err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(blockHash, txnHash, offboardTxnHash, nonce, batch, "normal")
	if err != nil {
		log.Errorf("Fail insert offboard: %v\n", err)
		return err
	}

	stmt2, err := handlerTx.Prepare(
		`update offboard_txn set offboard_txn_hash=$1, status=$2, nonce=$3 where block_hash=$4 and txn_hash=$5 and batch=$6 and status<>$7
	`)

	if err != nil {
		log.Errorf("Fail update offboard_txn: %v\n", err)
		return err
	}
	defer stmt1.Close()

	_, err = stmt2.Exec(offboardTxnHash, "pending", nonce, blockHash, txnHash, batch, "timeout")
	if err != nil {
		log.Errorf("Fail update offboard_txn: %v\n", err)
		return err
	}
	return nil

}
