package dtos

type CreateAttendanceResponse struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Date string `json:"date"`
}