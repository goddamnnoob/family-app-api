package domain

import (
	"github.com/goddamnnoob/family-app-api/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId          string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserName        string             `bson:"user_name" json:"user_name"`
	UserGender      string             `bson:"user_gender" json:"user_gender"`
	UserEmail       string             `bson:"user_email" json:"user_email"`
	UserPhoneNumber string             `bson:"user_phone_number" json:"user_phone_number"`
	UserPassword    string             `bson:"user_password" json:"user_password"`
	UserMother      string             `bson:"user_mother" json:"user_mother"`
	UserFather      string             `bson:"user_father" json:"user_father"`
	UserPartner     string             `bson:"user_partner" json:"user_partner"`
	UserBrothers    []string           `bson:"user_brothers" json:"user_brothers"`
	UserSisters     []string           `bson:"user_sisters" json:"user_sisters"`
	UserSibilings   []string           `bson:"user_sibilings" json:"user_sibilings"`
	UserLocation    string             `bson:"user_location" json:"user_location"`
}

type FamilyMembers struct {
	Father    *User   `json:"user_father,omitempty"`
	Mother    *User   `json:"user_mother,omitempty"`
	Partner   *User   `json:"user_partner,omitempty"`
	Brothers  []*User `json:"user_brothers,omitempty"`
	Sisters   []*User `json:"user_sisters,omitempty"`
	Sibilings []*User `json:"user_sibilings,omitempty"`
}

type UserRepository interface {
	GetAllFamilyMembers(string) (*FamilyMembers, *errs.AppError)
	CreateUser(User) (string, *errs.AppError)
	GetUserByUserId(string) (*User, *errs.AppError)
	SearchUser(string, string) ([]*User, *errs.AppError)
	FindRelationship(string, string) ([]*User, *errs.AppError)
}
