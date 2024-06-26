package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	username string `json:"username" bson:"username"`
	avatar_url string `json:"avatar" bson: "avatar_url"`
	Date        string             `json:"date" bson:"date"`
}