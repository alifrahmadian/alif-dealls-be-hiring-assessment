package errors

import "errors"

var (
	ErrEmployeeHasRecordAttendance = errors.New("employee has recorded the attendance today")
	ErrIsWeekend = errors.New("you can't check-in during the weekend")
)