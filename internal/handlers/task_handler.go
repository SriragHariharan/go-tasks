package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriraghariharan/gotasks/internal/middleware"
	"github.com/sriraghariharan/gotasks/internal/models"
	service "github.com/sriraghariharan/gotasks/internal/services"
	"github.com/sriraghariharan/gotasks/internal/validators"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	//verify json format
	var newTask models.Task
	jsonDecodeErr := json.NewDecoder(r.Body).Decode(&newTask)

	if jsonDecodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Json format",
		})
		return
	}
	
	//verify tasks for fields
	taskValidationErr := validators.VerifyNewTask(&newTask)

	if taskValidationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": taskValidationErr.Error(),
		})
		return
	}

	//call service layer
	ctx := r.Context()

	//extract userId from context
	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "user not found in context",
		})
		return
	}

	newTask, err := service.CreateTask(ctx, newTask, userObjectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	//extract userId from context
	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "user not found in context",
		})
		return
	}

	allTasks, err := service.GetAllTasks(ctx, userObjectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allTasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	//get task id from request
	params := mux.Vars(r)
	taskID := params["id"]

	//check if taskId is a valid mongodb
	taskObjectID, err := bson.ObjectIDFromHex(taskID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid task ID",
		})
		return
	}

	ctx := r.Context()

	//extract userId from context
	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "user not found in context",
		})
		return
	}

	_, err = service.DeleteTask(ctx, taskObjectID, userObjectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task deleted successfully",
	})
}

//update task handler
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request){

	//get task id from request
	params := mux.Vars(r)
	taskID := params["id"]

	//check if taskId is a valid mongodb id
	taskObjectID, err := bson.ObjectIDFromHex(taskID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid task ID",
		})
		return
	}

	//verify json format
	var updatedTask models.Task
	jsonDecodeErr := json.NewDecoder(r.Body).Decode(&updatedTask)

	if jsonDecodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Json format",
		})
		return
	}

	//verify tasks for fields
	taskValidationErr := validators.VerifyNewTask(&updatedTask)

	if taskValidationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": taskValidationErr.Error(),
		})
		return
	}

	ctx := r.Context()

	//extract userId from context
	userObjectID, ok := ctx.Value(middleware.UserIDKey).(bson.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "user not found in context",
		})
		return
	}

	updatedTask, err = service.UpdateTask(ctx, taskObjectID, updatedTask, userObjectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}