package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/message/", IsAuthorized(app.getAllMessages)).Methods("GET")
	r.HandleFunc("/message/{id}", IsAuthorized(app.findMessageByID)).Methods("GET")
	r.HandleFunc("/message/", IsAuthorized(app.insertMessage)).Methods("POST")
	r.HandleFunc("/message/{id}", IsAuthorized(app.deleteMessage)).Methods("DELETE")

	r.HandleFunc("/chat/", IsAuthorized(app.getAllChats)).Methods("GET")
	r.HandleFunc("/chat/{id}", IsAuthorized(app.findChatByID)).Methods("GET")
	r.HandleFunc("/chat/", IsAuthorized(app.insertChat)).Methods("POST")
	r.HandleFunc("/chat/{id}", IsAuthorized(app.deleteChat)).Methods("DELETE")

	r.HandleFunc("/disposableImage/", IsAuthorized(app.getAllDisposableImages)).Methods("GET")
	r.HandleFunc("/disposableImage/{id}", IsAuthorized(app.findDisposableImageByID)).Methods("GET")
	r.HandleFunc("/disposableImage/", IsAuthorized(app.insertDisposableImage)).Methods("POST")
	r.HandleFunc("/disposableImage/{id}", IsAuthorized(app.deleteDisposableImage)).Methods("DELETE")

	r.HandleFunc("/api/send", IsAuthorized(app.sendMessage)).Methods("POST")
	r.HandleFunc("/api/send/post", IsAuthorized(app.sendPostMessage)).Methods("POST")
	r.HandleFunc("/api/send/disposableImage/{sender}/{receiver}", IsAuthorized(app.sendDisposableImage)).Methods("POST")

	r.HandleFunc("/api/getMessages/{sender}/{receiver}", IsAuthorized(app.getMessages)).Methods("GET")
	r.HandleFunc("/api/disposableImage/file/{path}", IsAuthorized(app.getDisposableImage)).Methods("GET")
	r.HandleFunc("/api/deleteChat/{sender}/{receiver}", IsAuthorized(app.deleteChatBetweenUsers)).Methods("GET")
	r.HandleFunc("/api/isChatDeleted/{sender}/{receiver}", IsAuthorized(app.isChatDeleted)).Methods("GET")
	r.HandleFunc("/api/openDisposable/{id}", IsAuthorized(app.openDisposableImage)).Methods("GET")

	return r
}
