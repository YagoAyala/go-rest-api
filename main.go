package main

import (
	"github.com/YagoAyala/go-rest-api.git/src/database"
	"github.com/YagoAyala/go-rest-api.git/src/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
