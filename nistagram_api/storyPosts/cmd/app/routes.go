package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/highlight/", app.getAllHighlights).Methods("GET")
	r.HandleFunc("/api/highlight/{id}", app.findHighlightByID).Methods("GET")
	r.HandleFunc("/api/highlight/", app.insertHighlight).Methods("POST")
	r.HandleFunc("/api/highlight/{id}", app.deleteHighlight).Methods("DELETE")

	r.HandleFunc("/api/storyPost/", app.getAllStoryPosts).Methods("GET")
	r.HandleFunc("/api/storyPost/{id}", app.findStoryPostByID).Methods("GET")
	r.HandleFunc("/api/storyPost/", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/api/storyPost/{id}", app.deletePost).Methods("DELETE")

	return r
}
