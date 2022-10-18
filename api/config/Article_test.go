package config

import (
	"database/sql"
	"io/ioutil"
	"log"
	"testing"

	"github.com/go-yaml/yaml"
)

var testQueries *Queries

func TestGetMySQLDB(t *testing.T) {
	confContent, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}

	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}

	conn, err := sql.Open(conf.DB_Driver, conf.DB_User+":"+conf.DB_Pass+"@/"+conf.DB_Name)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
}
