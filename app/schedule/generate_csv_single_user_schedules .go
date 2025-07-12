package schedule

import (
	"bytes"
	"context"
	"encoding/csv"
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

type GenerateCSVSingleUserSchedules struct {
	scheduleRepo domainSchedule.Repository
}

func NewGenerateCSVSingleUserSchedules(scheduleRepo domainSchedule.Repository) GenerateCSVSingleUserSchedules {
	return GenerateCSVSingleUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h GenerateCSVSingleUserSchedules) Handle(ctx context.Context, userUUID string, query RangeDates) ([]byte, error) {
	header := []string{"ID", "Nama", query.From.Format(time.DateOnly)}
	for i := 1; i <= query.DiffDays(); i++ {
		header = append(header, query.From.Add(time.Hour*time.Duration(i*24)).Format(time.DateOnly))
	}

	content := [][]string{header}

	sc, err := h.scheduleRepo.GetScheduleByUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	wr := csv.NewWriter(&buf)

	row := []string{
		sc.UUID(),
		sc.UserName(),
	}

	if sc.StartDate().Sub(query.From).Microseconds() > 0 {

		for i := 0; i <= query.DiffDays(); i++ {
			row = append(row, "")
		}

		content = append(content, row)
		if err := wr.WriteAll(content); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	for i := 0; i <= query.DiffDays(); i++ {
		_, cycle := generate_schedule_dates(*sc, query.From, time.Hour*time.Duration(i*24))
		row = append(row, cycle)
	}

	content = append(content, row)
	if err := wr.WriteAll(content); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
