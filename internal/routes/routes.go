package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// mount sub routers
	r.PathPrefix("/auth").Handler(http.StripPrefix("/auth", AuthRoutes()))

	return r
}