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

	r.HandleFunc("/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/location/{id}", app.deleteLocation).Methods("DELETE")


	r.HandleFunc("/albumFeed/", app.getAllAlbumFeeds).Methods("GET")
	r.HandleFunc("/albumFeed/", app.insertAlbumFeed).Methods("POST")
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
	r.HandleFunc("/albumFeed/liked/{userId}", app.getLikedAlbums).Methods("GET")
	r.HandleFunc("/albumFeed/disliked/{userId}", app.getDislikedAlbums).Methods("GET")




	r.HandleFunc("/api/image/{userIdd}/{feedId}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/feed/username/{feedId}", app.getUsername).Methods("GET")
	r.HandleFunc("/api/album/username/{feedId}", app.getUsernameAlbum).Methods("GET")
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
	r.HandleFunc("/feed/liked/{userId}", app.getLikedPhotos).Methods("GET")
	r.HandleFunc("/feed/disliked/{userId}", app.getDislikedPhotos).Methods("GET")

	r.HandleFunc("/api/collection/allData/{userId}", app.insertAllDataCollection).Methods("POST")
	r.HandleFunc("/collection/", app.getAllCollections).Methods("GET")
	r.HandleFunc("/api/collection/{userId}", app.insertCollection).Methods("POST")
	r.HandleFunc("/api/collection/{id}", app.deleteCollection).Methods("DELETE")
	r.HandleFunc("/api/collection/user/{userId}", app.getUsersCollections).Methods("GET")
	r.HandleFunc("/api/collection/addPost/", app.insetPostInCollection).Methods("POST")

	r.HandleFunc("/api/collection/user/album/{userId}", app.getUsersCollectionAlbums).Methods("GET")
	r.HandleFunc("/api/collection/album/{userId}", app.insertCollectionAlbum).Methods("POST")
	r.HandleFunc("/api/collection/album/addPost/", app.insetAlbumInCollectionAlbum).Methods("POST")


	r.HandleFunc("/api/video/{userId}/{feedId}", app.saveVideo).Methods("POST")
	r.HandleFunc("/api/feed/file/{feedId}", app.GetFileByPostId).Methods("GET")
	r.HandleFunc("/api/feed/fileMessage/{feedId}/{userId}", app.GetFileMessageByPostId).Methods("GET")
	r.HandleFunc("/api/feedAlbum/files/{feedId}", app.GetFilesByAlbumPostId).Methods("GET")
	r.HandleFunc("/api/feedAlbum/images/{feedId}", app.GetImagesByAlbumId).Methods("GET")

	r.HandleFunc("/report", app.reportFeedPost).Methods("POST")
	r.HandleFunc("/feed/reports", app.getAllFeedReports).Methods("GET")
	r.HandleFunc("/albumFeed/reports", app.getAllAlbumReports).Methods("GET")
	r.HandleFunc("/report/remove/{id}", app.deleteReport).Methods("DELETE")
	r.HandleFunc("/feed/remove/{id}/{reportId}", app.deleteFeedPost).Methods("DELETE")

	r.HandleFunc("/albumFeed/remove/{id}/{reportId}", app.deleteAlbumFeed).Methods("DELETE")

	r.HandleFunc("/removeUserId/{id}", app.removeEverythingFromUser).Methods("DELETE")
	r.HandleFunc("/locationOk/{userId}/{country}/{city}/{street}", app.findIfLocationIsOk).Methods("GET")


	return r
}