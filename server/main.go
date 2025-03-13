package main

import (
	"go-chat/config"
	"go-chat/models"
	"go-chat/server"
)

func main() {
	cfg := config.LoadConfig()

	models.ConnectDatabase(cfg)

	server.Start()
}
