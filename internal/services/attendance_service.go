package services

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/utils"
)

type AttendanceService interface {
	CreateAttendance(attendance *models.Attendance) (*models.Attendance, error)
}

type attendanceService struct {
	AttendanceRepo r.AttendanceRepository
}

func NewAttendanceService(attendanceRepo r.AttendanceRepository) AttendanceService {
	return &attendanceService{
		AttendanceRepo: attendanceRepo,
	}
}

func (s *attendanceService) CreateAttendance(attendance *models.Attendance) (*models.Attendance, error) {
	isAttendanceRecorded, err := s.AttendanceRepo.CheckIfUserHasRecordAttendance(attendance.UserID, attendance.Date)
	if err != nil {
		return nil, err
	}

	if isAttendanceRecorded {
		return nil, e.ErrEmployeeHasRecordAttendance
	}

	
	if utils.IsWeekend(attendance.Date) {
		return nil, e.ErrIsWeekend
	}

	newAttendance, err := s.AttendanceRepo.CreateAttendance(attendance)
	if err != nil {
		return nil, err
	}

	return newAttendance, nil
}