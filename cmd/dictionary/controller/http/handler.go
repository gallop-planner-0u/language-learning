package http

import (
	"language-learning/cmd/dictionary"
)

type Handler struct {
	uc dictionary.Usecase
}

func New(uc dictionary.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

