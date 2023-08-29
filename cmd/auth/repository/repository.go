package repository

import (
	"context"
)

type Repository struct {
	client interface{}
}

func New(client interface{}) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Get(ctx context.Context, username, passwordHash string) (interface{}, error)
