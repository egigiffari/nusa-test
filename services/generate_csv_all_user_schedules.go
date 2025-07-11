package services

import (
	"context"
	"encoding/csv"
	"io"
	"time"

	"github.com/egigiffari/nusa-test/domain/schedule"
)

type generateAllCSVUserSchedules struct {
	scheduleRepo schedule.Repository
}

func NewGenerateAllCSVUserSchedules(scheduleRepo schedule.Repository) generateAllCSVUserSchedules {
	return generateAllCSVUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h generateAllCSVUserSchedules) Handle(ctx context.Context, query RangeDates, writer io.Writer) error {
	header := []string{"ID", "Nama", query.From.Format(time.DateOnly)}
	for i := 1; i <= query.DiffDays(); i++ {
		header = append(header, query.From.Add(time.Hour*time.Duration(i*24)).Format(time.DateOnly))
	}

	schedules := h.scheduleRepo.GetAllSchedules(ctx, query.From)

	body := [][]string{}
	for _, s := range schedules {

		row := []string{
			s.UUID(),
			s.UserName(),
		}

		for i := 0; i <= query.DiffDays(); i++ {
			_, cycle := generate_schedule_dates(s, query.From, time.Hour*time.Duration(i*24))
			row = append(row, cycle)
		}

		body = append(body, row)
	}

	content := [][]string{header}
	content = append(content, body...)

	wr := csv.NewWriter(writer)
	return wr.WriteAll(content)
}
