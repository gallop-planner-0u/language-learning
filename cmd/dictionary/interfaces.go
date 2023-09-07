package dictionary

import (
	"context"
	"language-learning/ent"
)

type Repository interface {
	GetByID(ctx context.Context, userID, id int) (*ent.Record, error)
	GetByWord(ctx context.Context, userID int, word string) (*ent.Record, error)
	Create(ctx context.Context, userID int, record *ent.Record) (int, error)
	Update(ctx context.Context, userID, id int, updatedRecord *ent.Record) error
	Delete(ctx context.Context, userID, id int) error
}

