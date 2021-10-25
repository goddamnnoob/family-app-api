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

func (d UserRepositoryDb) GetAllFamilyMembers(user_id string) (*FamilyMembers, *errs.AppError) {
	user, err := d.GetUserByUserId(user_id)
	if err != nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	var familyMembers FamilyMembers
	familyMembers.Father, err = d.GetUserByUserId(user.UserFather)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying UserFather " + err.Message)
	}
	familyMembers.Mother, err = d.GetUserByUserId(user.UserMother)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying UserMother" + err.Message)
	}
	var brothers []*User
	for _, us := range user.UserBrothers {
		u, _ := d.GetUserByUserId(us)
		brothers = append(brothers, u)
	}
	familyMembers.Brothers = brothers
	var sisters []*User
	for _, us := range user.UserBrothers {
		u, _ := d.GetUserByUserId(us)
		sisters = append(sisters, u)
	}
	familyMembers.Sisters = sisters
	return &familyMembers, nil

}

func (d UserRepositoryDb) GetUserByUserId(user_id string) (*User, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	var users []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := usersCollection.Find(ctx, bson.D{{"_id", user_id}})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying " + err.Error())
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while converting to []User from cursor " + err.Error())
	}

	if users == nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	return &users[0], nil
}

func (d UserRepositoryDb) CreateUser(u User) (string, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := usersCollection.InsertOne(ctx, &u)
	if err != nil {
		return "", errs.NewUnexpectedError("DB error")
	}
	userId := fmt.Sprint(result.InsertedID)[10:34]
	return userId, nil
}
