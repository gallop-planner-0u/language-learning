package repository

import (
	"context"
	"language-learning/ent"
)

type Repository struct {
	client *ent.Client
}

func New(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Get(ctx context.Context, username, passwordHash string) (*ent.User, error) {
	return nil, nil
}

func (r *Repository) Create(ctx context.Context, user *ent.User) (int, error) {
	return 0, nil
}
