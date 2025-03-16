package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret")

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &jwt.StandardClaims{
		Subject:   string(rune(userID)),
		ExpiresAt: expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
