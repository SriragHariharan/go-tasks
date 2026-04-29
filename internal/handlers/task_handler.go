package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sriraghariharan/gotasks/internal/models"
	service "github.com/sriraghariharan/gotasks/internal/services"
	"github.com/sriraghariharan/gotasks/internal/validators"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	newTask, err := service.CreateTask(ctx, newTask)

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