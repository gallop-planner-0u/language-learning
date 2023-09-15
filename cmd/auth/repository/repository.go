package repository

import (
	"context"
	"language-learning/ent"
	"language-learning/ent/user"

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

func (r *Repository) Get(ctx context.Context, username, passwordHash string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			sql.FieldEQ(user.FieldUsername, username),
			sql.FieldEQ(user.FieldPassword, passwordHash)).
		First(ctx)
	return u, err
}

func (r *Repository) GetById(ctx context.Context, id int) (*ent.User, error) {
	u, err := r.client.User.Get(ctx, id)
	return u, err
}

func (r *Repository) Create(ctx context.Context, user *ent.User) (int, error) {
	u, err := r.client.User.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}
