package config

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
)

type Config struct {
	DB_User   string `yaml:"DB_USER"`
	DB_Driver string `yaml:"DB_DRIVER"`
	DB_Name   string `yaml:"DB_NAME"`
	DB_Pass   string `yaml:"DB_PASS"`
}

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func GetMySQLDB() (db *sql.DB, err error) {

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
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
