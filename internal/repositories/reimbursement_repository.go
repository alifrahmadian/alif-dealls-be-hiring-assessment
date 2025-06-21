package repositories

import (
	"database/sql"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type ReimbursementRepository interface {
	CreateReimbursement(reimbursement *models.Reimbursement) (*models.Reimbursement, error)
	SumReimbursementAmountPerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error)
	InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error
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

func (r *reimbursementRepository) SumReimbursementAmountPerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error) {
	var sum uint64

	query := `
		SELECT COALESCE(SUM(reimbursement_amount), 0) FROM reimbursements
		WHERE user_id = $1
		AND payroll_id IS NULL
		AND date::date BETWEEN $2::date AND $3::date;
	`

	err := r.DB.QueryRow(
		query,
		userID,
		startDate,
		endDate,
	).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (r *reimbursementRepository) InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error {
	query := `
		UPDATE reimbursements
		SET payroll_id = $1, updated_at = NOW()
		WHERE user_id = $2 
		AND date::date BETWEEN $3::date AND $4::date;
	`
	_, err := tx.Exec(query, payrollID, userID, startDate, endDate)
	if err != nil {
		return err
	}

	return nil
}
