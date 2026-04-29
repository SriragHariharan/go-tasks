package main

import (
	"log"
	"net/http"

	database "github.com/sriraghariharan/gotasks/internal/db"
	"github.com/sriraghariharan/gotasks/internal/routes"
)

func main() {
	//connect to database
	database.Connect()

	r := routes.SetupRoutes()
	//server
	log.Fatal(http.ListenAndServe(":4000", r))
}