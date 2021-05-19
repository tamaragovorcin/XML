package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/albumStory/", app.getAllStory).Methods("GET")
	r.HandleFunc("/api/albumStory/{id}", app.findByIDStory).Methods("GET")
	r.HandleFunc("/api/albumStory/", app.insertStory).Methods("POST")
	r.HandleFunc("/api/albumStory/{id}", app.deleteStory).Methods("DELETE")

	r.HandleFunc("/api/albumFeed/", app.getAllFeed).Methods("GET")
	r.HandleFunc("/api/albumFeed/{id}", app.findByIDFeed).Methods("GET")
	r.HandleFunc("/api/albumFeed/", app.insertFeed).Methods("POST")
	r.HandleFunc("/api/albumFeed/{id}", app.deleteFeed).Methods("DELETE")

	return r
}
