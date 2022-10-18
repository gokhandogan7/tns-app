package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1231231234eBike"
	dbName := "firstdatabase"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
