package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gomod/pkg/models"
	"net/http"
)

func (app *application) getAllMessages(w http.ResponseWriter, r *http.Request) {
	message, err := app.messages.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(message)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Messages have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findMessageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.messages.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Message not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a message")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.messages.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New message have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.messages.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d message(s)", deleteResult.DeletedCount)
}
