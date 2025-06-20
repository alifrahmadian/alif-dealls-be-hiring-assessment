package errors

import "errors"

var (
	ErrOvertimeIsTakenToday = errors.New("you have submitted overtime for this day")
	ErrOvertimeMoreThanThreeHours = errors.New("overtime should not be more than 3 hours per day")
	ErrEmployeeHasNotWorkedToday = errors.New("you haven't worked for today")
	ErrOvertimeDateIsRequired = errors.New("date is required")
	ErrOvertimeHoursRequired = errors.New("overtime hour is required")
	ErrInvalidOvertimeDate = errors.New("invalid date format")
)