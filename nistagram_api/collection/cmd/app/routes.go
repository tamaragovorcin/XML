package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.all).Methods("GET")
	r.HandleFunc("/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/", app.insert).Methods("POST")
	r.HandleFunc("/{id}", app.delete).Methods("DELETE")

	return r
}
