package services

import (
	"context"
	"time"

	"github.com/egigiffari/nusa-test/domain/schedule"
)

type checkUserSchedule struct {
	scheduleRepo schedule.Repository
}

func NewCheckUserSchedule(scheduleRepo schedule.Repository) checkUserSchedule {
	return checkUserSchedule{
		scheduleRepo: scheduleRepo,
	}
}

func (h checkUserSchedule) Handle(ctx context.Context, userUUID string, from time.Time) (*UserScheduleStatus, error) {
	sc, err := h.scheduleRepo.GetScheduleByUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	sch := UserScheduleStatus{
		UserUUID: sc.UserUUID(),
		UserName: sc.UserName(),
	}

	if sc.StartDate().Sub(from).Microseconds() < 0 {
		return nil, nil
	}

	date, cycle := generate_schedule_dates(*sc, from, 0)
	sch.Date = date
	sch.Cycle = cycle

	return &sch, nil
}
