package main

import (
	"github.com/iurylemos/webapi-golang/database"
	"github.com/iurylemos/webapi-golang/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}