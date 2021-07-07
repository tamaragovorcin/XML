package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/story/{userId}", IsAuthorized(app.insertStoryPost)).Methods("POST")
	r.HandleFunc("/api/story/user/{userId}", IsAuthorized(app.getUsersStories)).Methods("GET")
	r.HandleFunc("/api/story/homePage/{userId}", IsAuthorized(app.getStoriesForHomePage)).Methods("GET")


	r.HandleFunc("/api/storyAlbum/{userId}", IsAuthorized(app.insertAlbumStory)).Methods("POST")


	r.HandleFunc("/api/storyAlbum/{userId}", IsAuthorized(app.insertAlbumStory)).Methods("POST")
	r.HandleFunc("/api/storyAlbum/usersAlbums/{userIdd}", IsAuthorized(app.getUsersStoryAlbums)).Methods("GET")
	r.HandleFunc("/api/storyAlbum/homePage/{userId}", IsAuthorized(app.getStoryAlbumsForHomePage)).Methods("GET")

	r.HandleFunc("/api/image/{userId}/{storyId}", IsAuthorized(app.saveImage)).Methods("POST")

	r.HandleFunc("/highlight/", IsAuthorized(app.getAllHighlights)).Methods("GET")
	r.HandleFunc("/api/highlight/{userId}", IsAuthorized(app.insertHighlight)).Methods("POST")
	r.HandleFunc("/highlight/{id}", IsAuthorized(app.deleteHighlight)).Methods("DELETE")
	r.HandleFunc("/api/highlight/user/{userId}", IsAuthorized(app.getUsersHiglights)).Methods("GET")
	r.HandleFunc("/api/highlight/addStory/", IsAuthorized(app.insetStoryInHighlight)).Methods("POST")

	r.HandleFunc("/api/highlight/album/{userId}", IsAuthorized(app.insertHighlightAlbum)).Methods("POST")
	r.HandleFunc("/api/highlight/user/album/{userId}", IsAuthorized(app.getUsersHiglightAlbums)).Methods("GET")


	r.HandleFunc("/api/", IsAuthorized(app.getAllStoryPosts)).Methods("GET")
	r.HandleFunc("/api/", IsAuthorized(app.insertStoryPost)).Methods("POST")
	r.HandleFunc("/api/{id}", IsAuthorized(app.deletePost)).Methods("DELETE")

	r.HandleFunc("/api/albumStory/", IsAuthorized(app.getAllStory)).Methods("GET")
	r.HandleFunc("/api/albumStory/", IsAuthorized(app.insertAlbumStory)).Methods("POST")
	r.HandleFunc("/api/albumStory/{id}", IsAuthorized(app.deleteStory)).Methods("DELETE")
	r.HandleFunc("/api/highlight/addStoryAlbum/", IsAuthorized(app.insetStoryAlbumInHighlight)).Methods("POST")

	r.HandleFunc("/api/story/file/{storyId}", IsAuthorized(app.GetFileByPostId)).Methods("GET")
	r.HandleFunc("/removeUserId/{id}", IsAuthorized(app.removeEverythingFromUser)).Methods("DELETE")

	r.HandleFunc("/api/story/username/{feedId}", IsAuthorized(app.getUsername)).Methods("GET")
	r.HandleFunc("/api/story/fileMessage/{feedId}/{userId}", IsAuthorized(app.GetFileMessageByPostId)).Methods("GET")

	return r
}