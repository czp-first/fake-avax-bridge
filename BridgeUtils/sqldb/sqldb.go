package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	PostgresDriver          = "postgres"
	PostgresFmt             = "postgresql://%s:%s@%s/%s?sslmode=disable" // postgresql://username:password@host:port/database?sslmode=disable
	DefaultPostgresPoolSize = 20

	transactionalMaxRetry   = 10
	transactionalRetryDelay = 10 * time.Millisecond
)

var (
	ErrTxConflict = errors.New("Transaction conflict")
	ErrNoRows     = errors.New("No matching rows in the database")
)

type Db struct {
	driver string  // database driver
	info   string  // database connection info
	db     *sql.DB // database accesss object
}

type DbTx struct {
	db *Db
	tx *sql.Tx
}

type SqlStorage interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TxFunc func(tx *DbTx, args ...interface{}) error

func NewDb(driver, info string, poolSize int) (*Db, error) {

	if driver == "sqlite3" {
		if ok, err := exists(info); err != nil {
			return nil, err
		} else if !ok {
			dir := filepath.Dir(info)
			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	db, err := sql.Open(driver, info)
	if err != nil {
		return nil, err
	}

	if driver == "postgres" {
		err = db.Ping()
		if err != nil {
			db.Close()
			return nil, err
		}

		if poolSize > 0 {
			db.SetMaxIdleConns(poolSize)
			db.SetMaxOpenConns(poolSize)
		}
	} else if driver == "sqlite3" {
		db.SetMaxIdleConns(1)
	}

	d := &Db{
		driver: driver,
		info:   info,
		db:     db,
	}
	return d, nil
}

func (d *Db) Close() {
	if d.db != nil {
		d.db.Close()
		d.db = nil
		d.driver = ""
		d.info = ""
	}
}

func (d *Db) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d *Db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d *Db) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d *Db) OpenTransaction() (*DbTx, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	t := &DbTx{
		db: d,
		tx: tx,
	}
	return t, nil
}

func (t *DbTx) Discard() {
	if t.tx != nil {
		t.tx.Rollback()
		t.tx = nil
	}
}

func (t *DbTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *DbTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.tx.Query(query, args...)
}

func (t *DbTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRow(query, args...)
}

func (t *DbTx) ConvertError(err error) error {
	if err == nil {
		return nil
	}

	var patterns []string
	if t.db.driver == "postgres" {
		patterns = []string{"retry transaction", "restart transaction",
			"current transaction is aborted", "40001", "cr000"}
	} else {
		patterns = []string{"database is locked"}
	}

	errMsg := strings.ToLower(err.Error())
	for _, pat := range patterns {
		if strings.Contains(errMsg, pat) {
			return ErrTxConflict
		}
	}
	return err
}

func (t *DbTx) Commit() error {
	err := t.tx.Commit()
	if err == nil {
		t.tx = nil
		return nil
	}
	return t.ConvertError(err)
}

func (d *Db) Transactional(callback TxFunc, args ...interface{}) error {
	for i := 0; i < transactionalMaxRetry; i++ {
		t, err := d.OpenTransaction()
		if err != nil {
			return err
		}

		err = callback(t, args...)
		if err == nil {
			err = t.Commit()
			if err == nil {
				return nil
			}
		}

		err = t.ConvertError(err)
		t.Discard()
		if err != ErrTxConflict {
			return err
		}

		time.Sleep(transactionalRetryDelay)
	}

	return fmt.Errorf("%d Tx commit retries", transactionalMaxRetry)
}

func exists(fpath string) (bool, error) {
	_, err := os.Stat(fpath)
	if err == nil || os.IsExist(err) {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ChkQueryRow(err error) (bool, error) {
	found := false
	if err == nil {
		found = true
	} else if err == sql.ErrNoRows {
		err = nil
	}
	return found, err
}

func ChkExec(res sql.Result, err error, want int64, info string) error {
	if err != nil {
		return err
	}

	got, err := res.RowsAffected()
	if err == nil && got != want {
		err = fmt.Errorf("%s: invalid SQL #rows: %d != %d", info, got, want)
		if got == 0 {
			err = fmt.Errorf("%s: %s", err, ErrNoRows)
		}
	}
	return err
}

func ChkExecDiffError(res sql.Result, err, diffErr error, want int64) error {
	if err != nil {
		return err
	}

	got, err := res.RowsAffected()
	if err == nil && got != want {
		err = diffErr
	}
	return err
}

func InClause(column string, num, start int) string {
	if column == "" || num < 1 || start < 1 {
		return ""
	}

	params := make([]string, num)
	for i := 0; i < num; i++ {
		params[i] = fmt.Sprintf("$%d", start+1)
	}
	return fmt.Sprintf("%s IN (%s)", column, strings.Join(params, ", "))
}
