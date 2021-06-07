package main

import (
	"feedPosts/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"io"

	"net/http"
	"os"
)


func (app *application) saveVideo (w http.ResponseWriter, r *http.Request) {
	fmt.Println("pogodiooo")
	vars := mux.Vars(r)
	userId := vars["userId"]
	feedId := vars["feedId"]
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	defer file.Close()
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	postIdPrimitive, _ :=primitive.ObjectIDFromHex(feedId)
	var video =models.Video {
		Media : path,
		UserId : userIdPrimitive,
		PostId : postIdPrimitive,
	}

	insertResult, err  := app.videos.Insert(video)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New video has been created, id=%s", insertResult.InsertedID)
}


