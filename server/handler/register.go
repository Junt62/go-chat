package handler

import (
	"go-chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var userReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if userReq.Username == "" || userReq.Password == "" || userReq.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username, password, and email are required"})
		return
	}

	err := services.Register(userReq.Username, userReq.Password, userReq.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
