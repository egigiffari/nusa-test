package user

import (
	"context"
	"errors"
)

type Repository interface {
	AddUser(ctx context.Context, usr *User) error
	GetUser(ctx context.Context, uuid string) (*User, error)
}

var (
	ErrUserNotFound    = errors.New("the user was not found")
	ErrFailedToAddUser = errors.New("failed to add user")
	ErrUpdateUser      = errors.New("failed to update user")
)
