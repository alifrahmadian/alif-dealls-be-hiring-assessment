package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
)

type PayrollRepository interface {
	GetDB() *sql.DB
	CreatePayroll(tx *sql.Tx, payroll *models.Payroll) (*models.Payroll, error)
	CheckIfPayrollHasBeenProcessed(userID, attendancePeriodID int64) (bool, error)
	GetEmployeePayrollByID(payrollID, userID int64) (*models.Payroll, error)
}

type payrollRepository struct {
	DB *sql.DB
}

func NewPayrollRepository(db *sql.DB) PayrollRepository {
	return &payrollRepository{
		DB: db,
	}
}

func (r *payrollRepository) GetDB() *sql.DB {
	return r.DB
}

func (r *payrollRepository) CreatePayroll(tx *sql.Tx, payroll *models.Payroll) (*models.Payroll, error) {
	query := `
		INSERT INTO payrolls(user_id, attendance_period_id, base_salary, attendance_days, attendance_amount, overtime_hours, overtime_amount, reimbursement_amount, total_take_home_pay, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`

	err := tx.QueryRow(
		query,
		payroll.UserID,
		payroll.AttendancePeriodID,
		payroll.BaseSalary,
		payroll.AttendanceDays,
		payroll.AttendanceAmount,
		payroll.OvertimeHours,
		payroll.OvertimeAmount,
		payroll.ReimbursementAmount,
		payroll.TotalTakeHomePay,
		payroll.CreatedBy,
		payroll.UpdatedBy,
		payroll.CreatedAt,
		payroll.UpdatedAt,
	).Scan(&payroll.ID)

	if err != nil {
		return nil, err
	}

	return payroll, nil
}

func (r *payrollRepository) CheckIfPayrollHasBeenProcessed(userID, attendancePeriodID int64) (bool, error) {
	query := `
		SELECT user_id, attendance_period_id FROM payrolls WHERE user_id = $1 AND attendance_period_id = $2
	`

	err := r.DB.QueryRow(query, userID, attendancePeriodID).Scan(&userID, &attendancePeriodID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *payrollRepository) GetEmployeePayrollByID(payrollID, userID int64) (*models.Payroll, error) {
	payroll := &models.Payroll{}

	query := `
		SELECT 
			id, 
			user_id, 
			attendance_period_id, 
			base_salary, 
			attendance_days, 
			attendance_amount, 
			overtime_hours, 
			overtime_amount, 
			reimbursement_amount,
			total_take_home_pay
		FROM
			payrolls
		WHERE id = $1 AND user_id = $2
	`

	err := r.DB.QueryRow(
		query,
		payrollID,
		userID,
	).Scan(
		&payroll.ID,
		&payroll.UserID,
		&payroll.AttendancePeriodID,
		&payroll.BaseSalary,
		&payroll.AttendanceDays,
		&payroll.AttendanceAmount,
		&payroll.OvertimeHours,
		&payroll.OvertimeAmount,
		&payroll.ReimbursementAmount,
		&payroll.TotalTakeHomePay,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.ErrPayrollNotFound
		}

		return nil, err
	}

	return payroll, nil
}
