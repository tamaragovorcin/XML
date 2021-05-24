package main

import (
	"encoding/json"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
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

func (app *application) insertAlbumFeed(w http.ResponseWriter, r *http.Request) {
	// Define booking model
	var m models.AlbumFeed
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new booking
	insertResult, err := app.albumFeeds.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New albumFeeds have been created, id=%s", insertResult.InsertedID)
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
