package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
)

type UserService interface {
	GetAllFamilyMembers(string) (*domain.FamilyMembers, *errs.AppError)
	CreateUser(domain.User) (string, *errs.AppError)
	GetUserByUserId(string) (*domain.User, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllFamilyMembers(id string) (*domain.FamilyMembers, *errs.AppError) {

	return u.repo.GetAllFamilyMembers(id)
}

func (u DefaultUserService) CreateUser(user domain.User) (string, *errs.AppError) {
	return u.repo.CreateUser(user)
}

func (u DefaultUserService) GetUserByUserId(id string) (*domain.User, *errs.AppError) {
	return u.repo.GetUserByUserId(id)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
