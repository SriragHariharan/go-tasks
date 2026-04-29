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
func GetTaskDetails(ctx context.Context, taskID bson.ObjectID) (models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": taskID}

	var taskDetails models.Task
	err := database.TasksCollection.FindOne(ctx, filter).Decode(&taskDetails)

	if err != nil {
		return models.Task{}, errors.New("Unable to get task details")
	}

	return taskDetails, nil
}

//get all tasks for a user
func GetAllTasksForUser(ctx context.Context, userID bson.ObjectID) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID}

	cursor, err := database.TasksCollection.Find(ctx, filter)

	if err != nil {
		return []models.Task{}, errors.New("Unable to get tasks")
	}

	var tasks []models.Task
	err = cursor.All(ctx, &tasks)

	if err != nil {
		return []models.Task{}, errors.New("Unable to get tasks")
	}

	return tasks, nil
}

// update task
func UpdateTask(ctx context.Context, taskId bson.ObjectID, task models.Task, userID bson.ObjectID) (bool, error) {

	fmt.Println(taskId, task, userID)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": taskId, "userId": userID}

	update := bson.M{"$set": bson.M{"title": task.Title, "isCompleted": task.IsCompleted}}

	updateResp, err := database.TasksCollection.UpdateOne(ctx, filter, update)

	fmt.Printf("%+v\n", updateResp)
	fmt.Println(err)
	
	if err != nil {
		return false, errors.New("Unable to update task")
	}

	if updateResp.ModifiedCount == 0 {
		return false, errors.New("Unable to update task")
	}

	return true, nil	
}

// delete task
func DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": taskID, "userId": userID}

	deletionResp, err := database.TasksCollection.DeleteOne(ctx, filter)

	if deletionResp.DeletedCount == 0 {
		return false, errors.New("Task not found")
	}

	if err != nil {
		return false, errors.New("Unable to delete task")
	}

	return true, nil
}