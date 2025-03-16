package services

import (
	"errors"
	"go-chat/models"
	"go-chat/utils"
)

func Login(username, password string) (string, error) {
	var user models.User
	if err := models.DB.Where("username=?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != password {
		return "", errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
