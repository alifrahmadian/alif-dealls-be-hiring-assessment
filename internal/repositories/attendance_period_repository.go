package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
)

type AttendancePeriodRepository interface {
	CreateAttendancePeriod(attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error)
	GetAttendancePeriodByID(id int64) (*models.AttendancePeriod, error)
}

type attendancePeriodRepository struct {
	DB *sql.DB
}

func NewAttendancePeriodRepository(db *sql.DB) AttendancePeriodRepository {
	return &attendancePeriodRepository{
		DB: db,
	}
}

func (r *attendancePeriodRepository) CreateAttendancePeriod(attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error) {
	query := "INSERT INTO attendance_periods(start_date, end_date, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := r.DB.QueryRow(
		query,
		attendancePeriod.StartDate,
		attendancePeriod.EndDate,
		attendancePeriod.CreatedBy,
		attendancePeriod.UpdatedBy,
		attendancePeriod.CreatedAt,
		attendancePeriod.UpdatedAt,
	).Scan(&attendancePeriod.ID)

	if err != nil {
		return nil, err
	}

	return attendancePeriod, nil
}

func (r *attendancePeriodRepository) GetAttendancePeriodByID(id int64) (*models.AttendancePeriod, error) {
	query := `
		SELECT id, start_date, end_date FROM attendance_periods WHERE id = $1
	`

	attendancePeriod := &models.AttendancePeriod{}

	err := r.DB.
		QueryRow(query, id).
		Scan(
			&attendancePeriod.ID,
			&attendancePeriod.StartDate,
			&attendancePeriod.EndDate,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.ErrAttendancePeriodNotFound
		}

		return nil, err
	}

	return attendancePeriod, nil
}
