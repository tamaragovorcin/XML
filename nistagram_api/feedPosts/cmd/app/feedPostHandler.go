package main

import (
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (app *application) getAllFeedPosts(w http.ResponseWriter, r *http.Request) {
	bookings, err := app.feedPosts.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Contents have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findFeedPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.feedPosts.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Booking not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a booking")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertFeedPost(w http.ResponseWriter, req *http.Request) {
	app.infoLog.Printf("888888888888888888888888888888888888888888888")

	var m dtos.FeedPostDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	var post = models.Post{
		Id : 15,
		User : m.User,
		DateTime : time.Now(),
		Tagged : m.Tagged,
		Description: m.Description,
		Hashtags: m.Hashtags,
		Location : m.Location,
		Blocked : false,
	}
	var feedPost = models.FeedPost{
		Id : 15,
		Post : post,
		Likes : nil,
		Dislikes: nil,
		Comments: nil,
	}


	insertResult, err := app.feedPosts.Insert(feedPost)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteFeedPost(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.feedPosts.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}