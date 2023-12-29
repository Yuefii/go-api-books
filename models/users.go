package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
