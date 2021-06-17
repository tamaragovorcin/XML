package main

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (app *application) removeEverythingFromUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(id)
	app.infoLog.Printf(id)

	removeFromStories(userIdPrimitive,app)
	removeFromStoryAlbums(userIdPrimitive,app)

	removeFromHighlights(userIdPrimitive,app)
	removeFromHighlightsAlbum(userIdPrimitive,app)

}

func removeFromHighlightsAlbum(idPrimitive primitive.ObjectID, app *application) {
	allAlbums,_ := app.highlightsAlbum.All()
	for _,album := range allAlbums {
		if album.User.Hex()==idPrimitive.Hex() {
			_, _ = app.highlightsAlbum.Delete(album.Id.Hex())
		}
	}
}


func removeFromHighlights(idPrimitive primitive.ObjectID, app *application) {
	allHighlights,_ := app.highlights.All()
	for _,highlight := range allHighlights {
		if highlight.User.Hex()==idPrimitive.Hex() {
			_, _ = app.highlights.Delete(highlight.Id.Hex())
		}
	}
}

func removeFromStoryAlbums(idPrimitive primitive.ObjectID, app *application) {
	allAlbums,_ := app.albumStories.All()
	for _,album := range allAlbums {
		if album.Post.User.Hex()==idPrimitive.Hex() {
			_, _ = app.albumStories.Delete(album.Id.Hex())
		}
	}
}

func removeFromStories(idPrimitive primitive.ObjectID, app *application) {
	allStories, _ := app.storyPosts.All()
	for _, story := range allStories {
		if story.Post.User.Hex() == idPrimitive.Hex() {
			_, _ = app.storyPosts.Delete(story.Id.Hex())
		}
	}
}