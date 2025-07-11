package schedule

import (
	"time"

	"github.com/egigiffari/nusa-test/common/errors"
)

type Schedule struct {
	uuid string

	userUUID string
	userName string

	startDate time.Time
	cycles    []string
}

func NewSchedule(
	uuid string,
	userUUID string,
	userName string,
	startDate time.Time,
	cycles []string,
) (*Schedule, error) {

	if uuid == "" {
		return nil, ErrInvalidUUID
	}

	if userUUID == "" {
		return nil, ErrInvalidUserUUID
	}

	if len([]rune(userName)) <= 2 {
		return nil, ErrUserNameTooShort
	}

	if startDate.IsZero() {
		return nil, ErrStartDateIsZero
	}

	if len(cycles)%7 != 0 {
		return nil, ErrInvalidCycles
	}

	return &Schedule{
		uuid:      uuid,
		userUUID:  userUUID,
		userName:  userName,
		startDate: startDate,
		cycles:    cycles,
	}, nil
}

func (s Schedule) UUID() string {
	return s.uuid
}

func (s Schedule) UserUUID() string {
	return s.userUUID
}

func (s Schedule) UserName() string {
	return s.userName
}

func (s Schedule) StartDate() time.Time {
	return s.startDate
}

func (s Schedule) Cycles() []string {
	return s.cycles
}

var (
	ErrInvalidUUID      = errors.NewIncorrectInputError("invalid uuid", "invalid_uuid")
	ErrInvalidUserUUID  = errors.NewIncorrectInputError("invalid user uuid", "invalid_user_uuid")
	ErrUserNameTooShort = errors.NewIncorrectInputError("user name too short", "user_name_too_short")
	ErrStartDateIsZero  = errors.NewIncorrectInputError("zero start date", "zero_start_date")
	ErrInvalidCycles    = errors.NewIncorrectInputError("invalid cycles not in 7 or multiples 7 days", "invalid_cycles_not_in_7_or_multiples_7_days")
)
