package handler

import (
	"github.com/checkaayush/authware/rbac"
	"github.com/checkaayush/authware/repository"
)

type Handler struct {
	db   repository.Repository
	auth *rbac.RBACService
}

func NewHandler(db repository.Repository, auth *rbac.RBACService) *Handler {
	return &Handler{
		db:   db,
		auth: auth,
	}
}
