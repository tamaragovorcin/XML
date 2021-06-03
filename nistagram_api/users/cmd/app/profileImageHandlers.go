package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"users/pkg/dtos"
	"users/pkg/models"
)

func (app *application) getAllImages(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.profileImage.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Images have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findImageByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.profileImage.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Image not found")
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

	app.infoLog.Println("Have been found a image")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) getUsersProfileImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages,_ := app.profileImage.All()
	usersFeedPosts,err :=findImageByUserId(allImages,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.ProfileImageInfoDTO{}
	feedPostResponse = append(feedPostResponse, toResponse(usersFeedPosts.Media))


	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func toResponse(image2 string) dtos.ProfileImageInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image, _, _ := image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}

	return dtos.ProfileImageInfoDTO{
		Media: buffer.Bytes(),
	}
}
func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	userId := vars["userId"]
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
	var path = "profileImages/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	var image =models.ProfileImage {
		Media : path,
		UserId : userIdPrimitive,
	}

	insertResult, err  := app.profileImage.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}


func findImageByUserId(images []models.ProfileImage, idPrimitive primitive.ObjectID) (models.ProfileImage, error) {
	imagesUser := models.ProfileImage{}

	for _, image := range images {
		if	image.UserId==idPrimitive {
			imagesUser = image
		}
	}
	return imagesUser, nil
}

func (app *application) deleteImage(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.profileImage.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d image(s)", deleteResult.DeletedCount)
}

