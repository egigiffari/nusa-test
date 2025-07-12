package schedule

import (
	"bytes"
	"context"
	"encoding/csv"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type GenerateCSVAllUserSchedules struct {
	scheduleRepo domainSchedule.Repository
}

func NewGenerateCSVAllUserSchedules(scheduleRepo domainSchedule.Repository) GenerateCSVAllUserSchedules {
	return GenerateCSVAllUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h GenerateCSVAllUserSchedules) Handle(ctx context.Context, query RangeDates) ([]byte, error) {
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

	var buf bytes.Buffer
	wr := csv.NewWriter(&buf)
	if err := wr.WriteAll(content); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
