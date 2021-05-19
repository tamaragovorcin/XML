package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/comment/", app.getAllComments).Methods("GET")
	r.HandleFunc("/api/comment/{id}", app.findCommentByID).Methods("GET")
	r.HandleFunc("/api/comment/", app.insertComment).Methods("POST")
	r.HandleFunc("/api/comment/{id}", app.deleteComment).Methods("DELETE")

	r.HandleFunc("/api/content/", app.getAllContents).Methods("GET")
	r.HandleFunc("/api/content/{id}", app.findContentByID).Methods("GET")
	r.HandleFunc("/api/content/", app.insertContent).Methods("POST")
	r.HandleFunc("/api/content/{id}", app.deleteContent).Methods("DELETE")

	r.HandleFunc("/api/feedPost/", app.getAllFeedPosts).Methods("GET")
	r.HandleFunc("/api/feedPost/{id}", app.findFeedPostByID).Methods("GET")
	r.HandleFunc("/api/feedPost/", app.insertFeedPost).Methods("POST")
	r.HandleFunc("/api/feedPost/{id}", app.deleteFeedPost).Methods("DELETE")

	r.HandleFunc("/api/post/", app.getAllPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", app.findPostByID).Methods("GET")
	r.HandleFunc("/api/post/", app.insertPost).Methods("POST")
	r.HandleFunc("/api/post/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/api/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/api/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/api/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/api/location/{id}", app.deleteLocation).Methods("DELETE")

	return r
}
