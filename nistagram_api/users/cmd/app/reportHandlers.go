package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"users/pkg/models"
	"net/http"
)

func (app *application) getAllReports(w http.ResponseWriter, r *http.Request) {
	chats, err := app.reports.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Reports have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findReportByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.reports.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Report not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a report")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertReports(w http.ResponseWriter, r *http.Request) {
	var m models.Report
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.reports.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New report have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.reports.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d report(s)", deleteResult.DeletedCount)
}
