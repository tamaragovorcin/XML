package main

import (
	"encoding/json"
	"net/http"

	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) getAllCollections(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	movies, err := app.collections.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(movies)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("collections have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findCollectionByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.collections.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("collections not found")
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

	app.infoLog.Println("Have been found a collections")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertCollection(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Collection
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.collections.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New collections have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteCollection(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.collections.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d collections(s)", deleteResult.DeletedCount)
}
