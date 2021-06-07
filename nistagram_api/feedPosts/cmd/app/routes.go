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
	r.HandleFunc("/albumFeed/", app.insertAlbumFeed).Methods("POST")
	r.HandleFunc("/albumFeed/{id}", app.deleteAlbumFeed).Methods("DELETE")
	r.HandleFunc("/api/feedAlbum/usersAlbums/{userIdd}", app.getUsersFeedAlbums).Methods("GET")
	r.HandleFunc("/api/albumFeed/homePage/{userId}", app.getAlbumsForHomePage).Methods("GET")

	r.HandleFunc("/api/albumFeed/like/", app.likeTheFeedAlbum).Methods("POST")
	r.HandleFunc("/api/albumFeed/dislike/", app.dislikeTheFeedAlbum).Methods("POST")
	r.HandleFunc("/api/albumFeed/comment/", app.commentTheFeedAlbum).Methods("POST")
	r.HandleFunc("/api/albumFeed/likes/{postId}", app.getlikesFeedAlbum).Methods("GET")
	r.HandleFunc("/api/albumFeed/dislikes/{postId}", app.getdislikesFeedAlbum).Methods("GET")
	r.HandleFunc("/api/albumFeed/comments/{postId}", app.getcommentsFeedAlbum).Methods("GET")
	r.HandleFunc("/api/albumFeed/searchByLocation/{country}/{city}/{street}", app.getFeedAlbumsByLocation).Methods("GET")
	r.HandleFunc("/api/albumFeed/searchByHashTags/", app.getFeedAlbumsByHashTags).Methods("POST")
	r.HandleFunc("/api/albumFeed/searchByTags/{userId}", app.getFeedAlbumsByTags).Methods("GET")

	r.HandleFunc("/api/image/{userIdd}/{feedId}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/feed/usersImages/{userIdd}", app.getUsersFeedPosts).Methods("GET")
	r.HandleFunc("/api/feed/searchByLocation/{country}/{city}/{street}", app.getFeedPostsByLocation).Methods("GET")
	r.HandleFunc("/api/feed/searchByHashTags/", app.getFeedPostsByHashTags).Methods("POST")
	r.HandleFunc("/api/feed/homePage/{userId}", app.getPhototsForHomePage).Methods("GET")
	r.HandleFunc("/api/feed/like/", app.likeTheFeedPost).Methods("POST")
	r.HandleFunc("/api/feed/dislike/", app.dislikeTheFeedPost).Methods("POST")
	r.HandleFunc("/api/feed/comment/", app.commentTheFeedPost).Methods("POST")
	r.HandleFunc("/api/feed/likes/{postId}", app.getlikesFeedPost).Methods("GET")
	r.HandleFunc("/api/feed/dislikes/{postId}", app.getdislikesFeedPost).Methods("GET")
	r.HandleFunc("/api/feed/comments/{postId}", app.getcommentsFeedPost).Methods("GET")
	r.HandleFunc("/api/feed/searchByTags/{userId}", app.getFeedPostByTags).Methods("GET")


	r.HandleFunc("/api/collection/allData/{userId}", app.insertAllDataCollection).Methods("POST")
	r.HandleFunc("/collection/", app.getAllCollections).Methods("GET")
	r.HandleFunc("/api/collection/{userId}", app.insertCollection).Methods("POST")
	r.HandleFunc("/api/collection/{id}", app.deleteCollection).Methods("DELETE")
	r.HandleFunc("/api/collection/user/{userId}", app.getUsersCollections).Methods("GET")
	r.HandleFunc("/api/collection/addPost/", app.insetPostInCollection).Methods("POST")


	r.HandleFunc("/api/video/{userId}", app.uploadFile).Methods("POST")
	r.HandleFunc("/api/feed/usersVideos/{userId}", app.GetVideo).Methods("GET")
	return r
}