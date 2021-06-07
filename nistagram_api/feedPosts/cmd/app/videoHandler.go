package main

import (
	"bytes"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"strconv"
	"strings"

	"io"

	"net/http"
	"os"
)


func (app *application) uploadFile (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]
	//feedId := vars["feedId"]
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("inputFile")
	if err != nil {
		fmt.Println(err.Error())

	}
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	defer file.Close()
	var path = "images/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	//postIdPrimitive, _ :=primitive.ObjectIDFromHex(feedId)
	var video =models.Video {
		Media : path,
		UserId : userIdPrimitive,
		//PostId : postIdPrimitive,
	}

	insertResult, err  := app.videos.Insert(video)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}

func(app *application) GetVideo(w http.ResponseWriter, r *http.Request){
	fmt.Println("--------------------------------------")
	file, err:=os.Open("images/IMG-0d12b381f8c253794f87cf8b907e3e21-V.jpg")
	if err!=nil{
		http.Error(w,"file not found",404)
		return
	}


	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)
	FileStat,_:= file.Stat()
	FileSize:= strconv.FormatInt(FileStat.Size(),10)
	w.Header().Set("Content-Disposition", "attachment; filename="+"images/IMG-0d12b381f8c253794f87cf8b907e3e21-V.jpg")
	w.Header().Set("Content-Type", ContentType)
	w.Header().Set("Content-Length", FileSize)

	file.Seek(0,0)
	io.Copy(w,file)
	return
	/*feedPostResponse := dtos.VideoDTO{}
	feedPostResponse = videoToResponse("videos/20190303_032235.mp4")
	imagesMarshaled, err := json.Marshal(feedPostResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)*/
}

func videoToResponse(image2 string) dtos.VideoDTO {
	f, _ := os.Open(image2)


	defer f.Close()
	image, _, _ := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}
	return dtos.VideoDTO{
		Media: buffer.Bytes(),

	}
}