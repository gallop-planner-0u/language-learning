package usecase

import (
	"context"
	"language-learning/cmd/auth"
	"language-learning/cmd/auth/repository"
	"language-learning/ent"
	"language-learning/models"
	"language-learning/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	client := pkg.GetEntClient(ctx)
	if assert.NotNil(t, client) {
		t.Fail()
	}
	repo := repository.New(client)
	uc := New(repo)

	tests := []struct {
		name    string
		args    *ent.User
		want    models.SignUpResponse
		wantErr error
	}{
		{
			name: "first sign-up",
			args: &ent.User{
				Username: "username",
				Password: "password",
				Name:     "name",
			},
			want: models.SignUpResponse{
				Id: 1,
			},
			wantErr: auth.ErrDublicateKey,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			_, err := uc.SignUp(ctx, test.args)
			assert.ErrorIs(tt, err, test.wantErr)
			if assert.Error(tt, err) {
				t.Fail()
			}
		})
	}
}
