package main

import (
	"api/config"
	router "api/routers"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {

	db, err := config.GetMySQLDB()
	// if err != nil {
	fmt.Println(db, err)
	// defer db.Close()
	// }
	limit := 15
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
	}

	/* err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} */
	//fmt.Println("success")

	router.HandleRequests()

}
