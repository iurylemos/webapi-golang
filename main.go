package main

import "github.com/iurylemos/webapi-golang/server"

func main() {
	server := server.NewServer()

	server.Run()
}