package schedule

type UserSchedule struct {
	UserUUID  string
	UserName  string
	Schedules map[string]string
}

type UserScheduleStatus struct {
	UserUUID string
	UserName string
	Date     string
	Cycle    string
}
