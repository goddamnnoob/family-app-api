package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllFamilyMembers(string) (*domain.FamilyMembers, *errs.AppError)
	CreateUser(domain.User) (string, *errs.AppError)
	GetUserByUserId(string) (*domain.User, *errs.AppError)
	SearchUser(string, string) ([]*domain.User, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllFamilyMembers(id string) (*domain.FamilyMembers, *errs.AppError) {
	return u.repo.GetAllFamilyMembers(id)
}

func (u DefaultUserService) CreateUser(user domain.User) (string, *errs.AppError) {
	user.UserId = uuid.New().String()
	return u.repo.CreateUser(user)
}

func (u DefaultUserService) GetUserByUserId(id string) (*domain.User, *errs.AppError) {
	return u.repo.GetUserByUserId(id)
}

func (u DefaultUserService) SearchUser(key string, searchText string) ([]*domain.User, *errs.AppError) {
	if key == "location" {
		key = "user_location"
	} else if key == "name" {
		key = "user_name"
	} else if key == "phone" {
		key = "user_phone_number"
	}
	return u.repo.SearchUser(key, searchText)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
