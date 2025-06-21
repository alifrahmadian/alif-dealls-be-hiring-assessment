package repositories

import (
	"database/sql"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
)

type AttendanceRepository interface {
	CreateAttendance(attendance *models.Attendance) (*models.Attendance, error)
	CheckIfUserHasRecordAttendance(userID int64, date time.Time) (bool, error)
	CountEmployeeAttendancePerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error)
	InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error
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

func (r *attendanceRepository) CountEmployeeAttendancePerPeriod(userID int64, startDate time.Time, endDate time.Time) (uint64, error) {
	var count uint64

	query := `
		SELECT COUNT(*) FROM attendances
		WHERE user_id = $1
		AND payroll_id IS NULL
		AND date BETWEEN $2 AND $3;
	`

	err := r.DB.QueryRow(
		query,
		userID,
		startDate,
		endDate,
	).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *attendanceRepository) InsertPayrollID(tx *sql.Tx, userID, payrollID int64, startDate, endDate time.Time) error {
	query := `
		UPDATE attendances
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
