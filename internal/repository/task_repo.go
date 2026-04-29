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

// Create new task
func CreateNewTask(ctx context.Context, task models.Task) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	taskSavedResponse, err := database.TasksCollection.InsertOne(ctx, task)

	if err != nil{
		fmt.Println(err)
		return bson.ObjectID{}, errors.New("Unable to create task")
	}

	return taskSavedResponse.InsertedID.(bson.ObjectID), nil
}

// Get task details
func GetTaskDetails(ctx context.Context, taskID string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	taskObjectID, err := bson.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Task{}, errors.New("Invalid task ID")
	}

	filter := bson.M{"_id": taskObjectID}

	var taskDetails models.Task
	err = database.TasksCollection.FindOne(ctx, filter).Decode(&taskDetails)

	if err != nil {
		return models.Task{}, errors.New("Unable to get task details")
	}

	return taskDetails, nil
}

//get all tasks for a user
func GetAllTasksForUser(ctx context.Context, userID string) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	fmt.Println("userid", userID)

	userObjectID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return []models.Task{}, errors.New("Invalid user ID")
	}

	filter := bson.M{"userId": userObjectID}

	cursor, err := database.TasksCollection.Find(ctx, filter)

	if err != nil {
		return []models.Task{}, errors.New("Unable to get tasks")
	}

	var tasks []models.Task
	err = cursor.All(ctx, &tasks)

	if err != nil {
		return []models.Task{}, errors.New("Unable to get tasks")
	}

	fmt.Println("tasks", tasks)

	return tasks, nil
}