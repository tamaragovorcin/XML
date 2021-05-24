package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/notification/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/notification/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/notification/", app.insertNotification).Methods("POST")
	r.HandleFunc("/notification/{id}", app.deleteNotification).Methods("DELETE")

	r.HandleFunc("/settings/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/settings/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/settings/", app.insertNotification).Methods("POST")
	r.HandleFunc("/settings/{id}", app.deleteNotification).Methods("DELETE")

	return r
}
