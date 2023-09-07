package usecase

import (
	"context"
	"language-learning/cmd/dictionary"
	"language-learning/ent"
)

type Usecase struct {
	repo dictionary.Repository
}

func New(repo dictionary.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

