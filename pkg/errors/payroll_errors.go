package errors

import "errors"

var (
	ErrPayrollHasBeenProcessed              = errors.New("payroll has been proceesed")
	ErrPayrollUserIDRequired                = errors.New("user id is required")
	ErrPayrollAttendancePeriodIDRequired    = errors.New("attendance period id is required")
	ErrPayrollAttendancePeriodNotStartedYet = errors.New("attendance period has not started yet")
	ErrPayrollNotFound                      = errors.New("payroll not found")
	ErrPayrollInvalidAttendancePeriodID     = errors.New("invalid attendance_period_id")
	ErrPayrollInvalidPayrollID              = errors.New("invalid payroll_id")
	ErrPayrollIDRequired                    = errors.New("payroll_id is required")
)
