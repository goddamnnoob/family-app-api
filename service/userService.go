package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
)

type UserService interface {
	GetAllFamilyMembers(string) ([]domain.User, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllFamilyMembers(id string) ([]domain.User, *errs.AppError) {

	return u.repo.GetAllFamilyMembers(id)
}

func NewUserService(repository domain.UserRepository) UserService {
	return DefaultUserService{repository}
}
