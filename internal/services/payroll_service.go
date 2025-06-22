package services

import (
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/utils"
)

type PayrollService interface {
	CreatePayroll(payroll *models.Payroll) (*models.Payroll, error)
	GenerateEmployeePayslip(userID, attendancePeriodID, payrollID int64) (*dtos.GeneratePayslipPayrollResponse, error)
}

type payrollService struct {
	PayrollRepo          r.PayrollRepository
	UserRepo             r.UserRepository
	AttendancePeriodRepo r.AttendancePeriodRepository
	AttendanceRepo       r.AttendanceRepository
	OvertimeRepo         r.OvertimeRepository
	ReimbursementRepo    r.ReimbursementRepository
}

func NewPayrollService(
	payrollRepo r.PayrollRepository,
	userRepo r.UserRepository,
	attendancePeriodRepo r.AttendancePeriodRepository,
	attendanceRepo r.AttendanceRepository,
	overtimeRepo r.OvertimeRepository,
	reimbursementRepo r.ReimbursementRepository,
) PayrollService {
	return &payrollService{
		PayrollRepo:          payrollRepo,
		UserRepo:             userRepo,
		AttendancePeriodRepo: attendancePeriodRepo,
		AttendanceRepo:       attendanceRepo,
		OvertimeRepo:         overtimeRepo,
		ReimbursementRepo:    reimbursementRepo,
	}
}

