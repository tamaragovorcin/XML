package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gomod/pkg/models"
	"net/http"
)

func (app *application) getAllChats(w http.ResponseWriter, r *http.Request) {
	chats, err := app.chats.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Chats have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findChatByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.chats.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Chat not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a chat")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertChat(w http.ResponseWriter, r *http.Request) {
	var m models.Chat
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.chats.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New chat have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.chats.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d chats(s)", deleteResult.DeletedCount)
}
