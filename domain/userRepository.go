package domain

import (
	"context"
	"fmt"
	"time"

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

func (d UserRepositoryDb) CreateUser(u User) (string, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := usersCollection.InsertOne(ctx, &u)
	if err != nil {
		return "", errs.NewUnexpectedError("DB error")
	}
	return fmt.Sprint(result.InsertedID), nil
}
