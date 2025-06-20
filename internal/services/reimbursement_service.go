package services

import (
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
)

type ReimbursementService interface {
	CreateReimbursement(reimbursement *models.Reimbursement) (*models.Reimbursement, error)
}

type reimbursementService struct {
	ReimbursementRepo r.ReimbursementRepository
}

func NewReimbursementService(reimbursementRepo r.ReimbursementRepository) ReimbursementService {
	return &reimbursementService{
		ReimbursementRepo: reimbursementRepo,
	}
}

func (s *reimbursementService) CreateReimbursement(reimbursement *models.Reimbursement) (*models.Reimbursement, error) {
	if reimbursement.Date.After(time.Now()) {
		return nil, e.ErrReimbursementDateIsAfterToday
	}

	newReimbursement, err := s.ReimbursementRepo.CreateReimbursement(reimbursement)
	if err != nil {
		return nil, err
	}

	return newReimbursement, nil
}
