package user

import (
	"context"
	"sync"

	"github.com/egigiffari/nusa-test/domain/user"
)

type Memory struct {
	users map[string]user.User
	sync.Mutex
}

func NewMemory(initial map[string]user.User) *Memory {
	return &Memory{
		users: initial,
	}
}

func (repo *Memory) AddUser(ctx context.Context, usr *user.User) error {
	if repo.users == nil {
		repo.Lock()
		repo.users = make(map[string]user.User)
		repo.Unlock()
	}

	if _, ok := repo.users[usr.UUID()]; ok {
		return user.ErrFailedToAddUser
	}

	repo.Lock()
	repo.users[usr.UUID()] = *usr
	repo.Unlock()

	return nil
}

func (repo *Memory) GetUser(ctx context.Context, uuid string) (*user.User, error) {

	if user, ok := repo.users[uuid]; ok {
		return &user, nil
	}

	return nil, user.ErrUserNotFound
}
