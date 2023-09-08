package repository

import (
	"context"
	"language-learning/ent"
	"language-learning/ent/record"
	"time"

	"entgo.io/ent/dialect/sql"
)

type Repository struct {
	client *ent.Client
}

func New(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetByID(ctx context.Context, userID, id int) (*ent.Record, error) {
	return nil, nil
}

func (r *Repository) GetByWord(ctx context.Context, userID int, word string) (*ent.Record, error) {
	return nil, nil
}

func (r *Repository) Create(ctx context.Context, userID int, newRecord *ent.Record) (int, error) {
	_r, err := r.client.Record.
		Create().
		SetUserID(userID).
		SetWord(newRecord.Word).
		SetTranslation(newRecord.Translation).
		Save(ctx)
	return _r.ID, err
}

func (r *Repository) Update(ctx context.Context, userID, id int, updatedRecord *ent.Record) error {
	err := r.client.Record.Update().Where(
		sql.FieldEQ(record.FieldID, id),
		sql.FieldEQ(record.FieldUserID, userID)).
		SetWord(updatedRecord.Word).
		SetTranslation(updatedRecord.Translation).
		SetUpdatedAt(time.Now()).Exec(ctx)
	return err
}

func (r *Repository) Delete(ctx context.Context, userID, id int) error {
	return nil
}
