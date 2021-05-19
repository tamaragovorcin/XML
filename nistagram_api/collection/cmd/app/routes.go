package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/collection/", app.all).Methods("GET")
	r.HandleFunc("/api/collection/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/api/collection/", app.insert).Methods("POST")
	r.HandleFunc("/api/collection/{id}", app.delete).Methods("DELETE")

	return r
}
