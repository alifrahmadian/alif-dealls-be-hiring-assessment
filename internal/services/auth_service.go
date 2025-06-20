package services

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	r "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/auth"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
)

type AuthService interface {
	Login(username, password string) (*models.User, error)
}

type authService struct {
	UserRepo r.UserRepository
}

func NewAuthService(userRepo r.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (s *authService) Login(username, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUser(username)
	if err != nil {
		return nil, err
	}

	if !auth.ComparePassword(password, user.Password) {
		return nil, e.ErrInvalidPassword
	}

	return user, nil
}
