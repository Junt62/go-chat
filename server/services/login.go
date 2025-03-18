package services

import (
	"errors"
	"go-chat/models"
	"go-chat/utils"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (string, error) {
	var user models.User
	if err := models.DB.Where("username=?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
