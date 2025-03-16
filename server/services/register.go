package services

import (
	"errors"
	"go-chat/models"
)

func Register(username, password, email string) error {
	existingUser := models.FindUserByEmail(email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	user := models.User{
		Username: username,
		Password: models.HashPassword(password),
		Email:    email,
	}

	err := models.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
