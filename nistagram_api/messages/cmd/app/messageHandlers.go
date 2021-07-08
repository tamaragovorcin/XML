package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gomod/pkg/dtos"
	"gomod/pkg/models"
	"net/http"
	"strings"
	"time"

)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			fmt.Println("No Token Found")

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}


		authStringHeader := r.Header.Get("Authorization")
		if authStringHeader == "" {
			fmt.Errorf("Neki eror za auth")
		}
		authHeader := strings.Split(authStringHeader, "Bearer ")
		jwtToken := authHeader[1]

		token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("luna") , nil
		})

		if err != nil {
			fmt.Println("Your Token has been expired.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}



		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rolesString, _ := claims["roles"].(string)
			fmt.Println(rolesString)
			var tokenRoles []models.Role

			if err := json.Unmarshal([]byte(rolesString), &tokenRoles); err != nil {
				fmt.Println("Usercccc.")
			}



		} else{
			fmt.Println("User authorize fail.")
		}
	}


}

func (app *application) getAllMessages(w http.ResponseWriter, r *http.Request) {
	message, err := app.messages.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(message)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Messages have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findMessageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.messages.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Message not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a message")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.messages.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New message have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.messages.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d message(s)", deleteResult.DeletedCount)
}
func (app *application) sendMessage(w http.ResponseWriter, req *http.Request) {
	var m dtos.MessageDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	var message = models.Message{
		Sender:  m.Sender,
		FeedPost: primitive.ObjectID{},
		AlbumPost:  primitive.ObjectID{},
		StoryPost:  primitive.ObjectID{},
		DisposableImage: primitive.ObjectID{},
		DateTime: time.Now(),
		Deleted: false,
		Text: m.Text,
	}
	allChats,_ := app.chats.GetAll()
	usersChat := getChat(app,allChats,m.Sender,m.Receiver)
	var chatUpdate = models.Chat{
		Id : usersChat.Id,
		User1 : usersChat.User1,
		User2: usersChat.User2,
		Messages: append(usersChat.Messages,message),
		UserThatDeletedChat: primitive.ObjectID{},
		Deleted: false,

	}

	insertResult, err := app.chats.Update(chatUpdate)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
//	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationPost/"+"Feed Post"+"/"+userId)
//	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}
func (app *application) sendPostMessage(w http.ResponseWriter, req *http.Request) {
	var m dtos.MessagePostDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	var message = models.Message{
		Sender:  m.Sender,
		FeedPost: m.FeedPost,
		AlbumPost: m.AlbumPost,
		StoryPost:  m.StoryPost,
		DisposableImage: primitive.ObjectID{},
		DateTime: time.Now(),
		Deleted: false,
		Text: "",
	}
	allChats,_ := app.chats.GetAll()
	for _, receiver := range m.Receivers {
		usersChat := getChat(app,allChats,m.Sender,receiver)
		var chatUpdate = models.Chat{
			Id : usersChat.Id,
			User1 : usersChat.User1,
			User2: usersChat.User2,
			Messages: append(usersChat.Messages,message),
			UserThatDeletedChat: primitive.ObjectID{},
			Deleted: false,

		}
		insertResult, err := app.chats.Update(chatUpdate)
		if err != nil {
			app.serverError(w, err)
		}

		app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
		//	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationPost/"+"Feed Post"+"/"+userId)
		//	fmt.Println(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		idMarshaled, err := json.Marshal(insertResult.UpsertedID)
		fmt.Println(insertResult.UpsertedID)
		w.Write(idMarshaled)
	}



}
func getChat(app *application,chats []models.Chat, sender primitive.ObjectID,receiver primitive.ObjectID) models.Chat {
	for _, chat := range chats {
		user1 :=chat.User1
		user2 := chat.User2
		if user1 ==sender && user2 == receiver {
			return chat
		}
		if user1 ==receiver && user2 == sender {
			return chat
		}
	}
	chatBetweenUsers := insertChat(app,sender,receiver)
	return chatBetweenUsers
}
func insertChat(app *application,user1 primitive.ObjectID,user2 primitive.ObjectID) models.Chat {


	var chat = models.Chat{
		User1: user1,
		User2:  user2,
		Messages: []models.Message{},
		Deleted: false,
		UserThatDeletedChat: primitive.ObjectID{},
	}
	insertResult, _ := app.chats.Insert(chat)
	idMarshaled, _ := json.Marshal(insertResult.InsertedID)

	stringId := string(idMarshaled)
	stringId = stringId[1:]
	stringId = stringId[:len(stringId)-1]

	primitiveId,_ :=primitive.ObjectIDFromHex(stringId)
	settingsInserted, _ :=app.chats.FindByID(primitiveId)
	return *settingsInserted
}
func (app *application) getMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	senderId := vars["sender"]
	receiverId := vars["receiver"]
	senderIdPrimitive, _ := primitive.ObjectIDFromHex(senderId)
	receiverIdPrimitive, _ := primitive.ObjectIDFromHex(receiverId)
	allMessages, _ := app.chats.GetAll()
	allImages, _ := app.disposableImages.GetAll()
	messages, err := findChatBetweenUsers(allMessages, senderIdPrimitive,receiverIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	chatResponse := []dtos.MessageFrontDTO{}
	for _, mes := range messages.Messages {

					chatResponse = append(chatResponse, toResponseChat(mes, allImages,messages.Deleted))
				}


	imagesMarshaled, err := json.Marshal(chatResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func (app *application) deleteChatBetweenUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	senderId := vars["sender"]
	receiverId := vars["receiver"]
	senderIdPrimitive, _ := primitive.ObjectIDFromHex(senderId)
	receiverIdPrimitive, _ := primitive.ObjectIDFromHex(receiverId)
	allMessages, _ := app.chats.GetAll()
	usersChat, err := findChatBetweenUsers(allMessages, senderIdPrimitive,receiverIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}


	var chatUpdate = models.Chat{
		Id : usersChat.Id,
		User1 : usersChat.User1,
		User2: usersChat.User2,
		Messages: usersChat.Messages,
		Deleted: true,
		UserThatDeletedChat: senderIdPrimitive,

	}

	insertResult, err := app.chats.Update(chatUpdate)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	//	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationPost/"+"Feed Post"+"/"+userId)
	//	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}
func (app *application) isChatDeleted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	senderId := vars["sender"]
	receiverId := vars["receiver"]
	senderIdPrimitive, _ := primitive.ObjectIDFromHex(senderId)
	receiverIdPrimitive, _ := primitive.ObjectIDFromHex(receiverId)
	allMessages, _ := app.chats.GetAll()
	usersChat, err := findChatBetweenUsers(allMessages, senderIdPrimitive,receiverIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	chatResponse := dtos.DeletedChatDTO{
		Deleted: usersChat.Deleted,
		ForUser: usersChat.UserThatDeletedChat,
	}



	imagesMarshaled, err := json.Marshal(chatResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)


}
func findChatBetweenUsers(messages []models.Chat, user1 primitive.ObjectID, user2 primitive.ObjectID) (models.Chat, error) {
	chat := models.Chat{}

	for _, mes := range messages {
		if	mes.User1 ==user1 && mes.User2 == user2 {
			chat = mes
		}
		if	mes.User1 ==user2 && mes.User2 == user1 {
			chat = mes
		}
	}
	return chat, nil
}
func findDisposableById(images []models.DisposableImage, id primitive.ObjectID) (models.DisposableImage, error) {
	image := models.DisposableImage{}

	for _, img := range images {
	fmt.Println(img.Id)
	fmt.Println(id)
		if	img.Id == id {
			fmt.Println("PRONASAO")
			image = img
		}
	}
	return image, nil
}
func toResponseChat(message models.Message, images []models.DisposableImage, deleted bool) dtos.MessageFrontDTO {
	disposableImg,_ := findDisposableById(images,message.DisposableImage)
	return dtos.MessageFrontDTO{
		Id: message.Id,
		DateTime : message.DateTime,
		Text: message.Text,
		Sender:  message.Sender.Hex(),
		FeedPost: message.FeedPost,
		StoryPost: message.StoryPost,
		DisposableImage: disposableImg.Media,
		AlbumPost: message.AlbumPost,
		DisposableImageId: disposableImg.Id,
		OpenedDisposable: disposableImg.Opened,

	}
}