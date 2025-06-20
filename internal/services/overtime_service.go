package services

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/constants"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/utils"
)

type OvertimeService interface {
	CreateOvertime(overtime *models.Overtime) (*models.Overtime, error)
}

type overtimeService struct {
	OvertimeRepo r.OvertimeRepository
	AttendanceRepo r.AttendanceRepository
}

func NewOvertimeService(overtimeRepo r.OvertimeRepository, attendanceRepo r.AttendanceRepository) OvertimeService {
	return &overtimeService {
		OvertimeRepo: overtimeRepo,
		AttendanceRepo: attendanceRepo,
	}
}

func (s *overtimeService) CreateOvertime(overtime *models.Overtime) (*models.Overtime, error) {
	isOvertimeTakenToday, err := s.OvertimeRepo.CheckIfOvertimeHasBeenTaken(overtime.Date)
	if err != nil {
		return nil, err
	}

	isEmployeeHasCheckedInToday, err := s.AttendanceRepo.CheckIfUserHasRecordAttendance(overtime.UserID, overtime.Date)
	if err != nil {
		return nil, err
	}

	if utils.IsWeekday(overtime.Date){
		if !isEmployeeHasCheckedInToday {
			return nil, e.ErrEmployeeHasNotWorkedToday
		}
	}

	if overtime.OvertimeHours > constants.MAX_OVERTIME_PER_DAY {
		return nil, e.ErrOvertimeMoreThanThreeHours
	}

	if isOvertimeTakenToday {
		return nil, e.ErrOvertimeIsTakenToday
	}

	newOvertime, err := s.OvertimeRepo.CreateOvertime(overtime)
	if err != nil {
		return nil, err
	}

	return newOvertime, nil
}