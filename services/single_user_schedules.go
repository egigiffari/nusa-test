package services

import (
	"context"
	"time"

	"github.com/egigiffari/nusa-test/domain/schedule"
)

type singleUserSchedules struct {
	scheduleRepo schedule.Repository
}

func NewSingleUserSchedules(scheduleRepo schedule.Repository) singleUserSchedules {
	return singleUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h singleUserSchedules) Handle(ctx context.Context, userUUID string, query RangeDates) (*UserSchedule, error) {
	sc, err := h.scheduleRepo.GetScheduleByUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	sch := UserSchedule{
		UserUUID:  sc.UserUUID(),
		UserName:  sc.UserName(),
		Schedules: make(map[string]string, 0),
	}

	if sc.StartDate().Sub(query.From).Microseconds() < 0 {
		return &sch, nil
	}

	for i := 0; i <= query.DiffDays(); i++ {
		date, cycle := generate_schedule_dates(*sc, query.From, time.Hour*time.Duration(i*24))
		sch.Schedules[date] = cycle
	}

	return &sch, nil
}
