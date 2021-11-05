package service

import (
	"github.com/faheemuddinkhan/clinic-auth-service/api/dto"
	domainport "github.com/faheemuddinkhan/clinic-auth-service/pkg/domain/port"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/errors"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
)


//DrivenPort: service interface for {User}
type IUserService interface {
	CreateUser(request dto.NewUserRequest) (*dto.NewUserResponse, *errors.AppError)
	CheckAvailable(string, interface{}) (bool, *errors.AppError)
}


//Adapter for IUserService
type DefaultUserService struct {
	repo domainport.IUserRepo
}

func (s DefaultUserService) CreateUser(request dto.NewUserRequest) (*dto.NewUserResponse, *errors.AppError) {
	logger.Info("UserCreate")
	return nil, nil
}

func (s DefaultUserService) CheckAvailable(field string, value interface{}) (bool, *errors.AppError) {
	exists, err := s.repo.CheckExists("field", value)

	if err != nil {
		return false, err
	}

	return !exists, nil
}

func NewUserService(repo domainport.IUserRepo) DefaultUserService {
	return DefaultUserService{repo: repo}
}
