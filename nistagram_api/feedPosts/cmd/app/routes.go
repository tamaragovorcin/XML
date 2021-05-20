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

	r.HandleFunc("/content/", app.getAllContents).Methods("GET")
	r.HandleFunc("/content/{id}", app.findContentByID).Methods("GET")
	r.HandleFunc("/content/", app.insertContent).Methods("POST")
	r.HandleFunc("/content/{id}", app.deleteContent).Methods("DELETE")

	r.HandleFunc("/", app.getAllFeedPosts).Methods("GET")
	r.HandleFunc("/{id}", app.findFeedPostByID).Methods("GET")
	r.HandleFunc("/", app.insertFeedPost).Methods("POST")
	r.HandleFunc("/{id}", app.deleteFeedPost).Methods("DELETE")

	r.HandleFunc("/post/", app.getAllPosts).Methods("GET")
	r.HandleFunc("/post/{id}", app.findPostByID).Methods("GET")
	r.HandleFunc("/post/", app.insertPost).Methods("POST")
	r.HandleFunc("/post/{id}", app.deletePost).Methods("DELETE")

	r.HandleFunc("/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/location/{id}", app.deleteLocation).Methods("DELETE")

	return r
}
