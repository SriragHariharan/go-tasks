package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	UserId   bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password" bson:"password" validate:"required,min=6"`
}