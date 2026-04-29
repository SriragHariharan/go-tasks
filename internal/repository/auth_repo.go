package repo

import (
	"context"
	"errors"
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

	if err != nil{
		return models.User{}, fmt.Errorf("Unable to create user")
	}

	//attach user id to user obj
	newUser.UserId = userSavedResponse.InsertedID.(bson.ObjectID)
	return newUser, nil
}

func GetUserPassword(ctx context.Context, userEmail string) (string, error){
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"email": userEmail}

	var existingUser models.User
	err := database.UsersCollection.FindOne(ctx, filter).Decode(&existingUser)

	if err != nil {
		return "", errors.New("Unable to login")
	}
	
	return existingUser.Password, nil
}

func GetUserDetails(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var userDetails models.User
	err := database.UsersCollection.FindOne(ctx, filter).Decode(&userDetails)

	fmt.Printf("user details:%+v", userDetails)

	if err != nil {
		return models.User{}, errors.New("Unable to get details")
	}

	//remove password from user obj
	userDetails.Password = ""
	
	return userDetails, nil
}