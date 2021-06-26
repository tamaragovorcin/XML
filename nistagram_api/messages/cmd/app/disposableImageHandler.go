package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gomod/pkg/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (app *application) getAllDisposableImages(w http.ResponseWriter, r *http.Request) {
	disposableImage, err := app.disposableImages.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(disposableImage)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("DisposableImages have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}

	defer file.Close()
	//var path = "/var/lib/feedposts/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	var image =models.DisposableImage {
		Media : path,
		Opened: false,
	}

	insertResult, err  := app.disposableImages.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}
func (app *application) sendDisposableImage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	sender := vars["sender"]
	receiver := vars["receiver"]
	senderIdPrimitive, _ := primitive.ObjectIDFromHex(sender)
	receiverIdPrimitive, _ :=primitive.ObjectIDFromHex(receiver)
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}

	defer file.Close()
	//var path = "/var/lib/feedposts/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	var image =models.DisposableImage {
		Media : path,
		Opened: false,
	}
	insertResult1, _ := app.disposableImages.Insert(image)
	idIm := insertResult1.InsertedID.(primitive.ObjectID)
	var message = models.Message{
		Sender:  senderIdPrimitive,
		FeedPost: primitive.ObjectID{},
		AlbumPost:  primitive.ObjectID{},
		StoryPost:  primitive.ObjectID{},
		DisposableImage: idIm,
		DateTime: time.Now(),
		Deleted: false,
		Text: "",
	}
	allChats,_ := app.chats.GetAll()
	usersChat := getChat(app,allChats,senderIdPrimitive,receiverIdPrimitive)
	var chatUpdate = models.Chat{
		Id : usersChat.Id,
		User1 : usersChat.User1,
		User2: usersChat.User2,
		Messages: append(usersChat.Messages,message),

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

func (app *application) findDisposableImageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.disposableImages.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("DisposableImage not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a disposableImage")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func(app *application) openDisposableImage(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	IdPrimitive, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("ID ZA UPDATE")
	fmt.Println(id)
	allImages, _:= app.disposableImages.GetAll()
	image,_:= findDisposableById(allImages,IdPrimitive)
	var disposableUpdate = models.DisposableImage{
		Id : IdPrimitive,
		Opened: true,
		Media: image.Media,
	}

	insertResult, err := app.disposableImages.Update(disposableUpdate)
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
func(app *application) getDisposableImage(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	path := vars["path"]


	file, err:=os.Open("files/"+path)
	if err!=nil{
		http.Error(w,"file not found",404)
		return
	}


	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)
	FileStat,_:= file.Stat()
	FileSize:= strconv.FormatInt(FileStat.Size(),10)
	w.Header().Set("Content-Disposition", "attachment; filename="+path)
	w.Header().Set("Content-Type", ContentType)
	w.Header().Set("Content-Length", FileSize)

	file.Seek(0,0)
	io.Copy(w,file)
	return



}
func (app *application) insertDisposableImage(w http.ResponseWriter, r *http.Request) {
	var m models.DisposableImage
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.disposableImages.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New disposableImage have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteDisposableImage(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.disposableImages.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d disposableImage(s)", deleteResult.DeletedCount)
}
