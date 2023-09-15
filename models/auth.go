package models

import "github.com/golang-jwt/jwt/v5"

type SignUpResponse struct {
	Id int
}

type CustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}
