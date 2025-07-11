package user_test

import (
	"testing"

	"github.com/egigiffari/nusa-test/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	id := uuid.NewString()
	name := "Windah"

	user, err := user.NewUser(id, name)
	require.NoError(t, err)

	assert.Equal(t, id, user.Id())
	assert.Equal(t, name, user.Name())
}
