package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sriraghariharan/gotasks/internal/models"
	repo "github.com/sriraghariharan/gotasks/internal/repository"
)

func CreateNewUser(ctx context.Context, user models.User)(models.User, error){
	//check if user exists || exists ? bool : err
	userExists, err := repo.CheckUserExists(ctx, user.Email)
	if err != nil{
		return models.User{}, err
	}

	if userExists{
		return models.User{}, fmt.Errorf("user already exists")
	}

	//save to db return user obj
	newUser, err := repo.CreateNewUser(ctx, user)

	return newUser, nil

}

func LoginUser(ctx context.Context, existingUser *models.User) (models.User, error) {
	fmt.Println(existingUser)

	userExists, err := repo.CheckUserExists(ctx, existingUser.Email)

	if err != nil {
		return models.User{}, err
	}

	if userExists == false {
		return models.User{}, errors.New("Invalid credentials")
	}

	//get user's password
	passwordFromDb, err:= repo.GetUserPassword(ctx, existingUser.Email)	

	if existingUser.Password != passwordFromDb {
		return models.User{}, errors.New("Invalid credentials")
	}

	return *existingUser, nil

}