package schedule

import (
	"context"
	"errors"
)

type Repository interface {
	AddSchedule(ctx context.Context, sch *Schedule) error
	GetSchedule(ctx context.Context, userUUID string) (*Schedule, error)
}

var (
	ErrFailedToAddSchedule = errors.New("failed to add schedule")
	ErrScheduleNotFound    = errors.New("schedule was not found")
)
