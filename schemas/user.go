package schemas

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

type User struct {
	Name         string
	Email        string
	PasswordHash string
	SubDomain    string
	IsAdmin      bool
	UserID       int64
	Urls         []Url
}

type UserResponse struct {
	ID           primitive.ObjectID `json:"_id"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	PasswordHash string             `json:"passwordhash"`
	SubDomain    string             `json:"subdomain"`
	IsAdmin      bool               `json:"isadmin"`
	UserID       int64              `json:"userid"`
	Urls         []Url              `json:"urls"`
}

func CreateNewUser(db *mongo.Database, collectionName string, user *User) error {
	_, err := db.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(db *mongo.Database, collectionName string, filter bson.M) *UserResponse {
	var result *UserResponse
	info := db.Collection(collectionName).FindOne(ctx, filter).Decode(&result)
	if info != nil {
		return nil
	}
	return result
}
