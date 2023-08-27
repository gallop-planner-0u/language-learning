package auth

import "context"

type Repository interface {
	Get(ctx context.Context, username, passwordHash string) (interface{}, error)
	Create(ctx context.Context, user interface{}) (int, error)
}
