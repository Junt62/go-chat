package router

import (
	"go-chat/config"
	"go-chat/handler"
	"go-chat/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware(cfg.WebAddress))

	// r.Use(middleware.LoggerMiddleware())

	r.GET("/api/ping", middleware.AuthMiddleware(), handler.PingHandler)

	r.GET("/api/user/:id", handler.UserHandler)

	r.GET("/api/search", handler.SearchHandler)

	r.POST("/api/register", handler.RegisterHandler)

	r.POST("/api/login", handler.LoginHandler)

	r.POST("/api/logout", handler.LogoutHandler)

	r.GET("/api/protected", middleware.AuthMiddleware(), handler.ProtectedHandler)

	return r
}
