package models

import "time"

type Payroll struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	User User `json:"user"`
	AttendancePeriodID int64 `json:"attendance_period_id"`
	AttendancePeriod AttendancePeriod `json:"attendance_period"`
	BaseSalary uint64 `json:"base_salary"`
	AttendanceDays uint64 `json:"attendance_days"`
	AttendanceAmount uint64 `json:"attendance_amount"`
	OvertimeHours uint64 `json:"overtime_hours"`
	OvertimeAmount uint64 `json:"overtime_amount"`
	ReimbursementAmount uint64 `json:"reimbursement_amount"`
	TotalTakeHomePay uint64 `json:"total_take_home_pay"`
	CreatedBy int64 `json:"created_by"`
	UpdatedBy int64 `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}