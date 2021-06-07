package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/story/{userId}", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/api/story/user/{userId}", app.getUsersStories).Methods("GET")
	r.HandleFunc("/api/story/homePage/{userId}", app.getStoriesForHomePage).Methods("GET")


	r.HandleFunc("/api/storyAlbum/{userId}", app.insertAlbumStory).Methods("POST")


	r.HandleFunc("/api/storyAlbum/{userId}", app.insertAlbumStory).Methods("POST")
	r.HandleFunc("/api/storyAlbum/usersAlbums/{userIdd}", app.getUsersStoryAlbums).Methods("GET")
	r.HandleFunc("/api/storyAlbum/homePage/{userId}", app.getStoryAlbumsForHomePage).Methods("GET")


	r.HandleFunc("/api/image/{userId}/{storyId}", app.saveImage).Methods("POST")

	r.HandleFunc("/highlight/", app.getAllHighlights).Methods("GET")
	r.HandleFunc("/api/highlight/{userId}", app.insertHighlight).Methods("POST")
	r.HandleFunc("/highlight/{id}", app.deleteHighlight).Methods("DELETE")
	r.HandleFunc("/api/highlight/user/{userId}", app.getUsersHiglights).Methods("GET")
	r.HandleFunc("/api/highlight/addStory/", app.insetStoryInHighlight).Methods("POST")

	r.HandleFunc("/api/highlight/album/{userId}", app.insertHighlightAlbum).Methods("POST")
	r.HandleFunc("/api/highlight/user/album/{userId}", app.getUsersHiglightAlbums).Methods("GET")


	r.HandleFunc("/api/", app.getAllStoryPosts).Methods("GET")
	r.HandleFunc("/api/", app.insertStoryPost).Methods("POST")
	r.HandleFunc("/api/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/api/albumStory/", app.getAllStory).Methods("GET")
	r.HandleFunc("/api/albumStory/", app.insertAlbumStory).Methods("POST")
	r.HandleFunc("/api/albumStory/{id}", app.deleteStory).Methods("DELETE")
	r.HandleFunc("/api/highlight/addStoryAlbum/", app.insetStoryAlbumInHighlight).Methods("POST")

	r.HandleFunc("/api/story/file/{storyId}", app.GetFileByPostId).Methods("GET")
	return r
}