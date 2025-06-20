package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type AttendancePeriodRepository interface {
	CreateAttendancePeriod (attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error)
}

type attendanceRepository struct {
	DB *sql.DB
}

func NewAttendancePeriodRepository(db *sql.DB) AttendancePeriodRepository {
	return &attendanceRepository{
		DB: db,
	}
}

func (r *attendanceRepository) CreateAttendancePeriod (attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error) {
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

	if err != nil{
		return nil, err
	}

	return attendancePeriod, nil
}
