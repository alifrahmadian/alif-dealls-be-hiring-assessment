package models

import "time"

type Overtime struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	User          User      `json:"user"`
	PayrollID     *int64    `json:"payroll_id"`
	Payroll       Payroll   `json:"payroll"`
	Date          time.Time `json:"date"`
	OvertimeHours uint64    `json:"overtime_hours"`
	CreatedBy     int64     `json:"created_by"`
	UpdatedBy     int64     `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
