package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/albumStory/", app.getAllStory).Methods("GET")
	r.HandleFunc("/albumStory/{id}", app.findByIDStory).Methods("GET")
	r.HandleFunc("/albumStory/", app.insertStory).Methods("POST")
	r.HandleFunc("/albumStory/{id}", app.deleteStory).Methods("DELETE")

	r.HandleFunc("/albumFeed/", app.getAllFeed).Methods("GET")
	r.HandleFunc("/albumFeed/{id}", app.findByIDFeed).Methods("GET")
	r.HandleFunc("/albumFeed/", app.insertFeed).Methods("POST")
	r.HandleFunc("/albumFeed/{id}", app.deleteFeed).Methods("DELETE")

	return r
}
