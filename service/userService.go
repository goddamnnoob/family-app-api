package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
)

type UserService interface {
	GetAllFamilyMembers(string) ([]domain.User, *errs.AppError)
	CreateUser(domain.User) (string, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllFamilyMembers(id string) ([]domain.User, *errs.AppError) {

	return u.repo.GetAllFamilyMembers(id)
}

func (u DefaultUserService) CreateUser(user domain.User) (string, *errs.AppError) {
	userid, err := u.repo.CreateUser(user)
	if err != nil {
		return "0", err
	}
	return userid, nil
}

func NewUserService(repository domain.UserRepository) UserService {
	return DefaultUserService{repository}
}
