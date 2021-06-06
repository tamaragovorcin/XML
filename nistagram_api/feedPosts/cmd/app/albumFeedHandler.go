package main

import (
	"bytes"
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (app *application) getAllAlbumFeeds(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	albumFeeds, err := app.albumFeeds.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(albumFeeds)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("albumFeeds have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func (app *application) insertAlbumFeed(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	var m dtos.FeedPostDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	var post = models.Post{
		User : userIdPrimitive,
		DateTime : time.Now(),
		Tagged : m.Tagged,
		Description: m.Description,
		Hashtags: parseHashTags(m.Hashtags),
		Location : m.Location,
		Blocked : false,
	}
	var feedPost = models.AlbumFeed{
		Post : post,
		Likes : []primitive.ObjectID{},
		Dislikes: []primitive.ObjectID{},
		Comments: []models.Comment{},
	}


	insertResult, err := app.albumFeeds.Insert(feedPost)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func (app *application) deleteAlbumFeed(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.albumFeeds.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d albumFeeds(s)", deleteResult.DeletedCount)
}

func (app *application) getUsersFeedAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages, _ := app.images.All()
	allAlbums, _ := app.albumFeeds.All()
	usersFeedAlbums, err := findFeedAlbumsByUserId(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range usersFeedAlbums {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}

		feedAlbumResponse = append(feedAlbumResponse, toResponseAlbum(album, images))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func findFeedAlbumsByUserId(albums []models.AlbumFeed, idPrimitive primitive.ObjectID) ([]models.AlbumFeed, error){
	feedAlbumsUser := []models.AlbumFeed{}

	for _, album := range albums {
		if	album.Post.User.String()==idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}
func findAlbumByPostId(images []models.Image, idFeedAlbum primitive.ObjectID) ([]string, error) {
	imageAlbumPost := []string{}

	for _, image := range images {

		if	image.PostId==idFeedAlbum {
			imageAlbumPost= append(imageAlbumPost, image.Media)
		}
	}
	return imageAlbumPost, nil
}
func toResponseAlbum(feedAlbum models.AlbumFeed, imageList []string) dtos.FeedAlbumInfoDTO {
	imagesBuffered := [][]byte{}
	for _, image2 := range imageList {
		f, _ := os.Open(image2)
		defer f.Close()
		image, _, _ := image.Decode(f)
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, image, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBuffered :=buffer.Bytes()
		imagesBuffered= append(imagesBuffered, imageBuffered)
	}


	return dtos.FeedAlbumInfoDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :feedAlbum.Post.Tagged,
		Location : locationToString(feedAlbum.Post.Location),
		Description : feedAlbum.Post.Description,
		Hashtags : hashTagsToString(feedAlbum.Post.Hashtags),
		Media : imagesBuffered,
		Username : "",

	}
}
func (app *application) getAlbumsForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages, _ := app.images.All()
	allAlbums, _ := app.albumFeeds.All()
	feedAlbumsForHomePage, err := findFeedAlbumsForHomePage(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range feedAlbumsForHomePage {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}
		userUsername :=getUserUsername(album.Post.User)
		feedAlbumResponse = append(feedAlbumResponse, toResponseAlbumHomePage(album, images,userUsername))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func toResponseAlbumHomePage(feedAlbum models.AlbumFeed, imageList []string, username string) dtos.FeedAlbumInfoDTO {
	imagesBuffered := [][]byte{}
	for _, image2 := range imageList {
		f, _ := os.Open(image2)
		defer f.Close()
		image, _, _ := image.Decode(f)
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, image, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBuffered :=buffer.Bytes()
		imagesBuffered= append(imagesBuffered, imageBuffered)
	}


	return dtos.FeedAlbumInfoDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :feedAlbum.Post.Tagged,
		Location : locationToString(feedAlbum.Post.Location),
		Description : feedAlbum.Post.Description,
		Hashtags : hashTagsToString(feedAlbum.Post.Hashtags),
		Media : imagesBuffered,
		Username : username,
	}
}

func findFeedAlbumsForHomePage(albums []models.AlbumFeed, idPrimitive primitive.ObjectID) ([]models.AlbumFeed, error) {
	feedAlbumsUser := []models.AlbumFeed{}

	for _, album := range albums {
		if	album.Post.User.String()!=idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	//usloc za pracenje!!!!!!!!!!!!!!!
	return feedAlbumsUser, nil
}
func (app *application) likeTheFeedAlbum(w http.ResponseWriter, r *http.Request) {

	var m dtos.PostReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbum, err := app.albumFeeds.FindByID(m.PostId)
	if feedAlbum == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedAlbum.Post.User,
		DateTime : feedAlbum.Post.DateTime,
		Tagged : feedAlbum.Post.Tagged,
		Description: feedAlbum.Post.Description,
		Hashtags: feedAlbum.Post.Hashtags,
		Location : feedAlbum.Post.Location,
		Blocked : feedAlbum.Post.Blocked,
	}
	var feedAlbumUpdate = models.AlbumFeed{
		Post : post,
		Likes : append(feedAlbum.Likes, m.UserId),
		Dislikes: feedAlbum.Dislikes,
		Comments: feedAlbum.Comments,
	}

	insertResult, err := app.albumFeeds.Update(feedAlbumUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) dislikeTheFeedAlbum(w http.ResponseWriter, r *http.Request) {

	var m dtos.PostReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.feedPosts.FindByID(m.PostId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedPost.Post.User,
		DateTime : feedPost.Post.DateTime,
		Tagged : feedPost.Post.Tagged,
		Description: feedPost.Post.Description,
		Hashtags: feedPost.Post.Hashtags,
		Location : feedPost.Post.Location,
		Blocked : feedPost.Post.Blocked,
	}
	var feedPostUpdate = models.FeedPost{
		Id: feedPost.Id,
		Dislikes:append(feedPost.Dislikes, m.UserId),
		Comments : feedPost.Comments,
		Post : post,
		Likes: feedPost.Likes,
	}

	insertResult, err := app.feedPosts.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) commentTheFeedAlbum(w http.ResponseWriter, r *http.Request) {

	var m dtos.CommentReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.feedPosts.FindByID(m.PostId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedPost.Post.User,
		DateTime : feedPost.Post.DateTime,
		Tagged : feedPost.Post.Tagged,
		Description: feedPost.Post.Description,
		Hashtags: feedPost.Post.Hashtags,
		Location : feedPost.Post.Location,
		Blocked : feedPost.Post.Blocked,
	}
	var comment = models.Comment{
		DateTime : time.Now(),
		Content : m.Content,
		Writer: m.UserId,
	}
	var feedPostUpdate = models.FeedPost{
		Id: feedPost.Id,
		Dislikes:feedPost.Dislikes,
		Comments : append(feedPost.Comments, comment),
		Post : post,
		Likes: feedPost.Likes,
	}

	insertResult, err := app.feedPosts.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getlikesFeedAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range likesForPost.Likes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}

	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getdislikesFeedAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range likesForPost.Dislikes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}

	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getcommentsFeedAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	commentsDtos :=getCommentDtos(likesForPost.Comments)


	usernamesMarshaled, err := json.Marshal(commentsDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}
