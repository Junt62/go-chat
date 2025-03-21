package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &jwt.StandardClaims{
		Subject:   strconv.FormatUint(uint64(userID), 10),
		ExpiresAt: expirationTime,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.StandardClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func generateSecretKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
