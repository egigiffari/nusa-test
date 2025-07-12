package schedule

type UserSchedule struct {
	UserUUID  string            `json:"id"`
	UserName  string            `json:"name"`
	Schedules map[string]string `json:"schedules"`
}

type UserScheduleStatus struct {
	UserUUID string `json:"id"`
	UserName string `json:"name"`
	Date     string `json:"date"`
	Cycle    string `json:"shift"`
}
