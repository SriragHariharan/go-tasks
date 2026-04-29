package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)

	//server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello welcome to go tasks"))
}