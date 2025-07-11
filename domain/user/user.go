package user

import "github.com/egigiffari/nusa-test/common/errors"

type User struct {
	id   string
	name string
}

func NewUser(id string, name string) (*User, error) {
	if len([]rune(id)) <= 0 {
		return nil, ErrInvalidUUID
	}

	if len([]rune(name)) <= 2 {
		return nil, ErrNameTooShort
	}

	return &User{
		id:   id,
		name: name,
	}, nil
}

func (u User) Id() string {
	return u.id
}

func (u User) Name() string {
	return u.name
}

var (
	ErrInvalidUUID  = errors.NewIncorrectInputError("id must be valid", "id_must_be_valid")
	ErrNameTooShort = errors.NewIncorrectInputError("name is too short", "name_is_too_short")
)
