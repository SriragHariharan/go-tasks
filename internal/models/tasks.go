package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	IsCompleted bool `json:"isCompleted" bson:"isCompleted"`
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
}