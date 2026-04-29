package service

import (
	"context"
	"errors"

	"github.com/sriraghariharan/gotasks/internal/models"
	repo "github.com/sriraghariharan/gotasks/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTask(ctx context.Context, task models.Task, userID bson.ObjectID) (models.Task, error) {

	task.UserId = userID

	// TODO: save to DB
	taskId, err := repo.CreateNewTask(ctx, task)

	if err != nil {
		return models.Task{}, err
	}

	task.Id = taskId

	return task, nil
}

func GetAllTasks(ctx context.Context, userID bson.ObjectID) ([]models.Task, error) {

	allTasks, err := repo.GetAllTasksForUser(ctx, userID)

	if err != nil {
		return []models.Task{}, err
	}

	return allTasks, nil
}

//update task
func UpdateTask(ctx context.Context, taskId bson.ObjectID, task models.Task, userID bson.ObjectID) (models.Task, error) {

	taskUpdated, err := repo.UpdateTask(ctx, taskId, task, userID)

	if err != nil {
		return models.Task{}, err
	}

	if taskUpdated == false {
		return models.Task{}, errors.New("Unable to update task")
	}

	//fetch task details
	taskDetails, err := repo.GetTaskDetails(ctx, taskId)

	if err != nil {
		return models.Task{}, err
	}

	return taskDetails, nil
}

//delete task
func DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (bool, error) {

	taskDeleted, err := repo.DeleteTask(ctx, taskID, userID)

	if err != nil {
		return false, err
	}

	return taskDeleted, nil
}