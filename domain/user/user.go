package user

import "github.com/egigiffari/nusa-test/common/errors"

type User struct {
	uuid string
	name string
}

func NewUser(uuid string, name string) (*User, error) {
	if len([]rune(uuid)) <= 0 {
		return nil, ErrInvalidUUID
	}

	if len([]rune(name)) <= 2 {
		return nil, ErrNameTooShort
	}

	return &User{
		uuid: uuid,
		name: name,
	}, nil
}

func (u User) UUID() string {
	return u.uuid
}

func (u User) Name() string {
	return u.name
}

var (
	ErrInvalidUUID  = errors.NewIncorrectInputError("id must be valid", "id_must_be_valid")
	ErrNameTooShort = errors.NewIncorrectInputError("name is too short", "name_is_too_short")
)
