package domain

import "github.com/goddamnnoob/family-app-api/errs"

type User struct {
	UserName        string
	UserId          string
	UserEmail       string
	UserPhoneNumber string
	UserMother      string
	UserFather      string
	UserBrothers    []string
	UserSisters     []string
}

type UserRepository interface {
	GetAllFamilyMemebers() ([]User, *errs.AppError)
}
