package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type ReimbursementRepository interface {
	CreateReimbursement(reimbursement *models.Reimbursement) (*models.Reimbursement, error)
}

type reimbursementRepository struct {
	DB *sql.DB
}

func NewReimbursementRepository(db *sql.DB) ReimbursementRepository {
	return &reimbursementRepository{
		DB: db,
	}
}

func (r *reimbursementRepository) CreateReimbursement(reimbursement *models.Reimbursement) (*models.Reimbursement, error) {
	query := `
		INSERT INTO reimbursements(user_id, reimbursement_amount, date, description, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		reimbursement.UserID,
		reimbursement.ReimbursementAmount,
		reimbursement.Date,
		reimbursement.Description,
		reimbursement.CreatedBy,
		reimbursement.UpdatedBy,
		reimbursement.CreatedAt,
		reimbursement.UpdatedAt,
	).Scan(&reimbursement.ID)

	if err != nil {
		return nil, err
	}

	return reimbursement, nil
}