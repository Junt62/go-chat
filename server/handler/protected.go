package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
}
