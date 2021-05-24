package main

import (
	"encoding/json"
	"net/http"

	"follows/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) getAllReports(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.reports.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Reports have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findReportByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.reports.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Repot not found")
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

	app.infoLog.Println("Have been found a report")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertReport(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Report
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.reports.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteReport(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.reports.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d report(s)", deleteResult.DeletedCount)
}