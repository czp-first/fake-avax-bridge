package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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
