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
	r.HandleFunc("/chat/",app.insertChat).Methods("POST")
	r.HandleFunc("/chat/{id}", app.deleteChat).Methods("DELETE")

	r.HandleFunc("/disposableImage/", app.getAllDisposableImages).Methods("GET")
	r.HandleFunc("/disposableImage/{id}", app.findDisposableImageByID).Methods("GET")
	r.HandleFunc("/disposableImage/", app.insertDisposableImage).Methods("POST")
	r.HandleFunc("/disposableImage/{id}",app.deleteDisposableImage).Methods("DELETE")

	r.HandleFunc("/api/send", app.sendMessage).Methods("POST")
	r.HandleFunc("/api/send/post", app.sendPostMessage).Methods("POST")
	r.HandleFunc("/api/send/disposableImage/{sender}/{receiver}", app.sendDisposableImage).Methods("POST")

	r.HandleFunc("/api/getMessages/{sender}/{receiver}", app.getMessages).Methods("GET")
	r.HandleFunc("/api/disposableImage/file/{path}",app.getDisposableImage).Methods("GET")
	r.HandleFunc("/api/deleteChat/{sender}/{receiver}", app.deleteChatBetweenUsers).Methods("GET")
	r.HandleFunc("/api/isChatDeleted/{sender}/{receiver}", app.isChatDeleted).Methods("GET")
	r.HandleFunc("/api/openDisposable/{id}", app.openDisposableImage).Methods("GET")

	return r
}