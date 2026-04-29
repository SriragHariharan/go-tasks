package service

import (
	"context"
	"errors"

	"github.com/sriraghariharan/gotasks/internal/middleware"
	"github.com/sriraghariharan/gotasks/internal/models"
	repo "github.com/sriraghariharan/gotasks/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTask(ctx context.Context, task models.Task) (models.Task, error) {

	userIdStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return models.Task{}, errors.New("user not found in context")
	}

	userObjectID, err := bson.ObjectIDFromHex(userIdStr)
	if err != nil {
		return models.Task{}, err
	}

	task.UserId = userObjectID

	// TODO: save to DB
	taskId, err := repo.CreateNewTask(ctx, task)

	if err != nil {
		return models.Task{}, err
	}

	task.Id = taskId

	return task, nil
}