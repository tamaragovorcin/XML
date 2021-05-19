package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/notification/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/api/notification/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/api/notification/", app.insertNotification).Methods("POST")
	r.HandleFunc("/api/notification/{id}", app.deleteNotification).Methods("DELETE")

	return r
}
