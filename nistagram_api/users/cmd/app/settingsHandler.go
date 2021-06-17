package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
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
func (app *application) findBlockedUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allSettings,_ := app.settings.GetAll()
	var listBlockedUsers []primitive.ObjectID
	for _, settingsItem := range allSettings {
		if settingsItem.User.Hex()==intVar.Hex() {
			listBlockedUsers = settingsItem.Blocked
		}
	}
	var listOfUsers []models.User
	for _,user := range  listBlockedUsers{
		us, _ := app.users.FindByID(user)
		listOfUsers = append(listOfUsers, *us)
	}
	response := []dtos2.BlockedUserDTO{}
	for _,blocked := range  listOfUsers{
			response = 	append(response,blockedToResponse(blocked))
	}
	b, err := json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func blockedToResponse(user models.User) dtos2.BlockedUserDTO {
	return dtos2.BlockedUserDTO{
		Id: user.Id,
		Username: user.ProfileInformation.Username,
	}
}
func (app *application) addSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	allSettings,_ := app.settings.GetAll()
	allNotifications,_ := app.notification.GetAll()
	settings := getUsersSettings(app,allSettings,userId)
	notifications := getUsersNotifications(app, allNotifications, userId)
	fmt.Println(notifications)
	b, err := json.Marshal(settings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) findMutedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allSettings,_ := app.settings.GetAll()

	usersCloseFriens := getMutedUsers(allSettings,intVar)
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
func (app *application) checkIfUserAllowsTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allSettings,_ := app.settings.GetAll()

	allows := getAllowTags(allSettings,intVar)

	b, err := json.Marshal(allows)

	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) checkIfUserIsMuted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	subjectId := vars["subjectId"]
	objectId := vars["objectId"]
	subjVar, err := primitive.ObjectIDFromHex(subjectId)
	objVar, err := primitive.ObjectIDFromHex(objectId)
	allSettings,_ := app.settings.GetAll()

	mutedList := getMutedUsers(allSettings,subjVar)

	found := false
	list  := strings.Split(mutedList, ",")
	for _, muted := range list {
		if muted == objVar.Hex() {
			found = true
		}

	}
	b, err := json.Marshal(found)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) checkIfUserIsBlocked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	subjectId := vars["subjectId"]
	objectId := vars["objectId"]
	subjVar, err := primitive.ObjectIDFromHex(subjectId)
	objVar, err := primitive.ObjectIDFromHex(objectId)
	allSettings,_ := app.settings.GetAll()

	blockedList1 := getBlockedUsers(allSettings,subjVar)

	found := false
	list  := strings.Split(blockedList1, ",")
	for _, muted := range list {
		if muted == objVar.Hex() {
			found = true
		}

	}
	blockedList2 := getBlockedUsers(allSettings,objVar)
	list1  := strings.Split(blockedList2, ",")
	for _, muted := range list1 {
		if muted == subjVar.Hex() {
			found = true
		}

	}
	b, err := json.Marshal(found)
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
func getBlockedUsers(settings []models.Settings, user primitive.ObjectID) string {
	var listCloseFriends []primitive.ObjectID
	listCloseFriendsString :=""
	for _, settingsItem := range settings {
		if settingsItem.User.Hex()==user.Hex() {
			listCloseFriends = settingsItem.Blocked
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
func getMutedUsers(settings []models.Settings, user primitive.ObjectID) string {
	var listCloseFriends []primitive.ObjectID
	listCloseFriendsString :=""
	for _, settingsItem := range settings {
		if settingsItem.User.Hex()==user.Hex() {
			listCloseFriends = settingsItem.Muted
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
func getAllowTags(settings []models.Settings, user primitive.ObjectID) bool {
	var allows bool
	fmt.Println(user.Hex())
	for _, settingsItem := range settings {
		fmt.Println(settingsItem.User.Hex())

		if settingsItem.User.Hex()==user.Hex() {
			fmt.Println(settingsItem.AllowTags)
			allows = settingsItem.AllowTags
		}
	}
	return allows
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
func getUsersNotifications(app *application,notifications []models.Notifications, logged string) models.Notifications {
	for _, settingsItem := range notifications {
		user :=settingsItem.User
		settinsUser:= user.Hex()
		if settinsUser==logged {
			return settingsItem
		}
	}
	usersSettings := insertNotificationsForUser(app,logged)
	return usersSettings
}

func (app *application) getUsersPrivacySettings(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	allSettings,_ := app.settings.GetAll()

	settings := getUsersSettings(app,allSettings,userId)

	b, err := json.Marshal(settings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
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
func (app *application) changePrivacySettings(w http.ResponseWriter, r *http.Request) {

	var m dtos2.SettingsDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Println(m.AllowTags)
	fmt.Println(m.AcceptMessages)
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.UserId)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : m.AllowTags,
		AcceptMessages : m.AcceptMessages,
		Muted: usersSettings.Muted,
		Blocked: usersSettings.Blocked,
		CloseFriends : usersSettings.CloseFriends,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) muteUser(w http.ResponseWriter, r *http.Request) {

	var m dtos2.MuteDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.Subject)
	muted, _ := primitive.ObjectIDFromHex(m.Object)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: append(usersSettings.Muted,muted),
		Blocked: usersSettings.Blocked,
		CloseFriends : usersSettings.CloseFriends,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) blockUser(w http.ResponseWriter, r *http.Request) {
	var m dtos2.MuteDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.Subject)
	blocked, _ := primitive.ObjectIDFromHex(m.Object)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: usersSettings.Muted,
		Blocked: append(usersSettings.Muted,blocked),
		CloseFriends : usersSettings.CloseFriends,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}

	resp, err := http.Get("http://localhost:4005/api/deleteFollow/"+m.Subject+"/"+m.Object)
	log.Println("unable to encode image.", resp)
	if err != nil {
		log.Fatalln(err)
	}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) unmuteUser(w http.ResponseWriter, r *http.Request) {

	var m dtos2.MuteDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.Subject)
	closeFriendsId, _ := primitive.ObjectIDFromHex(m.Object)
	newMutedList :=removeUserFromMutedList(usersSettings,closeFriendsId)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: newMutedList,
		Blocked: usersSettings.Blocked,
		CloseFriends : usersSettings.CloseFriends,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) unblockUser(w http.ResponseWriter, r *http.Request) {
	var m dtos2.MuteDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allSettings,_ := app.settings.GetAll()
	usersSettings := getUsersSettings(app,allSettings,m.Subject)
	closeFriendsId, _ := primitive.ObjectIDFromHex(m.Object)
	newBlockedList :=removeUserFromBlockList(usersSettings,closeFriendsId)
	var settingsUpdate = models.Settings{
		Id : usersSettings.Id,
		User : usersSettings.User,
		AllowTags : usersSettings.AllowTags,
		AcceptMessages : usersSettings.AcceptMessages,
		Muted: usersSettings.Muted,
		Blocked: newBlockedList,
		CloseFriends : usersSettings.CloseFriends,
	}


	insertResult, err := app.settings.Update(settingsUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func removeUserFromBlockList(settings models.Settings, id primitive.ObjectID) []primitive.ObjectID {
	listNew :=[]primitive.ObjectID{}
	for _, friend := range settings.Blocked {
		if friend.String()!=id.String() {
			listNew = append(listNew, friend)
		}
	}
	return listNew
}

func removeUserFromMutedList(settings models.Settings, id primitive.ObjectID) []primitive.ObjectID {
	listNew :=[]primitive.ObjectID{}
	for _, friend := range settings.Muted {
		if friend.String()!=id.String() {
			listNew = append(listNew, friend)
		}
	}
	return listNew
}

func removeFromMuted(id primitive.ObjectID, app *application) {
	allSettings,_ := app.settings.GetAll()
	for _,usersSettings := range allSettings {
		for _,mute := range usersSettings.Muted {
			if mute.Hex()==id.Hex() {
				newMutedList :=removeUserFromMutedList(usersSettings,id)
				var settingsUpdate = models.Settings{
					Id : usersSettings.Id,
					User : usersSettings.User,
					AllowTags : usersSettings.AllowTags,
					AcceptMessages : usersSettings.AcceptMessages,
					Muted: newMutedList,
					Blocked: usersSettings.Blocked,
					CloseFriends : usersSettings.CloseFriends,
				}
				_, _ = app.settings.Update(settingsUpdate)
			}
		}
	}
}



func removeFromBlocked(id primitive.ObjectID, app *application) {
	allSettings,_ := app.settings.GetAll()
	for _,usersSettings := range allSettings {
		for _,mute := range usersSettings.Blocked {
			if mute.Hex()==id.Hex() {
					newBlockedList :=removeUserFromBlockList(usersSettings,id)
				var settingsUpdate = models.Settings{
					Id : usersSettings.Id,
					User : usersSettings.User,
					AllowTags : usersSettings.AllowTags,
					AcceptMessages : usersSettings.AcceptMessages,
					Muted: usersSettings.Muted,
					Blocked: newBlockedList,
					CloseFriends : usersSettings.CloseFriends,
				}
				_, _ = app.settings.Update(settingsUpdate)
			}
		}
	}
}

func removeFromCloseFriends(id primitive.ObjectID, app *application) {
	allSettings,_ := app.settings.GetAll()
	for _,usersSettings := range allSettings {
		for _,mute := range usersSettings.CloseFriends {
			if mute.Hex()==id.Hex() {
				newCloseFriendsList :=removeCloseFriends(usersSettings,id)
				var settingsUpdate = models.Settings{
					Id : usersSettings.Id,
					User : usersSettings.User,
					AllowTags : usersSettings.AllowTags,
					AcceptMessages : usersSettings.AcceptMessages,
					Muted: usersSettings.Muted,
					Blocked: usersSettings.Blocked,
					CloseFriends : newCloseFriendsList,
				}
				_, _ = app.settings.Update(settingsUpdate)
			}
		}
	}
}
