package schedule

import (
	"context"
	"errors"
)

type Repository interface {
	AddSchedule(ctx context.Context, sch *Schedule) error
	GetScheduleByUser(ctx context.Context, userUUID string) (*Schedule, error)
}

var (
	ErrFailedToAddSchedule = errors.New("failed to add schedule")
	ErrScheduleNotFound    = errors.New("schedule was not found")
)
