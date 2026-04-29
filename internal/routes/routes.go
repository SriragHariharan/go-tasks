package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// mount sub routers
	r.PathPrefix("/auth").Handler(http.StripPrefix("/auth", AuthRoutes()))
	r.PathPrefix("/task").Handler(http.StripPrefix("/task", TaskRoutes()))

	return r
}