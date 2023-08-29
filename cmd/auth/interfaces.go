package auth

import "context"

type Repository interface {
	Get(ctx context.Context, username, passwordHash string) (interface{}, error)
	Create(ctx context.Context, user interface{}) (int, error)
}

type Usecase interface {
	SignIn(ctx context.Context, username, password string) (string, error)
	SignUp(ctx context.Context, user interface{}) (int, error)
}
