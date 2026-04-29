package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sriraghariharan/gotasks/internal/models"
	service "github.com/sriraghariharan/gotasks/internal/services"
	"github.com/sriraghariharan/gotasks/internal/validators"
)

func SignupHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	//get user from req
	var newUser models.User
	jsonDecodeErr := json.NewDecoder(r.Body).Decode(&newUser)

	if jsonDecodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Json format",
		})
		return
	}

	//validate user
	userValidationErr := validators.NewUserValidator(&newUser)

	if userValidationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": userValidationErr.Error(),
		})
		return
	}

	ctx := r.Context()
	
	//send user to service layer
	newUserFromDb, err := service.CreateNewUser(ctx, newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUserFromDb)
}