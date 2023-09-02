package models

import "github.com/golang-jwt/jwt/v5"

type SignUpResponse struct {
	Username string
	Password string
}

type CustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}
