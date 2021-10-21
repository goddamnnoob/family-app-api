package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/goddamnnoob/family-app-api/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDb struct {
	dbClient *mongo.Client
}

func NewUserRepository(dbClient *mongo.Client) UserRepositoryDb {
	return UserRepositoryDb{dbClient: dbClient}
}

func (d UserRepositoryDb) GetAllFamilyMembers(user_id string) ([]User, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	var users []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := usersCollection.Find(ctx, bson.D{{"user_id", user_id}})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying " + err.Error())
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while converting to []User from cursor " + err.Error())
	}
	return users, nil

}

func (d UserRepositoryDb) GetUserByUserId(user_id string) ([]User, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	var users []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := usersCollection.Find(ctx, bson.D{{"user_id", user_id}})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying " + err.Error())
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while converting to []User from cursor " + err.Error())
	}
	return users, nil
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
