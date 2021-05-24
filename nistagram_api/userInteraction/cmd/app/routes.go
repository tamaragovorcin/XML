package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.getAllFollows).Methods("GET")
	r.HandleFunc("/{id}", app.findByIDFollow).Methods("GET")
	r.HandleFunc("/", app.insertFollow).Methods("POST")
	r.HandleFunc("/{id}", app.deleteFollow).Methods("DELETE")

	r.HandleFunc("/report/", app.getAllReports).Methods("GET")
	r.HandleFunc("/report/{id}", app.findReportByID).Methods("GET")
	r.HandleFunc("/report/", app.insertReport).Methods("POST")
	r.HandleFunc("/report/{id}", app.deleteReport).Methods("DELETE")

	return r
}
