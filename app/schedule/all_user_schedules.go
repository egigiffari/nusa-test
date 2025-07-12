package schedule

import (
	"context"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type RangeDates struct {
	From time.Time
	To   time.Time
}

func (r RangeDates) DiffDays() int {
	return int(r.To.Sub(r.From).Hours() / 24)
}

type AllUserSchedules struct {
	scheduleRepo domainSchedule.Repository
}

func NewAllUserSchedules(scheduleRepo domainSchedule.Repository) AllUserSchedules {
	return AllUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h AllUserSchedules) Handle(ctx context.Context, query RangeDates) []UserSchedule {
	schedules := h.scheduleRepo.GetAllSchedules(ctx, query.From)

	userSchedules := make([]UserSchedule, 0)
	for _, s := range schedules {

		userSchedule := UserSchedule{
			UserUUID: s.UserUUID(),
			UserName: s.UserName(),
		}

		for i := 0; i <= query.DiffDays(); i++ {
			date, cycle := generate_schedule_dates(s, query.From, time.Hour*time.Duration(i*24))
			userSchedule.Schedules[date] = cycle
		}

		userSchedules = append(userSchedules, userSchedule)
	}

	return userSchedules
}
