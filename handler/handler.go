package handler

import (
	"github.com/checkaayush/authware/repository"
)

type Handler struct {
	db repository.Repository
}

func NewHandler(db repository.Repository) *Handler {
	return &Handler{
		db: db,
	}
}
