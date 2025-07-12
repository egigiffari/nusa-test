package app

import (
	"context"

	"github.com/egigiffari/nusa-test/adapters/schedule"
	appSchedule "github.com/egigiffari/nusa-test/app/schedule"
	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type Application struct {
	AllUserSchedules               appSchedule.AllUserSchedules
	SingleUserSchedules            appSchedule.SingleUserSchedules
	CheckUserSchedule              appSchedule.CheckUserSchedule
	GenerateCSVAllUserSchedules    appSchedule.GenerateCSVAllUserSchedules
	GenerateCSVSingleUserSchedules appSchedule.GenerateCSVSingleUserSchedules
}

func NewApplication(ctx context.Context) Application {

	// userRepo := user.NewMemory(make(map[string]domainUser.User, 0))
	scheduleRepo := schedule.NewMemory(make(map[string]domainSchedule.Schedule, 0))

	return Application{
		AllUserSchedules:               appSchedule.NewAllUserSchedules(scheduleRepo),
		SingleUserSchedules:            appSchedule.NewSingleUserSchedules(scheduleRepo),
		CheckUserSchedule:              appSchedule.NewCheckUserSchedule(scheduleRepo),
		GenerateCSVAllUserSchedules:    appSchedule.NewGenerateCSVAllUserSchedules(scheduleRepo),
		GenerateCSVSingleUserSchedules: appSchedule.NewGenerateCSVSingleUserSchedules(scheduleRepo),
	}
}
