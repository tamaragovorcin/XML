package main

import (
	"campaigns/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"strings"
)

func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	userId := vars["userIdd"]
	campaignId := vars["campaignId"]
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
	//var path = "/var/lib/campaigns/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	campaignIdPrimitive, _ :=primitive.ObjectIDFromHex(campaignId)
	var image =models.Image {
		Media : path,
		UserId : userIdPrimitive,
		CampaignId : campaignIdPrimitive,
	}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}