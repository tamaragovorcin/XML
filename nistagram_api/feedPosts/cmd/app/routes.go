package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/comment/", app.getAllComments).Methods("GET")
	r.HandleFunc("/comment/{id}", app.findCommentByID).Methods("GET")
	r.HandleFunc("/comment/", app.insertComment).Methods("POST")
	r.HandleFunc("/comment/{id}", app.deleteComment).Methods("DELETE")

	r.HandleFunc("/", app.getAllFeedPosts).Methods("GET")
	r.HandleFunc("/{id}", app.findFeedPostByID).Methods("GET")
	r.HandleFunc("/api/feedAlbum/{userId}", app.insertAlbumFeed).Methods("POST")
	r.HandleFunc("/api/feed/{userIdd}", app.insertFeedPost).Methods("POST")
	r.HandleFunc("/{id}", app.deleteFeedPost).Methods("DELETE")

	r.HandleFunc("/post/", app.getAllPosts).Methods("GET")
	r.HandleFunc("/post/{id}", app.findPostByID).Methods("GET")
	r.HandleFunc("/post/", app.insertPost).Methods("POST")
	r.HandleFunc("/post/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/location/{id}", app.deleteLocation).Methods("DELETE")


	r.HandleFunc("/albumFeed/", app.getAllAlbumFeeds).Methods("GET")
	r.HandleFunc("/albumFeed/{id}", app.findAlbumFeedByID).Methods("GET")
	r.HandleFunc("/albumFeed/", app.insertAlbumFeed).Methods("POST")
	r.HandleFunc("/albumFeed/{id}", app.deleteAlbumFeed).Methods("DELETE")

	r.HandleFunc("/collection/", app.getAllCollections).Methods("GET")
	r.HandleFunc("/collection/{id}", app.findCollectionByID).Methods("GET")
	r.HandleFunc("/collection/", app.insertCollection).Methods("POST")
	r.HandleFunc("/collection/{id}", app.deleteCollection).Methods("DELETE")

	r.HandleFunc("/api/image/{userIdd}/{feedId}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/feed/usersImages/{userIdd}", app.getUsersFeedPosts).Methods("GET")
	r.HandleFunc("/api/feed/searchByLocation/{country}/{city}/{street}", app.getFeedPostsByLocation).Methods("GET")
	r.HandleFunc("/api/feed/searchByHashTags/", app.getFeedPostsByHashTags).Methods("POST")


	return r
}
