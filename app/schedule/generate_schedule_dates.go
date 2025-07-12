package schedule

import (
	"time"

	domainSchedule "github.com/egigiffari/nusa-test/domain/schedule"
)

func generate_schedule_dates(sch domainSchedule.Schedule, from time.Time, addDuration time.Duration) (string, string) {
	date := from.Add(addDuration)
	diffDaysFromStart := int(date.Sub(sch.StartDate()).Hours() / 24)
	cycle := getShiftCycle(sch.StartDate(), sch.Cycles(), diffDaysFromStart)
	return date.Format(time.DateOnly), cycle
}
