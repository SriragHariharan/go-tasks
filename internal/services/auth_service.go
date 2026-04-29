package service

import (
	"context"
	"errors"

	"github.com/sriraghariharan/gotasks/internal/models"
	repo "github.com/sriraghariharan/gotasks/internal/repository"
	"github.com/sriraghariharan/gotasks/internal/utils"
)

func CreateNewUser(ctx context.Context, user models.User)(string, error){
	//check if user exists || exists ? bool : err
	userExists, err := repo.CheckUserExists(ctx, user.Email)
	if err != nil{
		return "", err
	}

	if userExists{
		return "", errors.New("user already exists")
	}

	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil{
		return "", errors.New("Something went wrong")
	}
	user.Password = hashedPassword

	//save to db return user obj
	newUser, err := repo.CreateNewUser(ctx, user)

	//generate jwt token
	token, err := utils.GenerateJwt(newUser.UserId, newUser.Email)
	if err != nil{
		return "", errors.New("Something went wrong")
	}

	return token, nil

}

func LoginUser(ctx context.Context, existingUser *models.User) (string, error) {

	userExists, err := repo.CheckUserExists(ctx, existingUser.Email)

	if err != nil {
		return "", err
	}

	if userExists == false {
		return "", errors.New("Invalid credentials")
	}

	//get hashed password from db
	passwordFromDb, err:= repo.GetUserPassword(ctx, existingUser.Email)	

	passwordValid := utils.VerifyPassword(existingUser.Password, passwordFromDb)

	if passwordValid == false {
		return "", errors.New("Invalid credentials")
	}

	//get userId from db
	userDetails, err := repo.GetUserDetails(ctx, existingUser.Email)

	if err != nil {
		return "", err
	}

	//generate jwt token
	token, err := utils.GenerateJwt(userDetails.UserId, userDetails.Email)
	if err != nil{
		return "", errors.New("Something went wrong")
	}

	return token, nil

}