package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	dtos2 "users/pkg/dtos"

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

func (app *application) findUserCloseFriends(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allSettings,_ := app.settings.GetAll()

	usersCloseFriens := getCloseFriends(allSettings,intVar)
	fmt.Println("list   " + usersCloseFriens)

	b, err := json.Marshal(usersCloseFriens)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getCloseFriends(settings []models.Settings, user primitive.ObjectID) string {
	var listCloseFriends []primitive.ObjectID
	 listCloseFriendsString :=""
	for _, settingsItem := range settings {
		if settingsItem.User.Hex()==user.Hex() {
			listCloseFriends = settingsItem.CloseFriends
		}
	}

	for _, closeFriend := range listCloseFriends {
		listCloseFriendsString += closeFriend.Hex()+ ","
	}
	if listCloseFriendsString!="" {
		listCloseFriendsString = listCloseFriendsString[:len(listCloseFriendsString)-1]
	}


	return listCloseFriendsString
}

func (app *application) addUserToCloseFriends(w http.ResponseWriter, r *http.Request) {

	var m dtos2.CloseFriendsDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.IdLogged)
	closeFriendsId, _ := primitive.ObjectIDFromHex(m.IdClose)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: usersSettings.Muted,
		Blocked: usersSettings.Blocked,
		CloseFriends : append(usersSettings.CloseFriends,closeFriendsId),
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}


func (app *application) removeUserFromCloseFriends(w http.ResponseWriter, r *http.Request) {

	var m dtos2.CloseFriendsDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.IdLogged)
	closeFriendsId, _ := primitive.ObjectIDFromHex(m.IdClose)
	closeFriendsNew :=removeCloseFriends(usersSettings,closeFriendsId)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: usersSettings.Muted,
		Blocked: usersSettings.Blocked,
		CloseFriends : closeFriendsNew,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func removeCloseFriends(settings models.Settings, id primitive.ObjectID) []primitive.ObjectID {
	listNew :=[]primitive.ObjectID{}
	for _, friend := range settings.CloseFriends {
		if friend.String()!=id.String() {
			listNew = append(listNew, friend)
		}
	}
	return listNew
}

func getUsersSettings(app *application,settings []models.Settings, logged string) models.Settings {
	for _, settingsItem := range settings {
		user :=settingsItem.User
		settinsUser:= user.Hex()
		if settinsUser==logged {
			return settingsItem
		}
	}
	usersSettings := insertSettingsForUser(app,logged)
	return usersSettings
}

func insertSettingsForUser(app *application,logged string) models.Settings {
	userId, _ := primitive.ObjectIDFromHex(logged)
	var settings = models.Settings{
		User :userId,
		AllowTags : true,
		AcceptMessages : true,
		Muted: []primitive.ObjectID{},
		Blocked: []primitive.ObjectID{},
		CloseFriends : []primitive.ObjectID{},
	}
	insertResult, _ := app.settings.Insert(settings)
	idMarshaled, _ := json.Marshal(insertResult.InsertedID)

	stringId := string(idMarshaled)
	stringId = stringId[1:]
	stringId = stringId[:len(stringId)-1]

	primitiveId,_ :=primitive.ObjectIDFromHex(stringId)
	settingsInserted, _ :=app.settings.FindByID(primitiveId)
	return *settingsInserted
}
