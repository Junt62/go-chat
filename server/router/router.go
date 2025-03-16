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

	r.GET("ping", handler.PingHandler)

	r.GET("/user/:id", handler.UserHandler)

	r.GET("/search", handler.SearchHandler)

	r.POST("/register", handler.RegisterHandler)

	r.POST("/login", handler.LoginHandler)

	r.POST("/logout", handler.LogoutHandler)

	r.GET("/protected", middleware.AuthMiddleware(), handler.ProtectedHandler)

	return r
}
