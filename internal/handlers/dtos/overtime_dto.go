package dtos

type CreateOvertimeRequest struct {
	Date string `json:"date" binding:"required"`
	OvertimeHours uint64 `json:"overtime_hours" binding:"required"`
}

type CreateOvertimeResponse struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Date string `json:"date"`
	OvertimeHours uint64 `json:"overtime_hours"`
}