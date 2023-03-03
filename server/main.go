package main

import (
	"tsi.co/go-api2/database"
	"tsi.co/go-api2/resources/models"
	"tsi.co/go-api2/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&models.Actor{})

	server.Init()
}
