package user_test

import (
	"context"
	"testing"

	"github.com/egigiffari/nusa-test/adapters/user"
	domainUser "github.com/egigiffari/nusa-test/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_GetUser(t *testing.T) {
	type testCase struct {
		name        string
		uuid        string
		expectedErr error
	}

	usr, err := domainUser.NewUser(uuid.NewString(), "Arya")
	require.NoError(t, err)

	init := make(map[string]domainUser.User)
	init[usr.UUID()] = *usr

	repo := user.NewMemory(init)

	testCases := []testCase{
		{
			name:        "No User By UUID",
			uuid:        uuid.NewString(),
			expectedErr: domainUser.ErrUserNotFound,
		},
		{
			name:        "User By UUID",
			uuid:        usr.UUID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetUser(context.Background(), tc.uuid)
			assert.Equal(t, err, tc.expectedErr)
		})
	}
}

func TestMemory_AddUser(t *testing.T) {
	type testCase struct {
		name        string
		usr         string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add User",
			usr:         "Arya",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := user.NewMemory(make(map[string]domainUser.User))

			usr, err := domainUser.NewUser(uuid.NewString(), tc.usr)
			require.NoError(t, err)

			err = repo.AddUser(context.Background(), usr)
			assert.Equal(t, err, tc.expectedErr)
		})
	}
}
