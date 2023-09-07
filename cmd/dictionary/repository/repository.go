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

func (r *Repository) GetByID(ctx context.Context, userID, id int) (*ent.Record, error)
func (r *Repository) GetByWord(ctx context.Context, userID int, word string) (*ent.Record, error)
func (r *Repository) Create(ctx context.Context, userID int, record *ent.Record) (int, error)
func (r *Repository) Update(ctx context.Context, userID, id int, updatedRecord *ent.Record) error
func (r *Repository) Delete(ctx context.Context, userID, id int) error
