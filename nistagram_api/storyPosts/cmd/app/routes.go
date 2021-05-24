package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/highlight/", app.getAllHighlights).Methods("GET")
	r.HandleFunc("/highlight/{id}", app.findHighlightByID).Methods("GET")
	r.HandleFunc("/highlight/", app.insertHighlight).Methods("POST")
	r.HandleFunc("/highlight/{id}", app.deleteHighlight).Methods("DELETE")

	r.HandleFunc("/", app.getAllStoryPosts).Methods("GET")
	r.HandleFunc("/{id}", app.findStoryPostByID).Methods("GET")
	r.HandleFunc("/", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/albumStory/", app.getAllStory).Methods("GET")
	r.HandleFunc("/albumStory/{id}", app.findAlbumStoryByID).Methods("GET")
	r.HandleFunc("/albumStory/", app.insertAlbumStory).Methods("POST")
	r.HandleFunc("/albumStory/{id}", app.deleteStory).Methods("DELETE")

	return r
}
