package main

import (
	"go-chat/config"
	"go-chat/models"
	"go-chat/router"
)

func main() {
	cfg := config.LoadConfig()

	models.ConnectDatabase(cfg)

	r := router.SetupRouter(cfg)

	r.Run(":8080")
}
