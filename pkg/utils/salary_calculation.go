package utils

import "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/constants"

func CalculateDailyRate(baseSalary uint64) uint64 {
	return baseSalary / constants.TOTAL_WORKING_DAYS_PER_MONTH
}

func CalculateHourlyRate(baseSalary uint64) uint64 {
	return CalculateDailyRate(baseSalary) / constants.TOTAL_WORKING_HOURS_PER_DAY
}

func CalculateAttendanceAmount(baseSalary, totalAttendanceDays uint64) uint64 {
	return CalculateDailyRate(baseSalary) * totalAttendanceDays
}

func CalculateOvertimeAmount(baseSalary, totalOvertimeHours uint64) uint64 {
	return (constants.OVERTIME_PAYMENT_MULTIPLIER * CalculateHourlyRate(baseSalary)) * totalOvertimeHours
}

func CalculateTotalTakeHomePay(attendanceAmount, overtimeAmount, reimbursementAmount uint64) uint64 {
	return attendanceAmount + overtimeAmount + reimbursementAmount
}
