package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"users/pkg/models"
)

func (app *application) getAllProfileInformation(w http.ResponseWriter, r *http.Request) {
	chats, err := app.profileInformation.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Profile information have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findProfileInformationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.profileInformation.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("ProfileInformation not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a ProfileInformation")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertProfileInformation(w http.ResponseWriter, r *http.Request) {
	var m models.ProfileInformation
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.profileInformation.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New ProfileInformation have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteProfileInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.profileInformation.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d ProfileInformation(s)", deleteResult.DeletedCount)
}
