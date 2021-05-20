package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/user/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/", app.insertUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", app.deleteUser).Methods("DELETE")
	return r
}
