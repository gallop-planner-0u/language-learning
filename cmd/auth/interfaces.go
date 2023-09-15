package auth

import (
	"context"
	"errors"
	"language-learning/ent"
	"language-learning/models"
)

var (
	ErrDublicateKey = errors.New("pq: ")
)

type Repository interface {
	Get(ctx context.Context, username, passwordHash string) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) (int, error)
}

type Usecase interface {
	SignIn(ctx context.Context, username, password string) (string, error)
	SignUp(ctx context.Context, user *ent.User) (int, error)
	ParseToken(accessToken string) (*models.CustomClaims, error)
}
