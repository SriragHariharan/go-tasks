package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const dbConnectionString = "mongodb://localhost:27017"
const databaseName = "gotasks"

const TasksCollectionName = "tasks"
const UsersCollectionName = "users"

var TasksCollection *mongo.Collection
var UsersCollection *mongo.Collection

func Connect() {
	// create client
	client, err := mongo.Connect(options.Client().ApplyURI(dbConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	// context for operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo not reachable:", err)
	}

	fmt.Println("MongoDB connected")

	db := client.Database(databaseName)

	TasksCollection = db.Collection(TasksCollectionName)
	UsersCollection = db.Collection(UsersCollectionName)
}