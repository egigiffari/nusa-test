package schedule

import (
	"context"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type CheckUserSchedule struct {
	scheduleRepo domainSchedule.Repository
}

func NewCheckUserSchedule(scheduleRepo domainSchedule.Repository) CheckUserSchedule {
	return CheckUserSchedule{
		scheduleRepo: scheduleRepo,
	}
}

func (h CheckUserSchedule) Handle(ctx context.Context, userUUID string, from time.Time) (*UserScheduleStatus, error) {
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
