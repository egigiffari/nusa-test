package schedule

import (
	"context"
	"errors"
	"time"
)

type Repository interface {
	AddSchedule(ctx context.Context, sch *Schedule) error
	GetScheduleByUser(ctx context.Context, userUUID string) (*Schedule, error)
	GetAllSchedules(ctx context.Context, from time.Time) []Schedule
}

var (
	ErrFailedToAddSchedule = errors.New("failed to add schedule")
	ErrScheduleNotFound    = errors.New("schedule was not found")
)
