package schedule_test

import (
	"context"
	"testing"
	"time"

	"github.com/egigiffari/nusa-test/adapters/schedule"
	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_GetScheduleByUser(t *testing.T) {
	type testCase struct {
		name        string
		uuid        string
		expectedErr error
	}

	sch, err := domainSchedule.NewSchedule(
		uuid.NewString(),
		uuid.NewString(),
		"Arya",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"P", "P", "S", "S", "M", "M", "L"},
	)

	require.NoError(t, err)

	init := make(map[string]domainSchedule.Schedule)
	init[sch.UUID()] = *sch

	repo := schedule.NewMemory(init)

	testCases := []testCase{
		{
			name:        "No Schedule By User UUID",
			uuid:        uuid.NewString(),
			expectedErr: domainSchedule.ErrScheduleNotFound,
		},
		{
			name:        "Schedule By User UUID",
			uuid:        sch.UserUUID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetScheduleByUser(context.Background(), tc.uuid)
			assert.Equal(t, err, tc.expectedErr)
		})
	}
}

func TestMemory_AddSchedule(t *testing.T) {
	type testCase struct {
		name        string
		usr         string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Schedule",
			usr:         "Arya",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := schedule.NewMemory(make(map[string]domainSchedule.Schedule))

			sch, err := domainSchedule.NewSchedule(
				uuid.NewString(),
				uuid.NewString(),
				tc.usr,
				time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
				[]string{"P", "P", "S", "S", "M", "M", "L"},
			)
			require.NoError(t, err)

			err = repo.AddSchedule(context.Background(), sch)
			assert.Equal(t, err, tc.expectedErr)
		})
	}
}

func TestMemory_GetAllSchedules(t *testing.T) {
	type testCase struct {
		name   string
		from   time.Time
		length int
	}

	sch, err := domainSchedule.NewSchedule(
		uuid.NewString(),
		uuid.NewString(),
		"Arya",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"P", "P", "S", "S", "M", "M", "L"},
	)

	require.NoError(t, err)

	init := make(map[string]domainSchedule.Schedule)
	init[sch.UUID()] = *sch

	repo := schedule.NewMemory(init)

	testCases := []testCase{
		{
			name:   "Get Empty Schedules",
			from:   time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
			length: 0,
		},
		{
			name:   "Get Schedules",
			from:   time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
			length: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schedules := repo.GetAllSchedules(context.Background(), tc.from)
			assert.Equal(t, len(schedules), tc.length)
		})
	}
}
