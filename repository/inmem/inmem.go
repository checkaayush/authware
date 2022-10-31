package inmem

import (
	"github.com/checkaayush/authware/repository"
)

type inMemRepo struct{}

func NewInMemRepository() repository.Repository {
	return new(inMemRepo)
}
