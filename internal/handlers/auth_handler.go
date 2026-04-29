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

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	//decode json to struct
	var existingUser models.User;
	jsonDecodeErr := json.NewDecoder(r.Body).Decode(&existingUser)

	if jsonDecodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Json format",
		})
		return
	}

	//validate user json for fields
	userValidationErr := validators.ValidateExistingUser(&existingUser)
	if userValidationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": userValidationErr.Error(),
		})
	}

	ctx := r.Context()
	//call the service layer
	userDetails, err := service.LoginUser(ctx, &existingUser)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDetails)
}