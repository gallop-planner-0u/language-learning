package usecase

import (
	"context"
	"fmt"
	"language-learning/cmd/auth"
	"language-learning/ent"
	"language-learning/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Usecase struct {
	repo       auth.Repository
	signingKey string
}

func New(repo auth.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) SignUp(ctx context.Context, user *ent.User) (int, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *Usecase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, err := uc.repo.Get(ctx, username, password)
	accessToken, err := uc.GenerateToken(ctx, user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (uc *Usecase) ParseToken(accessToken string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return uc.signingKey, nil
	})

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (uc *Usecase) GenerateToken(ctx context.Context, user *ent.User) (string, error) {
	claims := models.CustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(uc.signingKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
