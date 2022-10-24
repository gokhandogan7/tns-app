package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
)

type Config struct {
	DB_User   string `yaml:"MYSQL_USER"`
	DB_Driver string `yaml:"MYSQL_DRIVER"`
	DB_Name   string `yaml:"MYSQL_DATABASE"`
	DB_Pass   string `yaml:"MYSQL_ROOT_PASSWORD"`
	DB_Host   string `yaml:"MYSQL_HOST"`
	DB_Port   string `yaml:"MYSQL_PORT"`
}

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func GetMySQLDB() (db *sql.DB, err error) {

	/* limit := 10
	// check that the database is reachable; try at least 3 times to connect
	for i := 0; i <= limit; i++ {
		err := db.Ping()
		if err != nil && i == limit {
			fmt.Errorf("couldn't connect to database after %d tries: %s", i, err)
			break
		} else if err != nil {
			log.Info("Couldn't connect to database, retrying in 1 second ...")
			time.Sleep(5 * time.Second)
		} else {
			log.Info("Successfully connected to database")
			break
		}
	} */

	confContent, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}

	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}

	dbDriver := conf.DB_Driver
	dbUser := conf.DB_User
	dbPass := conf.DB_Pass
	dbName := conf.DB_Name
	dbHost := conf.DB_Host
	dbPort := conf.DB_Port
	fmt.Println("Connecting to database...")

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db, err
}
