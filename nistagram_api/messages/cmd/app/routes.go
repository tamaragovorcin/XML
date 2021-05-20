package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/message/", app.getAllMessages).Methods("GET")
	r.HandleFunc("/message/{id}", app.findMessageByID).Methods("GET")
	r.HandleFunc("/message/", app.insertMessage).Methods("POST")
	r.HandleFunc("/message/{id}", app.deleteMessage).Methods("DELETE")

	r.HandleFunc("/chat/", app.getAllChats).Methods("GET")
	r.HandleFunc("/chat/{id}", app.findChatByID).Methods("GET")
	r.HandleFunc("/chat/", app.insertChat).Methods("POST")
	r.HandleFunc("/chat/{id}", app.deleteChat).Methods("DELETE")

	r.HandleFunc("/disposableImage/", app.getAllDisposableImages).Methods("GET")
	r.HandleFunc("/disposableImage/{id}", app.findDisposableImageByID).Methods("GET")
	r.HandleFunc("/disposableImage/", app.insertDisposableImage).Methods("POST")
	r.HandleFunc("/disposableImage/{id}", app.deleteDisposableImage).Methods("DELETE")

	return r
}
