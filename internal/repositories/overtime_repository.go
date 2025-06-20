package repositories

import (
	"database/sql"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type OvertimeRepository interface {
	CreateOvertime(overtime *models.Overtime) (*models.Overtime, error)
	CheckIfOvertimeHasBeenTaken(date time.Time) (bool, error)
}

type overtimeRepository struct {
	DB *sql.DB
}

func NewOvertimeRepository(db *sql.DB) OvertimeRepository {
	return &overtimeRepository{
		DB: db,
	}
}

func (r *overtimeRepository) CreateOvertime(overtime *models.Overtime) (*models.Overtime, error) {
	query := `
		INSERT INTO overtimes(user_id, date, overtime_hours, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		overtime.UserID,
		overtime.Date,
		overtime.OvertimeHours,
		overtime.CreatedBy,
		overtime.UpdatedBy,
		overtime.CreatedAt,
		overtime.UpdatedAt,
	).Scan(&overtime.ID)

	if err != nil {
		return nil, err
	}

	return overtime, nil
}

func (r *overtimeRepository) CheckIfOvertimeHasBeenTaken(date time.Time) (bool, error) {
	query := "SELECT id FROM overtimes where date = $1"

	var id int64
	
	err := r.DB.QueryRow(query, date).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}