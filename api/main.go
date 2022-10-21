package main

import (
	"api/config"
	router "api/routers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := config.GetMySQLDB()
	// if err != nil {
	fmt.Println(db, err)
	// defer db.Close()
	// }

	err = db.Ping()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("success")

	router.HandleRequests()

}
