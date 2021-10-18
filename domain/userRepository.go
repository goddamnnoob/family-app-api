package domain

import (
	"github.com/goddamnnoob/family-app-api/errs"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDb struct {
	dbClient *mongo.Client
}

func NewUserRepository(dbClient *mongo.Client) UserRepositoryDb {
	return UserRepositoryDb{dbClient: dbClient}
}

func (d UserRepositoryDb) GetAllFamilyMembers(id string) ([]User, *errs.AppError) {
	return make([]User, 0), nil

}

func GetAllFamilyMembers(id string) ([]User, *errs.AppError) {
	users := make([]User, 0)
	return users, nil
}
