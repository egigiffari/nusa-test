package schedule

import (
	"context"
	"sync"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type Memory struct {
	schedules map[string]domainSchedule.Schedule
	sync.Mutex
}

func NewMemory(initial map[string]domainSchedule.Schedule) *Memory {
	return &Memory{
		schedules: initial,
	}
}

func (repo *Memory) AddSchedule(ctx context.Context, sch *domainSchedule.Schedule) error {
	if repo.schedules == nil {
		repo.Lock()
		repo.schedules = make(map[string]domainSchedule.Schedule)
		repo.Unlock()
	}

	if _, ok := repo.schedules[sch.UUID()]; ok {
		return domainSchedule.ErrFailedToAddSchedule
	}

	repo.Lock()
	repo.schedules[sch.UUID()] = *sch
	repo.Unlock()
	return nil
}

func (repo *Memory) GetScheduleByUser(ctx context.Context, userUUID string) (*domainSchedule.Schedule, error) {
	repo.Lock()

	for _, s := range repo.schedules {
		if s.UserUUID() == userUUID {
			return &s, nil
		}
	}

	repo.Unlock()

	return nil, domainSchedule.ErrScheduleNotFound
}

func (repo *Memory) GetAllSchedules(ctx context.Context, from time.Time) []domainSchedule.Schedule {
	repo.Lock()

	schedules := make([]domainSchedule.Schedule, 0)
	for _, sc := range repo.schedules {
		if sc.StartDate().Sub(from).Milliseconds() >= 0 {
			schedules = append(schedules, sc)
		}
	}

	repo.Unlock()
	return schedules
}
