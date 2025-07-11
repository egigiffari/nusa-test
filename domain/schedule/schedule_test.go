package schedule_test

import (
	"testing"
	"time"

	"github.com/egigiffari/nusa-test/domain/schedule"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSchedule(t *testing.T) {
	t.Parallel()

	id := uuid.NewString()
	userUUID := uuid.NewString()
	userName := "Windah"
	startDate := time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC)
	cycles := []string{"P", "P", "S", "S", "M", "M", "L"}

	sch, err := schedule.NewSchedule(id, userUUID, userName, startDate, cycles)
	require.NoError(t, err)

	assert.Equal(t, id, sch.UUID())
	assert.Equal(t, userUUID, sch.UserUUID())
	assert.Equal(t, userName, sch.UserName())
	assert.Equal(t, startDate, sch.StartDate())
	assert.Equal(t, cycles, sch.Cycles())
}
