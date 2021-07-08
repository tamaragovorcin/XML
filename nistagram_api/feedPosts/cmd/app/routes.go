package main


import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/comment/", IsAuthorized(app.getAllComments)).Methods("GET")
	r.HandleFunc("/comment/{id}", IsAuthorized(app.findCommentByID)).Methods("GET")
	r.HandleFunc("/comment/", IsAuthorized(app.insertComment)).Methods("POST")
	r.HandleFunc("/comment/{id}", IsAuthorized(app.deleteComment)).Methods("DELETE")

	r.HandleFunc("/", IsAuthorized(app.getAllFeedPosts)).Methods("GET")
	r.HandleFunc("/{id}", IsAuthorized(app.findFeedPostByID)).Methods("GET")
	r.HandleFunc("/api/feedAlbum/{userId}", IsAuthorized(app.insertAlbumFeed)).Methods("POST")
	r.HandleFunc("/api/feed/{userIdd}", IsAuthorized(app.insertFeedPost)).Methods("POST")

	r.HandleFunc("/location/", IsAuthorized(app.getAllLocations)).Methods("GET")
	r.HandleFunc("/location/{id}", IsAuthorized(app.findLocationByID)).Methods("GET")
	r.HandleFunc("/location/", IsAuthorized(app.insertLocation)).Methods("POST")
	r.HandleFunc("/location/{id}", IsAuthorized(app.deleteLocation)).Methods("DELETE")


	r.HandleFunc("/albumFeed/", IsAuthorized(app.getAllAlbumFeeds)).Methods("GET")
	r.HandleFunc("/albumFeed/", IsAuthorized(app.insertAlbumFeed)).Methods("POST")
	r.HandleFunc("/api/feedAlbum/usersAlbums/{userIdd}", IsAuthorized(app.getUsersFeedAlbums)).Methods("GET")
	r.HandleFunc("/api/albumFeed/homePage/{userId}", IsAuthorized(app.getAlbumsForHomePage)).Methods("GET")

	r.HandleFunc("/api/albumFeed/like/", IsAuthorized(app.likeTheFeedAlbum)).Methods("POST")
	r.HandleFunc("/api/albumFeed/dislike/", IsAuthorized(app.dislikeTheFeedAlbum)).Methods("POST")
	r.HandleFunc("/api/albumFeed/comment/", IsAuthorized(app.commentTheFeedAlbum)).Methods("POST")
	r.HandleFunc("/api/albumFeed/likes/{postId}", IsAuthorized(app.getlikesFeedAlbum)).Methods("GET")
	r.HandleFunc("/api/albumFeed/dislikes/{postId}", IsAuthorized(app.getdislikesFeedAlbum)).Methods("GET")
	r.HandleFunc("/api/albumFeed/comments/{postId}", IsAuthorized(app.getcommentsFeedAlbum)).Methods("GET")
	r.HandleFunc("/api/albumFeed/searchByLocation/{country}/{city}/{street}", IsAuthorized(app.getFeedAlbumsByLocation)).Methods("GET")
	r.HandleFunc("/api/albumFeed/searchByHashTags/", IsAuthorized(app.getFeedAlbumsByHashTags)).Methods("POST")
	r.HandleFunc("/api/albumFeed/searchByTags/{userId}", IsAuthorized(app.getFeedAlbumsByTags)).Methods("GET")
	r.HandleFunc("/albumFeed/liked/{userId}", IsAuthorized(app.getLikedAlbums)).Methods("GET")
	r.HandleFunc("/albumFeed/disliked/{userId}", IsAuthorized(app.getDislikedAlbums)).Methods("GET")




	r.HandleFunc("/api/image/{userIdd}/{feedId}", IsAuthorized(app.saveImage)).Methods("POST")
	r.HandleFunc("/api/feed/username/{feedId}", IsAuthorized(app.getUsername)).Methods("GET")
	r.HandleFunc("/api/album/username/{feedId}", IsAuthorized(app.getUsernameAlbum)).Methods("GET")
	r.HandleFunc("/api/feed/usersImages/{userIdd}", IsAuthorized(app.getUsersFeedPosts)).Methods("GET")
	r.HandleFunc("/api/feed/searchByLocation/{country}/{city}/{street}", IsAuthorized(app.getFeedPostsByLocation)).Methods("GET")
	r.HandleFunc("/api/feed/searchByHashTags/", IsAuthorized(app.getFeedPostsByHashTags)).Methods("POST")
	r.HandleFunc("/api/feed/homePage/{userId}", IsAuthorized(app.getPhototsForHomePage)).Methods("GET")
	r.HandleFunc("/api/feed/like/", IsAuthorized(app.likeTheFeedPost)).Methods("POST")
	r.HandleFunc("/api/feed/dislike/", IsAuthorized(app.dislikeTheFeedPost)).Methods("POST")
	r.HandleFunc("/api/feed/comment/", IsAuthorized(app.commentTheFeedPost)).Methods("POST")
	r.HandleFunc("/api/feed/likes/{postId}", IsAuthorized(app.getlikesFeedPost)).Methods("GET")
	r.HandleFunc("/api/feed/dislikes/{postId}", IsAuthorized(app.getdislikesFeedPost)).Methods("GET")
	r.HandleFunc("/api/feed/comments/{postId}", IsAuthorized(app.getcommentsFeedPost)).Methods("GET")
	r.HandleFunc("/api/feed/searchByTags/{userId}", IsAuthorized(app.getFeedPostByTags)).Methods("GET")
	r.HandleFunc("/feed/liked/{userId}", IsAuthorized(app.getLikedPhotos)).Methods("GET")
	r.HandleFunc("/feed/disliked/{userId}", IsAuthorized(app.getDislikedPhotos)).Methods("GET")

	r.HandleFunc("/api/collection/allData/{userId}", IsAuthorized(app.insertAllDataCollection)).Methods("POST")
	r.HandleFunc("/collection/", IsAuthorized(app.getAllCollections)).Methods("GET")
	r.HandleFunc("/api/collection/{userId}", IsAuthorized(app.insertCollection)).Methods("POST")
	r.HandleFunc("/api/collection/{id}", IsAuthorized(app.deleteCollection)).Methods("DELETE")
	r.HandleFunc("/api/collection/user/{userId}", IsAuthorized(app.getUsersCollections)).Methods("GET")
	r.HandleFunc("/api/collection/addPost/", IsAuthorized(app.insetPostInCollection)).Methods("POST")

	r.HandleFunc("/api/collection/user/album/{userId}", IsAuthorized(app.getUsersCollectionAlbums)).Methods("GET")
	r.HandleFunc("/api/collection/album/{userId}", IsAuthorized(app.insertCollectionAlbum)).Methods("POST")
	r.HandleFunc("/api/collection/album/addPost/", IsAuthorized(app.insetAlbumInCollectionAlbum)).Methods("POST")


	r.HandleFunc("/api/video/{userId}/{feedId}", IsAuthorized(app.saveVideo)).Methods("POST")
	r.HandleFunc("/api/feed/file/{feedId}", app.GetFileByPostId).Methods("GET")
	r.HandleFunc("/api/feed/fileMessage/{feedId}/{userId}", IsAuthorized(app.GetFileMessageByPostId)).Methods("GET")
	r.HandleFunc("/api/feedAlbum/files/{feedId}", IsAuthorized(app.GetFilesByAlbumPostId)).Methods("GET")
	r.HandleFunc("/api/feedAlbum/images/{feedId}", IsAuthorized(app.GetImagesByAlbumId)).Methods("GET")

	r.HandleFunc("/report", IsAuthorized(app.reportFeedPost)).Methods("POST")
	r.HandleFunc("/feed/reports", IsAuthorized(app.getAllFeedReports)).Methods("GET")
	r.HandleFunc("/albumFeed/reports", IsAuthorized(app.getAllAlbumReports)).Methods("GET")
	r.HandleFunc("/report/remove/{id}", IsAuthorized(app.deleteReport)).Methods("DELETE")
	r.HandleFunc("/feed/remove/{id}/{reportId}", IsAuthorized(app.deleteFeedPost)).Methods("DELETE")

	r.HandleFunc("/albumFeed/remove/{id}/{reportId}", IsAuthorized(app.deleteAlbumFeed)).Methods("DELETE")

	r.HandleFunc("/removeUserId/{id}", IsAuthorized(app.removeEverythingFromUser)).Methods("DELETE")
	r.HandleFunc("/locationOk/{userId}/{country}/{city}/{street}", IsAuthorized(app.findIfLocationIsOk)).Methods("GET")


	return r
}