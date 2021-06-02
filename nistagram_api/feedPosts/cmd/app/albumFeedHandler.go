package main

import (
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
