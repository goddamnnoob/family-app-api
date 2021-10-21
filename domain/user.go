package domain

import (
	"github.com/goddamnnoob/family-app-api/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName        string             `bson:"user_name" json:"user_name"`
	UserId          string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserEmail       string             `bson:"user_email" json:"user_email"`
	UserPhoneNumber string             `bson:"user_phone_number" json:"user_phone_number"`
	UserMother      string             `bson:"user_mother" json:"user_mother"`
	UserFather      string             `bson:"user_father" json:"user_father"`
	UserBrothers    []string           `bson:"user_brothers" json:"user_brothers"`
	UserSisters     []string           `bson:"user_sisters" json:"user_sisters"`
}

type UserRepository interface {
	GetAllFamilyMembers(string) ([]User, *errs.AppError)
	CreateUser(User) (string, *errs.AppError)
	GetUserByUserId(string) ([]User, *errs.AppError)
}
