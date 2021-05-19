package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/message/", app.getAllMessages).Methods("GET")
	r.HandleFunc("/api/message/{id}", app.findMessageByID).Methods("GET")
	r.HandleFunc("/api/message/", app.insertMessage).Methods("POST")
	r.HandleFunc("/api/message/{id}", app.deleteMessage).Methods("DELETE")

	r.HandleFunc("/api/chat/", app.getAllChats).Methods("GET")
	r.HandleFunc("/api/chat/{id}", app.findChatByID).Methods("GET")
	r.HandleFunc("/api/chat/", app.insertChat).Methods("POST")
	r.HandleFunc("/api/chat/{id}", app.deleteChat).Methods("DELETE")

	r.HandleFunc("/api/disposableImage/", app.getAllDisposableImages).Methods("GET")
	r.HandleFunc("/api/disposableImage/{id}", app.findDisposableImageByID).Methods("GET")
	r.HandleFunc("/api/disposableImage/", app.insertDisposableImage).Methods("POST")
	r.HandleFunc("/api/disposableImage/{id}", app.deleteDisposableImage).Methods("DELETE")

	return r
}
