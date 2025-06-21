package models

import "time"

type Reimbursement struct {
	ID                  int64     `json:"id"`
	UserID              int64     `json:"user_id"`
	User                User      `json:"user"`
	PayrollID           *int64    `json:"payroll_id"`
	Payroll             Payroll   `json:"payroll"`
	ReimbursementAmount uint64    `json:"reimbursement_amount"`
	Date                time.Time `json:"date"`
	Description         string    `json:"description"`
	CreatedBy           int64     `json:"created_by"`
	UpdatedBy           int64     `json:"updated_by"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
