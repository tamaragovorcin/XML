package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"partnerships/pkg/models"
)

func (app *application) getAllPartnerships(w http.ResponseWriter, r *http.Request) {
	chats, err := app.partnerships.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("partnerships have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findPartnershipByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.partnerships.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("partnership not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a partnerships")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertPartnership(w http.ResponseWriter, r *http.Request) {
	var m models.Partnership
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.partnerships.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New partnership have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deletePartnership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.partnerships.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d chats(s)", deleteResult.DeletedCount)
}
