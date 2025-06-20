package models

import "time"

type Attendance struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	PayrollID *int64 `json:"payroll_id"`
	Date time.Time `json:"date"`
	CreatedBy int64 `json:"created_by"`
	UpdatedBy int64 `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}