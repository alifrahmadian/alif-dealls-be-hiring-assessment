package services

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
)

type AttendancePeriodService interface {
	CreateAttendancePeriod(attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error)
}

type attendancePeriodService struct {
	AttendancePeriodRepo r.AttendancePeriodRepository
}

func NewAttendancePeriodService(attendancePeriodRepo r.AttendancePeriodRepository) AttendancePeriodService {
	return &attendancePeriodService{
		AttendancePeriodRepo: attendancePeriodRepo,
	}
}

func (s *attendancePeriodService) CreateAttendancePeriod(attendancePeriod *models.AttendancePeriod) (*models.AttendancePeriod, error) {
	newAttendancePeriod, err := s.AttendancePeriodRepo.CreateAttendancePeriod(attendancePeriod)
	if err != nil {
		return nil, err
	}

	return newAttendancePeriod, nil
}