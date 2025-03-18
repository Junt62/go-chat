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

	hashedPassword, err := models.HashPassword(password)
	if err != nil {
		return errors.New("hash password error")
	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	err = models.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
