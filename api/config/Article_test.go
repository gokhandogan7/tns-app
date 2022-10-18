package config

import (
	"database/sql"
	"log"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = "root:1231231234eBike@/firstdatabase"
)

var testQueries *Queries

func TestGetMySQLDB(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
}
