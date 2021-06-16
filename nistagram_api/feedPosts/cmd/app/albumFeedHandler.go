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
	listTagged := taggedUsersToPrimitiveObject(m)

	var post = models.Post{
		User : userIdPrimitive,
		DateTime : time.Now(),
		Tagged : listTagged,
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

	taggedPeople :=getTaggedPeople(feedAlbum.Post.Tagged)

	return dtos.FeedAlbumInfoDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
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
		if iAmFollowingThisUser(userId,album.Post.User.Hex()) {
			if (!iBlockedThisUser(userId, album.Post.User.Hex())) {
				if (!iMutedThisUser(userId, album.Post.User.Hex())) {

					images, err := findAlbumByPostId(allImages, album.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(album.Post.User)
					feedAlbumResponse = append(feedAlbumResponse, toResponseAlbumHomePage(album, images, userUsername))
				}
			}
		}
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

	taggedPeople :=getTaggedPeople(feedAlbum.Post.Tagged)

	return dtos.FeedAlbumInfoDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
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
		app.infoLog.Println("Feed Album not found")
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
		Id: feedAlbum.Id,
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
		Id: feedAlbum.Id,
		Dislikes:append(feedAlbum.Dislikes, m.UserId),
		Comments : feedAlbum.Comments,
		Post : post,
		Likes: feedAlbum.Likes,
	}

	insertResult, err := app.albumFeeds.Update(feedAlbumUpdate)
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
	var comment = models.Comment{
		DateTime : time.Now(),
		Content : m.Content,
		Writer: m.UserId,
	}
	var feedAlbumUpdate = models.AlbumFeed{
		Id: feedAlbum.Id,
		Dislikes:feedAlbum.Dislikes,
		Comments : append(feedAlbum.Comments, comment),
		Post : post,
		Likes: feedAlbum.Likes,
	}

	insertResult, err := app.albumFeeds.Update(feedAlbumUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getlikesFeedAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForAlbum,err :=app.albumFeeds.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range likesForAlbum.Likes {

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


	dislikesForAlbum,err :=app.albumFeeds.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	dislikesDtos := []dtos.LikeDTO{}
	for _, user := range dislikesForAlbum.Dislikes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		dislikesDtos = append(dislikesDtos, like)

	}

	usernamesMarshaled, err := json.Marshal(dislikesDtos)
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


	albumPostComments,err :=app.albumFeeds.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	commentsDtos :=getCommentDtos(albumPostComments.Comments)


	usernamesMarshaled, err := json.Marshal(commentsDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getFeedAlbumsByLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	city :=vars["city"]
	street :=vars["street"]
	allImages,_ := app.images.All()
	locationFeedAlbums, _ :=app.albumFeeds.All()

	if country!="n" || city!="n" || street!="n" {
		locationFeedAlbums,_ =findFeedAlbumsByLocation(locationFeedAlbums,country,city,street)
	}
	feedAlbumsResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range locationFeedAlbums {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}

		userUsername :=getUserUsername(album.Post.User)

		feedAlbumsResponse = append(feedAlbumsResponse, toResponseAlbumHomePage(album, images,userUsername))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func (app *application) getFeedAlbumsByTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(user)

	allImages,_ := app.images.All()
	tagsFeedAlbums, _ :=app.albumFeeds.All()

	tagsFeedAlbums =findFeedAlbumsByTags(tagsFeedAlbums,userIdPrimitive)

	feedAlbumsResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range tagsFeedAlbums {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}
		userUsername :=getUserUsername(album.Post.User)

		feedAlbumsResponse = append(feedAlbumsResponse, toResponseAlbumHomePage(album, images,userUsername))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findFeedAlbumsByTags(albums []models.AlbumFeed, idPrimitive primitive.ObjectID) []models.AlbumFeed {
	listAlbums:=[]models.AlbumFeed{}
	for _, album := range albums {
		if userIsPublic(album.Post.User)==true {

			for _, tag := range album.Post.Tagged {
				if tag.String() == idPrimitive.String() {
					listAlbums = append(listAlbums, album)
				}
			}
		}
	}
	return listAlbums
}
func findFeedAlbumsByLocation(posts []models.AlbumFeed, country string, city string, street string)([]models.AlbumFeed, error) {
	feedPostsLocation := []models.AlbumFeed{}

	for _, feedPost := range posts {
		if userIsPublic(feedPost.Post.User)==true {
			if	feedPost.Post.Location.Country==country {
				if city=="n" {
					feedPostsLocation = append(feedPostsLocation, feedPost)
				} else if feedPost.Post.Location.Town==city {
					if street== "n" {
						feedPostsLocation = append(feedPostsLocation, feedPost)
					} else if feedPost.Post.Location.Street==street {
						feedPostsLocation = append(feedPostsLocation, feedPost)
					}
				}
			}
		}

	}
	return feedPostsLocation, nil
}


func (app *application) getFeedAlbumsByHashTags(w http.ResponseWriter, r *http.Request) {
	var hashtags dtos.HashTagDTO
	err := json.NewDecoder(r.Body).Decode(&hashtags)
	if err != nil {
		app.serverError(w, err)
	}
	allImages,_ := app.images.All()
	hashTagsFeedAlbums, _ :=app.albumFeeds.All()

	if hashtags.HashTags!="n" {
		hashTagsFeedAlbums,_ =findFeedAlbumsByHashTags(hashTagsFeedAlbums,parseHashTags(hashtags.HashTags))
	}

	feedAlbumsResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range hashTagsFeedAlbums {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}

		userUsername :=getUserUsername(album.Post.User)

		feedAlbumsResponse = append(feedAlbumsResponse, toResponseAlbumHomePage(album, images,userUsername))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findFeedAlbumsByHashTags(posts []models.AlbumFeed, hashtags []string) ([]models.AlbumFeed, error) {
	feedPostsHashTags := []models.AlbumFeed{}

	for _, feedPost := range posts {
		if userIsPublic(feedPost.Post.User)==true {
			feedPostsHashTagsList := feedPost.Post.Hashtags
			if feedPostsHashTagsList != nil {
				if postContainsAllHashTags(feedPostsHashTagsList, hashtags) {
					feedPostsHashTags = append(feedPostsHashTags, feedPost)
				}
			}
		}

	}
	return feedPostsHashTags, nil
}

func (app *application) getLikedAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages, _ := app.images.All()
	allAlbums, _ := app.albumFeeds.All()
	feedAlbumsForHomePage, err := findLikedAlbumsByUser(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range feedAlbumsForHomePage {
		if iAmFollowingThisUser(userId,album.Post.User.Hex()) {
			if !iBlockedThisUser(userId, album.Post.User.Hex()) {

					images, err := findAlbumByPostId(allImages, album.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(album.Post.User)
					feedAlbumResponse = append(feedAlbumResponse, toResponseAlbumHomePage(album, images, userUsername))
				}
			}

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func (app *application) getDislikedAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages, _ := app.images.All()
	allAlbums, _ := app.albumFeeds.All()
	feedAlbumsForHomePage, err := findLikedAlbumsByUser(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.FeedAlbumInfoDTO{}
	for _, album := range feedAlbumsForHomePage {
		if iAmFollowingThisUser(userId,album.Post.User.Hex()) {
			if !iBlockedThisUser(userId, album.Post.User.Hex()) {

					images, err := findAlbumByPostId(allImages, album.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(album.Post.User)
					feedAlbumResponse = append(feedAlbumResponse, toResponseAlbumHomePage(album, images, userUsername))
				}
			}

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func findLikedAlbumsByUser(posts []models.AlbumFeed, idPrimitive primitive.ObjectID) ([]models.AlbumFeed, error){
	feedPostUser := []models.AlbumFeed{}

	for _, feedPost := range posts {

		if	userLikedThePhoto(feedPost.Likes,idPrimitive) {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}
