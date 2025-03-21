package middleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(webAddress string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     strings.Split(webAddress, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}
