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

	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		return models.Task{}, errors.New("user not found in context")
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

func GetAllTasks(ctx context.Context) ([]models.Task, error) {

	// get UserId from context object
	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		return []models.Task{}, errors.New("user not found in context")
	}

	allTasks, err := repo.GetAllTasksForUser(ctx, userObjectID)

	if err != nil {
		return []models.Task{}, err
	}

	return allTasks, nil
}

func DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (bool, error) {

	taskDeleted, err := repo.DeleteTask(ctx, taskID, userID)

	if err != nil {
		return false, err
	}

	return taskDeleted, nil
}