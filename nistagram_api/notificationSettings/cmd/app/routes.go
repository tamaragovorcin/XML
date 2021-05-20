package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/", app.insertNotification).Methods("POST")
	r.HandleFunc("/{id}", app.deleteNotification).Methods("DELETE")

	return r
}