func (s *payrollService) CreatePayroll(payroll *models.Payroll) (*models.Payroll, error) {
	user, err := s.UserRepo.GetUserByID(payroll.UserID)
	if err != nil {
		return nil, err
	}

	attendancePeriod, err := s.AttendancePeriodRepo.GetAttendancePeriodByID(payroll.AttendancePeriodID)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(24 * time.Hour)
	startDate := attendancePeriod.StartDate.Truncate(24 * time.Hour)

	if now.Before(startDate) {
		return nil, e.ErrPayrollAttendancePeriodNotStartedYet
	}

	isPayrollHasBeenProcessed, err := s.PayrollRepo.CheckIfPayrollHasBeenProcessed(user.ID, attendancePeriod.ID)
	if err != nil {
		return nil, err
	}

	if isPayrollHasBeenProcessed {
		return nil, e.ErrPayrollHasBeenProcessed
	}

	totalAttendanceDays, err := s.AttendanceRepo.CountEmployeeAttendancePerPeriod(user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if err != nil {
		return nil, err
	}

	totalOvertimeHours, err := s.OvertimeRepo.SumOvertimeHoursPerPeriod(user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if err != nil {
		return nil, err
	}

	totalReimbursementAmount, err := s.ReimbursementRepo.SumReimbursementAmountPerPeriod(user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if err != nil {
		return nil, err
	}

	attendanceAmount := utils.CalculateAttendanceAmount(*user.BaseSalary, totalAttendanceDays)
	overtimeAmount := utils.CalculateOvertimeAmount(*user.BaseSalary, totalOvertimeHours)

	totalTakeHomePay := utils.CalculateTotalTakeHomePay(attendanceAmount, overtimeAmount, totalReimbursementAmount)

	payroll = &models.Payroll{
		UserID:              payroll.UserID,
		AttendancePeriodID:  payroll.AttendancePeriodID,
		BaseSalary:          *user.BaseSalary,
		AttendanceDays:      totalAttendanceDays,
		AttendanceAmount:    attendanceAmount,
		OvertimeHours:       totalOvertimeHours,
		OvertimeAmount:      overtimeAmount,
		ReimbursementAmount: totalReimbursementAmount,
		TotalTakeHomePay:    totalTakeHomePay,
		CreatedBy:           payroll.CreatedBy,
		UpdatedBy:           payroll.UpdatedBy,
		CreatedAt:           payroll.CreatedAt,
		UpdatedAt:           payroll.UpdatedAt,
	}

	tx, err := s.PayrollRepo.GetDB().Begin()
	if err != nil {
		return nil, err
	}

	txErr := error(nil)
	defer func() {
		if txErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	newPayroll, txErr := s.PayrollRepo.CreatePayroll(tx, payroll)
	if txErr != nil {
		return nil, txErr
	}

	txErr = s.AttendanceRepo.InsertPayrollID(tx, newPayroll.ID, user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if txErr != nil {
		tx.Rollback()
		return nil, txErr
	}

	txErr = s.OvertimeRepo.InsertPayrollID(tx, newPayroll.ID, user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if txErr != nil {
		tx.Rollback()
		return nil, txErr
	}

	txErr = s.ReimbursementRepo.InsertPayrollID(tx, newPayroll.ID, user.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if txErr != nil {
		tx.Rollback()
		return nil, txErr
	}

	return newPayroll, err
}

func (s *payrollService) GenerateEmployeePayslip(userID, attendancePeriodID, payrollID int64) (*dtos.GeneratePayslipPayrollResponse, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	attendanceDetails, err := s.AttendanceRepo.GetEmployeeAttendancesByPayrollID(payrollID, userID)
	if err != nil {
		return nil, err
	}

	overtimeDetails, err := s.OvertimeRepo.GetEmployeeOvertimesByPayrollID(payrollID, userID)
	if err != nil {
		return nil, err
	}

	reimbursements, err := s.ReimbursementRepo.GetEmployeeReimbursementsByPayrollID(payrollID, userID)
	if err != nil {
		return nil, err
	}

	payslip, err := s.PayrollRepo.GetEmployeePayrollByID(payrollID, userID)
	if err != nil {
		return nil, err
	}

	attendanceResponses := make([]*dtos.GeneratePayslipAttendanceResponse, len(attendanceDetails))
	for i, attendanceDetail := range attendanceDetails {
		attendanceResponses[i] = &dtos.GeneratePayslipAttendanceResponse{
			ID:             attendanceDetail.ID,
			UserID:         attendanceDetail.UserID,
			PayrollID:      attendanceDetail.PayrollID,
			Date:           attendanceDetail.Date.Format("2006-01-02"),
			AttendanceRate: utils.CalculateDailyRate(payslip.BaseSalary),
		}
	}

	overtimeResponses := make([]*dtos.GeneratePayslipOvertimeResponse, len(overtimeDetails))
	for i, overtimeDetail := range overtimeDetails {
		overtimeResponses[i] = &dtos.GeneratePayslipOvertimeResponse{
			ID:            overtimeDetail.ID,
			UserID:        overtimeDetail.UserID,
			PayrollID:     *overtimeDetail.PayrollID,
			Date:          overtimeDetail.Date.Format("2006-01-02"),
			OvertimeHours: overtimeDetail.OvertimeHours,
			OvertimeRate:  utils.CalculateOvertimeAmount(payslip.BaseSalary, overtimeDetail.OvertimeHours),
		}
	}

	reimbursementResponses := make([]*dtos.GeneratePayslipReimbursementResponse, len(reimbursements))
	for i, reimbursementDetail := range reimbursements {
		reimbursementResponses[i] = &dtos.GeneratePayslipReimbursementResponse{
			ID:                  reimbursementDetail.ID,
			UserID:              reimbursementDetail.UserID,
			PayrollID:           *reimbursementDetail.PayrollID,
			ReimbursementAmount: reimbursementDetail.ReimbursementAmount,
			Date:                reimbursementDetail.Date.Format("2006-01-02"),
			Description:         reimbursementDetail.Description,
		}
	}

	response := &dtos.GeneratePayslipPayrollResponse{
		ID:                  payslip.ID,
		Username:            user.Username,
		AttendancePeriodID:  payslip.AttendancePeriodID,
		BaseSalary:          payslip.BaseSalary,
		AttendanceDays:      payslip.AttendanceDays,
		AttendanceAmount:    payslip.AttendanceAmount,
		AttendanceDetails:   attendanceResponses,
		OvertimeHours:       payslip.OvertimeHours,
		OvertimeAmount:      payslip.OvertimeAmount,
		OvertimeDetails:     overtimeResponses,
		ReimbursementAmount: payslip.ReimbursementAmount,
		Reimbursements:      reimbursementResponses,
		TotalTakeHomePay:    payslip.TotalTakeHomePay,
	}

	return response, nil
}
