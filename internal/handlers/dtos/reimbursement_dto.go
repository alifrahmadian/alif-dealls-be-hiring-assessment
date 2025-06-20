package dtos

type CreateReimbursementRequest struct {
	ReimbursementAmount uint64 `json:"reimbursement_amount" binding:"required"`
	Date string `json:"date" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateReimbursementResponse struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	ReimbursementAmount uint64 `json:"reimbursement_amount"`
	Date                string`json:"date"`
    Description         string    `json:"description"`
    CreatedBy           int64     `json:"created_by"`
    UpdatedBy           int64     `json:"updated_by"`
    CreatedAt           string`json:"created_at"`
    UpdatedAt           string`json:"updated_at"`
}