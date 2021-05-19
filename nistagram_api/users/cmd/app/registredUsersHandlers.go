package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"users/pkg/models"
	"net/http"
)

func (app *application) getAllRegistredUsers(w http.ResponseWriter, r *http.Request) {
	chats, err := app.registredUsers.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("RegistredUsers have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findRegistredUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.registredUsers.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("RegistredUser not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a RegistredUser")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertRegistredUser(w http.ResponseWriter, r *http.Request) {
	var m models.RegisteredUser
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.registredUsers.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New RegistredUser have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteRegistredUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.registredUsers.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d RegistredUsers(s)", deleteResult.DeletedCount)
}
