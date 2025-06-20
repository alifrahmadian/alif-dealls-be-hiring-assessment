package repositories

import (
	"database/sql"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type AttendanceRepository interface {
	CreateAttendance(attendance *models.Attendance) (*models.Attendance, error)
	CheckIfUserHasRecordAttendance(userID int64, date time.Time) (bool, error)
}

type attendanceRepository struct {
	DB *sql.DB
}

func NewAttendanceRepository(db *sql.DB) AttendanceRepository {
	return &attendanceRepository{
		DB: db,
	}
}

func (r *attendanceRepository) CreateAttendance(attendance *models.Attendance) (*models.Attendance, error) {
	query := `
		INSERT INTO attendances(user_id, date, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		attendance.UserID,
		attendance.Date,
		attendance.CreatedBy,
		attendance.UpdatedBy,
		attendance.CreatedAt,
		attendance.UpdatedAt,
	).Scan(&attendance.ID)

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (r *attendanceRepository) CheckIfUserHasRecordAttendance(userID int64, date time.Time) (bool, error) {
	query := `
		SELECT id FROM attendances WHERE user_id = $1 AND date = $2
	`

	var id int64

	err := r.DB.QueryRow(query, userID, date).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
