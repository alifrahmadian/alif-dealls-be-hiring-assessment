package dtos

type CreatePayrollRequest struct {
	UserID             int64 `json:"user_id" binding:"required"`
	AttendancePeriodID int64 `json:"attendance_period_id" binding:"required"`
}

type CreatePayrollResponse struct {
	ID                  int64  `json:"id"`
	UserID              int64  `json:"user_id"`
	AttendancePeriodID  int64  `json:"attendance_period_id"`
	BaseSalary          uint64 `json:"base_salary"`
	AttendanceDays      uint64 `json:"attendance_days"`
	AttendanceAmount    uint64 `json:"attendance_amount"`
	OvertimeHours       uint64 `json:"overtime_hours"`
	OvertimeAmount      uint64 `json:"overtime_amount"`
	ReimbursementAmount uint64 `json:"reimbursement_amount"`
	TotalTakeHomePay    uint64 `json:"total_take_home_pay"`
	CreatedBy           int64  `json:"created_by"`
	UpdatedBy           int64  `json:"updated_by"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type GeneratePayslipPayrollResponse struct {
	ID                   int64                                   `json:"id"`
	Username             string                                  `json:"user_name"`
	AttendancePeriodID   int64                                   `json:"attendance_period_id"`
	PeriodStartDate      string                                  `json:"period_start_date"`
	PeriodEndDate        string                                  `json:"period_end_date"`
	BaseSalary           uint64                                  `json:"base_salary"`
	AttendanceDays       uint64                                  `json:"attendance_days"`
	AttendanceAmount     uint64                                  `json:"attendance_amount"`
	AttendanceDailyRate  uint64                                  `json:"attendance_daily_rate"`
	AttendanceHourlyRate uint64                                  `json:"attendance_hourly_rate"`
	AttendanceDetails    []*GeneratePayslipAttendanceResponse    `json:"attendance_details"`
	OvertimeHours        uint64                                  `json:"overtime_hours"`
	OvertimeAmount       uint64                                  `json:"overtime_amount"`
	OvertimeDetails      []*GeneratePayslipOvertimeResponse      `json:"overtime_details"`
	ReimbursementAmount  uint64                                  `json:"reimbursement_amount"`
	Reimbursements       []*GeneratePayslipReimbursementResponse `json:"reimbursements"`
	TotalTakeHomePay     uint64                                  `json:"total_take_home_pay"`
}

type EmployeePayslipSummary struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	TakeHomePay uint64 `json:"take_home_pay"`
}

type PayslipSummary struct {
	AttendancePeriodID int64                     `json:"attendance_period_id"`
	EmployeeSummaries  []*EmployeePayslipSummary `json:"employee_summaries"`
	PeriodStartDate    string                    `json:"period_start_date"`
	PeriodEndDate      string                    `json:"period_end_date"`
	TotalTakeHomePay   uint64                    `json:"total_take_home_pay"`
}
