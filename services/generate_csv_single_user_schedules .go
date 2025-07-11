package services

import (
	"context"
	"encoding/csv"
	"io"
	"time"

	"github.com/egigiffari/nusa-test/domain/schedule"
)

type generateSingleUserSchedules struct {
	scheduleRepo schedule.Repository
}

func NewGenerateSingleUserSchedules(scheduleRepo schedule.Repository) generateSingleUserSchedules {
	return generateSingleUserSchedules{
		scheduleRepo: scheduleRepo,
	}
}

func (h generateSingleUserSchedules) Handle(ctx context.Context, userUUID string, query RangeDates, writer io.Writer) error {
	header := []string{"ID", "Nama", query.From.Format(time.DateOnly)}
	for i := 1; i <= query.DiffDays(); i++ {
		header = append(header, query.From.Add(time.Hour*time.Duration(i*24)).Format(time.DateOnly))
	}

	content := [][]string{header}

	sc, err := h.scheduleRepo.GetScheduleByUser(ctx, userUUID)
	if err != nil {
		return err
	}

	wr := csv.NewWriter(writer)

	row := []string{
		sc.UUID(),
		sc.UserName(),
	}

	if sc.StartDate().Sub(query.From).Microseconds() < 0 {

		for i := 0; i <= query.DiffDays(); i++ {
			row = append(row, "")
		}

		content = append(content, row)
		return wr.WriteAll(content)
	}

	for i := 0; i <= query.DiffDays(); i++ {
		_, cycle := generate_schedule_dates(*sc, query.From, time.Hour*time.Duration(i*24))
		row = append(row, cycle)
	}

	content = append(content, row)
	return wr.WriteAll(content)
}
