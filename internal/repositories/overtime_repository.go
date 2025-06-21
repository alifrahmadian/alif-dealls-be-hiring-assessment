package repositories

import (
	"database/sql"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type OvertimeRepository interface {
	CreateOvertime(overtime *models.Overtime) (*models.Overtime, error)
	CheckIfOvertimeHasBeenTaken(date time.Time) (bool, error)
	SumOvertimeHoursPerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error)
	InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error
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

func (r *overtimeRepository) SumOvertimeHoursPerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error) {
	var sum uint64

	query := `
		SELECT COALESCE(SUM(overtime_hours), 0) FROM overtimes
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

func (r *overtimeRepository) InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error {
	query := `
		UPDATE overtimes
		SET payroll_id = $1, updated_at = NOW()
		WHERE user_id = $2 
		AND date BETWEEN $3 AND $4;
	`
	_, err := tx.Exec(query, payrollID, userID, startDate, endDate)
	if err != nil {
		return err
	}

	return nil
}
