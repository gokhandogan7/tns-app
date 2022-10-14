package main

import (
	"api/config"
	router "api/routers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
		defer db.Close()
	}
	fmt.Println("success")

	router.HandleRequests()

}
