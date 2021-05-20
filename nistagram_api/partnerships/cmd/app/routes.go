package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.getAllPartnerships).Methods("GET")
	r.HandleFunc("/{id}", app.findPartnershipByID).Methods("GET")
	r.HandleFunc("/", app.insertPartnership).Methods("POST")
	r.HandleFunc("/{id}", app.deletePartnership).Methods("DELETE")

	return r
}
