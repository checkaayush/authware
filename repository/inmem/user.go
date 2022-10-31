package inmem

import (
	"context"
	"sync"

	"github.com/checkaayush/authware/model"
)

var (
	users = map[int]*model.User{
		1: {
			ID:    1,
			Name:  "Aayush",
			Email: "aayush@authware.io",
			Role:  "member",
		},
	}
	userSeq  = 2
	userLock = sync.Mutex{}
)

func (i *inMemRepo) ListUsers(ctx context.Context) ([]model.User, error) {
	userLock.Lock()
	defer userLock.Unlock()

	us := []model.User{}
	for _, user := range users {
		us = append(us, *user)
	}
	return us, nil
}

func (i *inMemRepo) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	userLock.Lock()
	defer userLock.Unlock()

	u.ID = userSeq
	users[u.ID] = u
	userSeq++
	return u, nil
}

func (i *inMemRepo) DeleteUser(ctx context.Context, id int) error {
	userLock.Lock()
	defer userLock.Unlock()

	delete(users, id)
	return nil
}
