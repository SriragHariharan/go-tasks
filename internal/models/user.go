package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"-" bson:"password" validate:"required,min=6"`
}