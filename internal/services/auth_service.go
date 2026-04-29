package service

import (
	"context"
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