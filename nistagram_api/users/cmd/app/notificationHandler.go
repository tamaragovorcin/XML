package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	dtos2 "users/pkg/dtos"
	"users/pkg/models"
)

func (app *application) getAllNotification(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.notification.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Movies have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) getPostNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allNotifications,_ := app.notificationContent.GetAll()
	var listNotifications []dtos2.NotificationContentDTO
	for _, settingsItem := range allNotifications {
		if settingsItem.Subject.Hex()==intVar.Hex() &&( settingsItem.Posted == "Feed Post" ||
			settingsItem.Posted == "Album Feed Post" || settingsItem.Posted =="Story Post" || settingsItem.Posted =="Album Story Post"){
			listNotifications = append(listNotifications,toResponseNotificationContent(app,settingsItem))
		}
	}


	b, err := json.Marshal(listNotifications)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) getCommentNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	allNotifications,_ := app.notificationContent.GetAll()
	var listNotifications []dtos2.NotificationContentDTO
	for _, settingsItem := range allNotifications {
		if settingsItem.Subject.Hex()==intVar.Hex() && settingsItem.Posted != "Feed Post" &&
			settingsItem.Posted!= "Album Feed Post" && settingsItem.Posted!="Story Post" && settingsItem.Posted!="Album Story Post"{
			listNotifications = append(listNotifications,toResponseNotificationContent(app,settingsItem))
		}
	}


	b, err := json.Marshal(listNotifications)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) insertNotification(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Notifications
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.notification.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)
}
func toResponseNotificationContent(app *application,user models.NotificationContent) dtos2.NotificationContentDTO {
	userId := user.Object
	m, _ := app.users.FindByID(userId)

	return dtos2.NotificationContentDTO{
		Username: m.ProfileInformation.Username,
		Posted: user.Posted,
	}
}
func insertNotificationsForUser(app *application,logged string) models.Notifications {
	userId, _ := primitive.ObjectIDFromHex(logged)
	var settings = models.Notifications{
		User :userId,
		NotificationsComments: true,
		NotificationsMessages: true,
	}

	insertResult, _ := app.notification.Insert(settings)

	idMarshaled, _ := json.Marshal(insertResult.InsertedID)

	stringId := string(idMarshaled)
	stringId = stringId[1:]
	stringId = stringId[:len(stringId)-1]

	primitiveId,_ :=primitive.ObjectIDFromHex(stringId)
	settingsInserted, _ :=app.notification.FindByID(primitiveId)
	return *settingsInserted
}
func (app *application) sendNotificationPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("////////////////////////////////////")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	postType := vars["postType"]
	allNotifications,_ := app.notificationForUser.GetAll()
	allNotificationSettings, _ := app.notificationForUser.GetAll()
	for _, settingsItem := range allNotifications {
		object :=settingsItem.Object
		objectHex:= object.Hex()
		if objectHex== userId {
			var settings = models.NotificationContent{
				Subject: settingsItem.Subject,
				Object: settingsItem.Object,
				Posted: postType,
			}
			fmt.Println(postType)
			if postType == "Feed Post" || postType =="Album Feed Post"{
				for _,set := range  allNotificationSettings{
					if set.Subject.Hex() == settingsItem.Subject.Hex() && set.Object.Hex() == settingsItem.Object.Hex() && set.Posts == true{
						insertResult, err := app.notificationContent.Insert(settings)
						if err != nil {
							app.serverError(w, err)
						}

						app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)

					}
				}
			}else if postType == "Story Post" || postType == "Album Story Post"{
				for _,set := range  allNotificationSettings{
					if set.Subject.Hex() == settingsItem.Subject.Hex() && set.Object.Hex() == settingsItem.Object.Hex() && set.Stories == true{
						insertResult, err := app.notificationContent.Insert(settings)
						if err != nil {
							app.serverError(w, err)
						}

						app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)

					}
				}
			}


		}

	}

}
func (app *application) sendNotificationComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("////////////////////////////////////")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	writerId := vars["writer"]
	userId := vars["user"]
	content := vars["content"]
	writerIdPrimitive, _ := primitive.ObjectIDFromHex(writerId)

	allNotifications,_ := app.notification.GetAll()
	for _, settingsItem := range allNotifications {
		object :=settingsItem.User
		objectHex:= object.Hex()
		if objectHex== userId {
			var settings = models.NotificationContent{
				Subject: settingsItem.User,
				Object: writerIdPrimitive,
				Posted: content,
			}
			if settingsItem.NotificationsComments{
					insertResult, err := app.notificationContent.Insert(settings)
						if err != nil {
							app.serverError(w, err)
						}
						app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)


			}


		}

	}

}
func (app *application) deleteNotification(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.notification.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}
func (app *application) updateNotifications(w http.ResponseWriter, r *http.Request) {
	var m dtos2.NotificationSettingsDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allNotifications,_ := app.notification.GetAll()

	var found = false
	var notification models.Notifications
	for _, settingsItem := range allNotifications {
		subject :=settingsItem.User
		subjectHex:= subject.Hex()
		if subjectHex== m.User {
			found = true
			notification = settingsItem

		}

	}
	if found == true{
		sb := m.User
		sb = sb[1:]
		sb = sb[:len(sb)-1]
		Subject, err := primitive.ObjectIDFromHex(m.User)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}


		var UserNotification = models.Notifications{
			Id : notification.Id,
			User: Subject,
			NotificationsMessages: m.Messages,
			NotificationsComments: m.Comments,
		}
		insertResult, _ := app.notification.Update(UserNotification)
		app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
	}else{

		app.infoLog.Printf("Notifications not found")

	}


	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) turnOnNotificationsForUser(w http.ResponseWriter, r *http.Request) {
	var m dtos2.NotificationDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	allNotifications,_ := app.notificationForUser.GetAll()

	var found = false
	var notification models.NotificationForUser
	for _, settingsItem := range allNotifications {
		subject :=settingsItem.Subject
		subjectHex:= subject.Hex()
		object :=settingsItem.Object
		objectHex:= object.Hex()
		if subjectHex== m.Subject && objectHex == m.Object {
			found = true
			notification = settingsItem

		}

	}
fmt.Println(notification.Id)
	if found == true{
		sb := m.Subject
		sb = sb[1:]
		sb = sb[:len(sb)-1]
		Subject, err := primitive.ObjectIDFromHex(m.Subject)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		obj := m.Object
		obj = obj[1:]
		obj = obj[:len(obj)-1]
		Object, err := primitive.ObjectIDFromHex(m.Object)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(m.Posts)
		fmt.Println(m.Stories)

		var UserNotification = models.NotificationForUser{
			Id : notification.Id,
			Subject: Subject,
			Object: Object,
			Posts: m.Posts,
			Stories: m.Stories,
		}
		insertResult, _ := app.notificationForUser.Update(UserNotification)
		app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
	}else{
		sb := m.Subject
		sb = sb[1:]
		sb = sb[:len(sb)-1]
		Subject, err := primitive.ObjectIDFromHex(m.Subject)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		obj := m.Object
		obj = obj[1:]
		obj = obj[:len(obj)-1]
		Object, err := primitive.ObjectIDFromHex(m.Object)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		var UserNotification = models.NotificationForUser{
			Subject: Subject,
			Object: Object,
			Posts: m.Posts,
			Stories: m.Stories,
		}
		insertResult, _ := app.notificationForUser.Insert(UserNotification)
		app.infoLog.Printf("New user have been created, id=%s", insertResult.InsertedID)

	}


	if err != nil {
		app.serverError(w, err)
	}

}