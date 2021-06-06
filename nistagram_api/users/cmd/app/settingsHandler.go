package main
import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gorilla/mux"
	"users/pkg/models"
)

func (app *application) getAllSettings(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.settings.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("settings have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}



func (app *application) insertSettings(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Settings
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.settings.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New Settings have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteSettings(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.settings.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d Settings(s)", deleteResult.DeletedCount)
}


func (app *application) findUserCloseFriends(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allSettings,_ := app.settings.GetAll()

	usersCloseFriens := getCloseFriends(allSettings,intVar)

	b, err := json.Marshal(usersCloseFriens)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getCloseFriends(settings []models.Settings, user primitive.ObjectID) []string {
	listCloseFriends := []primitive.ObjectID{}
	listCloseFriendsString := []string{}

	for _, settingsItem := range settings {
		if settingsItem.User==user {
			listCloseFriends = settingsItem.CloseFriends
		}
	}
	for _, closeFriends := range listCloseFriends {
		listCloseFriendsString = append(listCloseFriendsString, closeFriends.String())
	}
	return listCloseFriendsString
}
