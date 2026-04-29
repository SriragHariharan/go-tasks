package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/sriraghariharan/gotasks/internal/handlers"
	"github.com/sriraghariharan/gotasks/internal/middleware"
)

func TaskRoutes() *mux.Router {
	r := mux.NewRouter()

	// use middleware for all routes in this router
	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/", handler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/all", handler.GetAllTasksHandler).Methods("GET")

	return r
}