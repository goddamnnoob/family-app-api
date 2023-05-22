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
	var children []*User
	for _, us := range user.UserBrothers {
		u, _ := d.GetUserByUserId(us)
		children = append(children, u)
	}
	familyMembers.Children = children
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
	if len(u.UserFather) > 0 {
		father, er := d.GetUserByUserId(u.UserFather)
		if er == nil {
			var children []string
			if father.UserChildren != nil && len(father.UserChildren) > 0 {
				if !ContainsString(father.UserChildren, u.UserId) {
					children = father.UserChildren
				}
				children = append(children, u.UserId)
				_, err = usersCollection.ReplaceOne(ctx,
					bson.M{"user_id": father.UserId},
					bson.M{"user_children": children},
				)
				if err != nil {
					return "", errs.NewUnexpectedError("DB error while updating father sibiling relationship ")
				}
			}
		}
	}
	if len(u.UserMother) > 0 {
		mother, er := d.GetUserByUserId(u.UserMother)
		if er == nil {
			var children []string
			if mother.UserChildren != nil && len(mother.UserChildren) > 0 {
				if !ContainsString(mother.UserChildren, u.UserId) {
					children = mother.UserChildren
				}
				children = append(children, u.UserId)
				_, err = usersCollection.ReplaceOne(ctx, bson.M{"user_id": mother.UserId}, bson.M{"user_children": children})
				if err != nil {
					return "", errs.NewUnexpectedError("DB error while updating mother sibiling relationship " + err.Error())
				}

			}
		}
	}
	if len(u.UserBrothers) > 0 {
		for _, brotherId := range u.UserBrothers {
			brother, er := d.GetUserByUserId(brotherId)
			if er == nil {
				if u.UserGender == "m" {
					var brothers []string
					if brother.UserBrothers != nil && len(brother.UserBrothers) > 0 {
						if !ContainsString(brother.UserBrothers, u.UserId) {
							brothers = brother.UserBrothers
						}
						brothers = append(brothers, u.UserId)
						_, err = usersCollection.ReplaceOne(ctx, bson.M{"user_id": brother.UserId}, bson.M{"user_brothers": brothers})
						if err != nil {
							return "", errs.NewUnexpectedError("DB error while updating brother brother relationship" + err.Error())
						}
					}
				} else if u.UserGender == "f" {
					var sisters []string
					if brother.UserSisters != nil && len(brother.UserSisters) > 0 {
						if !ContainsString(brother.UserSisters, u.UserId) {
							sisters = brother.UserSisters
						}
						sisters = append(sisters, u.UserId)
						_, err = usersCollection.ReplaceOne(ctx, bson.M{"user_id": brother.UserId}, bson.M{"user_sisters": sisters})
						if err != nil {
							return "", errs.NewUnexpectedError("DB error while updating brother sister relationship" + err.Error())
						}
					}
				}
			}
		}
	}
	if len(u.UserSisters) > 0 {
		for _, sisterId := range u.UserSisters {
			sister, er := d.GetUserByUserId(sisterId)
			if er == nil {
				if u.UserGender == "m" {
					var brothers []string
					if sister.UserBrothers != nil && len(sister.UserBrothers) > 0 {
						if !ContainsString(sister.UserBrothers, u.UserId) {
							brothers = sister.UserBrothers
						}
						brothers = append(brothers, u.UserId)
						_, err = usersCollection.ReplaceOne(ctx, bson.M{"user_id": sister.UserId}, bson.M{"user_brothers": brothers})
						if err != nil {
							return "", errs.NewUnexpectedError("DB error while updating sister brother relationship" + err.Error())
						}
					}
				} else if u.UserGender == "f" {
					var sisters []string
					if sister.UserSisters != nil && len(sister.UserSisters) > 0 {
						if !ContainsString(sister.UserSisters, u.UserId) {
							sisters = sister.UserSisters
						}
						sisters = append(sisters, u.UserId)
						_, err = usersCollection.ReplaceOne(ctx, bson.M{"user_id": sister.UserId}, bson.M{"user_sisters": sisters})
						if err != nil {
							return "", errs.NewUnexpectedError("DB error while updating sister sister relationship " + err.Error())
						}
					}
				}
			}
		}
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
		if familyMembers.Mother != nil && !ContainsUser(users, familyMembers.Mother) {
			t := helper(familyMembers.Mother.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}
		if familyMembers.Father != nil && !ContainsUser(users, familyMembers.Father) {
			t := helper(familyMembers.Father.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}

		if familyMembers.Partner != nil && !ContainsUser(users, familyMembers.Partner) {
			t := helper(familyMembers.Partner.UserId, s2, depth+1)
			if t {
				return true
			} else {
				users = users[:len(users)-1]
			}
		}
		if familyMembers.Brothers != nil {
			for _, bro := range familyMembers.Brothers {
				if !ContainsUser(users, bro) {
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
				if !ContainsUser(users, sis) {
					t := helper(sis.UserId, s2, depth+1)
					if t {
						return true
					} else {
						users = users[:len(users)-1]
					}
				}
			}
		}
		if familyMembers.Children != nil {
			for _, sib := range familyMembers.Children {
				if !ContainsUser(users, sib) {
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

func ContainsUser(users []*User, user *User) bool {
	for _, u := range users {
		if u.UserId == user.UserId {
			return true
		}
	}
	return false
}

func ContainsString(strings []string, str string) bool {
	for _, s := range strings {
		if s == str {
			return true
		}
	}
	return false
}

//func (d UserRepositoryDb) UpdateUser()
