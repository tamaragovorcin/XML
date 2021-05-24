package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"storyPosts/pkg/models"
)

func (app *application) getAllStory(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.albumStory.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("albumStory have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findAlbumStoryByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.albumStory.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("albumStory not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert movie to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a albumStory")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertAlbumStory(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.AlbumStory
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.albumStory.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New albumStory have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteStory(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.albumStory.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d albumStory(s)", deleteResult.DeletedCount)
}

