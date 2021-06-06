package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/story/{userId}", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/story/user/{userId}", app.getUsersStories).Methods("GET")
	r.HandleFunc("/api/story/homePage/{userId}", app.getStoriesForHomePage).Methods("GET")


	r.HandleFunc("/api/storyAlbum/{userId}", app.insertAlbumStory).Methods("POST")

	r.HandleFunc("/api/image/{userId}/{storyId}", app.saveImage).Methods("POST")

	r.HandleFunc("/highlight/", app.getAllHighlights).Methods("GET")
	r.HandleFunc("/highlight/{userId}", app.insertHighlight).Methods("POST")
	r.HandleFunc("/highlight/{id}", app.deleteHighlight).Methods("DELETE")
	r.HandleFunc("/highlight/user/{userId}", app.getUsersHiglights).Methods("GET")
	r.HandleFunc("/highlight/addStory/", app.insetStoryInHighlight).Methods("POST")


	r.HandleFunc("/api/", app.getAllStoryPosts).Methods("GET")
	r.HandleFunc("/api/", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/api/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/api/albumStory/", app.getAllStory).Methods("GET")
	r.HandleFunc("/api/albumStory/{id}", app.findAlbumStoryByID).Methods("GET")
	r.HandleFunc("/api/albumStory/", app.insertAlbumStory).Methods("POST")
	r.HandleFunc("/api/albumStory/{id}", app.deleteStory).Methods("DELETE")

	return r
}
