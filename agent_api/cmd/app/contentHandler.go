package main

import (
	"AgentApp/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) getAllContents(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	contents, err := app.contents.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(contents)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("contents have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findContentByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.contents.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("content not found")
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

	app.infoLog.Println("Have been found a content")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertContent(w http.ResponseWriter, r *http.Request) {
	// Define booking model
	var m models.Content
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new booking
	insertResult, err := app.contents.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New Content have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteContent(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.contents.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d Content(s)", deleteResult.DeletedCount)
}
