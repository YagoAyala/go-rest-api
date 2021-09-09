package main

import (
	"github.com/YagoAyala/go-rest-api.git/database"
	"github.com/YagoAyala/go-rest-api.git/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
