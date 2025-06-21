package errors

import "errors"

var (
	ErrPayrollHasBeenProcessed           = errors.New("payroll has been proceesed")
	ErrPayrollUserIDRequired             = errors.New("user id is required")
	ErrPayrollAttendancePeriodIDRequired = errors.New("attendance period id is required")
)
