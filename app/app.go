package app

import (
	"context"
	"time"

	"github.com/egigiffari/nusa-test/adapters/schedule"
	appSchedule "github.com/egigiffari/nusa-test/app/schedule"
	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
	"github.com/google/uuid"
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

	ahmad, _ := domainSchedule.NewSchedule(
		uuid.NewString(),
		"001",
		"Ahmad",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"P", "P", "S", "S", "M", "M", "L"},
	)

	widi, _ := domainSchedule.NewSchedule(
		uuid.NewString(),
		"002",
		"Widi",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"S", "S", "M", "M", "L", "P", "S"},
	)

	yono, _ := domainSchedule.NewSchedule(
		uuid.NewString(),
		"003",
		"Yono",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"M", "M", "P", "L", "P", "P", "M"},
	)

	yohan, _ := domainSchedule.NewSchedule(
		uuid.NewString(),
		"004",
		"Yohan",
		time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
		[]string{"L", "P", "P", "P", "S", "S", "P", "L", "S", "S", "P", "S", "S", "P"},
	)

	schedules := make(map[string]domainSchedule.Schedule)
	schedules[ahmad.UUID()] = *ahmad
	schedules[widi.UUID()] = *widi
	schedules[yono.UUID()] = *yono
	schedules[yohan.UUID()] = *yohan

	scheduleRepo := schedule.NewMemory(schedules)

	return Application{
		AllUserSchedules:               appSchedule.NewAllUserSchedules(scheduleRepo),
		SingleUserSchedules:            appSchedule.NewSingleUserSchedules(scheduleRepo),
		CheckUserSchedule:              appSchedule.NewCheckUserSchedule(scheduleRepo),
		GenerateCSVAllUserSchedules:    appSchedule.NewGenerateCSVAllUserSchedules(scheduleRepo),
		GenerateCSVSingleUserSchedules: appSchedule.NewGenerateCSVSingleUserSchedules(scheduleRepo),
	}
}
