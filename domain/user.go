package domain

import "github.com/goddamnnoob/family-app-api/errs"

type User struct {
	UserName        string   `bson:"user_name" json:"user_name"`
	UserId          string   `bson:"user_id" json:"user_id"`
	UserEmail       string   `bson:"user_email" json:"user_email"`
	UserPhoneNumber string   `bson:"user_phone_number" json:"user_phone_number"`
	UserMother      string   `bson:"user_mother" json:"user_mother"`
	UserFather      string   `bson:"user_father" json:"user_father"`
	UserBrothers    []string `bson:"user_brothers" json:"user_brothers"`
	UserSisters     []string `bson:"user_sisters" json:"user_sisters"`
}

type UserRepository interface {
	GetAllFamilyMembers(string) ([]User, *errs.AppError)
	CreateUser(User) (string, *errs.AppError)
}
