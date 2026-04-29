package repo

import (
	"context"
	"fmt"
	"time"

	database "github.com/sriraghariharan/gotasks/internal/db"
	"github.com/sriraghariharan/gotasks/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CheckUserExists(ctx context.Context, email string) (bool, error){
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//check if user exists
	count, err := database.UsersCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil{
		return false, fmt.Errorf("Something went wrong")
	}
	return count > 0, nil
}

func CreateNewUser(ctx context.Context, newUser models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userSavedResponse, err := database.UsersCollection.InsertOne(ctx, newUser)
	
	fmt.Printf("%+v", userSavedResponse)
	fmt.Printf("%+v", err)

	if err != nil{
		return models.User{}, fmt.Errorf("Unable to create user")
	}

	//attach user id to user obj
	newUser.UserId = userSavedResponse.InsertedID.(bson.ObjectID)
	return newUser, nil
}