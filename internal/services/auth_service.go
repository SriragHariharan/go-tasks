package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sriraghariharan/gotasks/internal/models"
	repo "github.com/sriraghariharan/gotasks/internal/repository"
	"github.com/sriraghariharan/gotasks/internal/utils"
)

func CreateNewUser(ctx context.Context, user models.User)(models.User, error){
	//check if user exists || exists ? bool : err
	userExists, err := repo.CheckUserExists(ctx, user.Email)
	if err != nil{
		return models.User{}, err
	}

	if userExists{
		return models.User{}, errors.New("user already exists")
	}

	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil{
		return models.User{}, errors.New("Something went wrong")
	}
	user.Password = hashedPassword

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

	//get hashed password from db
	passwordFromDb, err:= repo.GetUserPassword(ctx, existingUser.Email)	

	passwordValid := utils.VerifyPassword(existingUser.Password, passwordFromDb)

	if passwordValid == false {
		return models.User{}, errors.New("Invalid credentials")
	}

	return *existingUser, nil

}