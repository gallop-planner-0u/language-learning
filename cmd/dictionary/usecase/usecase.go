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

func (uc *Usecase) GetByID(ctx context.Context, userID, id int) (*ent.Record, error) {
	return uc.repo.GetByID(ctx, userID, id)
}

func (uc *Usecase) GetByWord(ctx context.Context, userID int, word string) (*ent.Record, error) {
	return uc.repo.GetByWord(ctx, userID, word)
}

func (uc *Usecase) Create(ctx context.Context, userID int, record *ent.Record) (int, error) {
	return uc.repo.Create(ctx, userID, record)
}

func (uc *Usecase) Update(ctx context.Context, userID, id int, updatedRecord *ent.Record) error {
	return uc.repo.Update(ctx, userID, id, updatedRecord)
}

func (uc *Usecase) Delete(ctx context.Context, userID, id int) error {
	return uc.repo.Delete(ctx, userID, id)
}
