package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/sriraghariharan/gotasks/internal/handlers"
)

func AuthRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handler.SignupHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	return r
}