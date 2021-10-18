package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
)

type UserService interface {
	GetAllFamilyMembers(string) ([]domain.User, *errs.AppError)
}

type DefaUserService struct {
	repo domain.UserRepository
}

func (u DefaUserService) GetAllFamilyMembers(id string) ([]domain.User, *errs.AppError) {

	return u.repo.GetAllFamilyMembers(id)
}
