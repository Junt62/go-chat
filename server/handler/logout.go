package handler

import (
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logout successful"})
}
