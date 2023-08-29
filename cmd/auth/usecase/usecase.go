package usecase

import (
	"context"
	"language-learning/cmd/auth"
	"language-learning/ent"
)

type Usecase struct {
	repo auth.Repository
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

func (uc *Usecase) ParseToken(accessToken string) (*ent.User, error) {
	return nil, nil
}

func (uc *Usecase) GenerateToken(ctx context.Context, user *ent.User) (string, error) {
	var accessToken string
	var err error
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
