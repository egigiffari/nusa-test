package schedule

import (
	"context"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type SingleUserSchedules struct {
	scheduleRepo domainSchedule.Repository
}

func NewSingleUserSchedules(scheduleRepo domainSchedule.Repository) SingleUserSchedules {
	return SingleUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h SingleUserSchedules) Handle(ctx context.Context, userUUID string, query RangeDates) (*UserSchedule, error) {
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
