package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/partnership/", app.getAllPartnerships).Methods("GET")
	r.HandleFunc("/api/partnership/{id}", app.findPartnershipByID).Methods("GET")
	r.HandleFunc("/api/partnership/", app.insertPartnership).Methods("POST")
	r.HandleFunc("/api/partnership/{id}", app.deletePartnership).Methods("DELETE")

	return r
}
