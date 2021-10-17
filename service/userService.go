package service

import (
	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
)

type UserService interface {
	GetAllFamilyMemebers() ([]domain.User, *errs.AppError)
}

type DefaUserService struct {
}
