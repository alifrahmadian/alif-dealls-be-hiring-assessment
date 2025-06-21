package errors

import "errors"

var (
	ErrStartDateRequired        = errors.New("start date is required")
	ErrEndDateRequired          = errors.New("end date is required")
	ErrInvalidStartDate         = errors.New("invalid start date format")
	ErrInvalidEndDate           = errors.New("invalid end date format")
	ErrAttendancePeriodNotFound = errors.New("attendance period not found")
)
