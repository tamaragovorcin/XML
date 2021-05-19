package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/follows/", app.getAllFollows).Methods("GET")
	r.HandleFunc("/api/follows/{id}", app.findByIDFollow).Methods("GET")
	r.HandleFunc("/api/follows/", app.insertFollow).Methods("POST")
	r.HandleFunc("/api/follows/{id}", app.deleteFollow).Methods("DELETE")

	return r
}
