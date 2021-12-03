package domain

import (
	"context"
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
	familyMembers.Father, _ = d.GetUserByUserId(user.UserFather)

	familyMembers.Mother, _ = d.GetUserByUserId(user.UserMother)

	familyMembers.Partner, _ = d.GetUserByUserId(user.UserPartner)
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
	var sibilings []*User
	for _, us := range user.UserBrothers {
		u, _ := d.GetUserByUserId(us)
		sibilings = append(sibilings, u)
	}
	familyMembers.Sibilings = sibilings
	return &familyMembers, nil

}

func (d UserRepositoryDb) GetUserByUserId(user_id string) (*User, *errs.AppError) {
	usersCollection := d.dbClient.Database("users").Collection("users")
	var users []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := usersCollection.Find(ctx, bson.D{{"user_id", user_id}})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying db" + err.Error())
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
	_, err := usersCollection.InsertOne(ctx, &u)
	if err != nil {
		return "", errs.NewUnexpectedError("DB error")
	}
	userId := u.UserId
	return userId, nil
}

func (d UserRepositoryDb) SearchUser(key string, value string) ([]*User, *errs.AppError) {
	var users []*User
	usersCollection := d.dbClient.Database("users").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	search := bson.D{{"$search", bson.D{{"text", bson.D{{"path", key}, {"query", value}}}}}}
	cursor, err := usersCollection.Aggregate(ctx, mongo.Pipeline{search})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying db" + err.Error())
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while converting to []User from cursor" + err.Error())
	}
	if users == nil {
		return nil, errs.NewUserNotFoundError("Search user not found")
	}
	return users, nil
}

/*func (d UserRepositoryDb) FindRelationship() ([]*User, *errs.AppError) {
	var users []*User
	usersCollection := d.dbClient.Database("users").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipeline := bson.D{{"$graphlookup", bson.D{{"from", "users"}, {"startWith", "$user_mother"}, {"connectFromField", "user_"}, {"connectToField", "user_id"}, {"as", "user"}, {"maxDepth", "3"}}}}
	cursor, err := usersCollection.Aggregate(ctx, mongo.Pipeline{pipeline})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying db " + err.Error())
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while converting to []User from cursor " + err.Error())
	}
	if users == nil {
		return nil, errs.NewUserNotFoundError("Search user not found")
	}
	return users, nil
}
Graphlookup not available for free tier
*/

func (d UserRepositoryDb) FindRelationship(start string, end string) ([]*User, *errs.AppError) {
	var users []*User
	var err *errs.AppError
	var isExist bool
	var helper func(string, string, int32) bool
	helper = func(s1, s2 string, depth int32) bool {
		if depth >= 5 {
			return false
		}
		u, err := d.GetUserByUserId(s1)
		if err != nil {
			return false
		}
		users = append(users, u)
		if s1 == s2 {
			return true
		}
		familyMembers, err := d.GetAllFamilyMembers(u.UserId)
		if err != nil {
			return false
		}
		if familyMembers.Mother != nil && !Contains(users, familyMembers.Mother) {
			t := helper(familyMembers.Mother.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}
		if familyMembers.Father != nil && !Contains(users, familyMembers.Father) {
			t := helper(familyMembers.Father.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}

		if familyMembers.Partner != nil && !Contains(users, familyMembers.Partner) {
			t := helper(familyMembers.Partner.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}
		if familyMembers.Brothers != nil {
			for _, bro := range familyMembers.Brothers {
				if !Contains(users, bro) {
					t := helper(bro.UserId, s2, depth+1)
					if t {
						return true
					} else {
						users = users[:len(users)-1]
					}
				}
			}
		}
		if familyMembers.Sisters != nil {
			for _, sis := range familyMembers.Sisters {
				if !Contains(users, sis) {
					t := helper(sis.UserId, s2, depth+1)
					if t {
						return true
					} else {
						users = users[:len(users)-1]
					}
				}
			}
		}
		if familyMembers.Sibilings != nil {
			for _, sib := range familyMembers.Sibilings {
				if !Contains(users, sib) {
					t := helper(sib.UserId, s2, depth+1)
					if t {
						return true
					} else {
						users = users[:len(users)-1]
					}
				}
			}
		}
		return false
	}
	isExist = helper(start, end, 0)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errs.NewRelationshipNotFoundError("Relationship not found")
	}
	return users, nil
}

func Contains(users []*User, user *User) bool {
	for _, u := range users {
		if u.UserId == user.UserId {
			return true
		}
	}
	return false
}

//func (d UserRepositoryDb) UpdateUser()
