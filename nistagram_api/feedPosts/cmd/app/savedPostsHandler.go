package main

import (
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (app *application) getAllSavedPosts(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.savedPosts.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Contents have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findSavedPostByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.savedPosts.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Booking not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert booking to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a booking")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertSavedPost(w http.ResponseWriter, r *http.Request) {
	// Define booking model
	var m models.SavedPost
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new booking
	insertResult, err := app.savedPosts.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteSavedPost(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.savedPosts.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}
func (app *application) getUsersSavedPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allCollection,_ := app.collections.All()
	allSavedPosts, _ :=app.savedPosts.All()
	usersCollections,err :=findCollectionsByUserId(allCollection,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.UserCollectionsDTO{}
	for _, feedPost := range usersCollections {

		usersSavedPosts, err := findUsersSavedPosts(allSavedPosts,userIdPrimitive)
		if err != nil {
			app.serverError(w, err)
		}

		feedPostResponse = append(feedPostResponse, collectionPostsToResponse(feedPost, usersSavedPosts))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func collectionPostsToResponse(collection models.Collection, posts []models.SavedPost) dtos.UserCollectionsDTO {
	return dtos.UserCollectionsDTO{
		Id: collection.Id,
		Name: collection.Name,
		SavedPosts: posts,
	}
}

func findUsersSavedPosts(posts []models.SavedPost, idPrimitive primitive.ObjectID) ([]models.SavedPost, error){
	feedPostUser := []models.SavedPost{}

	for _, feedPost := range posts {
		if	feedPost.User.String()==idPrimitive.String() {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}
