package utils

import (
	"go-api-books/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret_key_here")

func GenerateToken(email string) (string, error) {
	claims := models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
