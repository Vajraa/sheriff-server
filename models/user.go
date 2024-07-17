package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	AvatarURL string             `json:"avatar_url" bson:"avatar_url"`
	Date      string             `json:"date" bson:"date"`
}
