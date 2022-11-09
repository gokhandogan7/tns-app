package main

import (
	"api/config"
	router "api/routers"
)

func main() {

	config.ConnectMySQLDB()
	router.HandleRequests()

}
