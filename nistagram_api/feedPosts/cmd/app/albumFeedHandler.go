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

func (app *application) findAlbumFeedByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.albumFeeds.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("albumFeeds not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert booking to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a albumFeeds")

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
		Likes : nil,
		Dislikes: nil,
		Comments: nil,
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
