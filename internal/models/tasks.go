package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Task struct {
	Id bson.ObjectID `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	IsCompleted bool `json:"isCompleted" bson:"isCompleted"`
	UserId bson.ObjectID `json:"userId" bson:"userId"`
}