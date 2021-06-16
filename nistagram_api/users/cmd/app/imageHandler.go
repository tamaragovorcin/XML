package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"strings"
	"users/pkg/models"
)

func (app *application) getAllImagesVerification(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.images.All()
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

func (app *application) findImageVerificationByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.images.FindByID(id)
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
func (app *application) saveImageVerification(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("TU SAM")

	vars := mux.Vars(r)
	userId := vars["userIdd"]
	verificationId := vars["verificationId"]
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
	//var path = "/var/lib/feedposts/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	postIdPrimitive, _ :=primitive.ObjectIDFromHex(verificationId)
	var image =models.Image {
		Media : path,
		UserId : userIdPrimitive,
		VerificationId : postIdPrimitive,
	}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}


func findImagesByUserId(images []models.Image, idPrimitive primitive.ObjectID) ([]models.Image, error) {
	imagesUser := []models.Image{}

	for _, image := range images {
		if	image.UserId==idPrimitive {
			imagesUser = append(imagesUser, image)
		}
	}
	return imagesUser, nil
}
func findImageByPostId(images []models.Image, idFeedPost primitive.ObjectID) (models.Image, error) {
	imageFeedPost := models.Image{}

	for _, image := range images {
		if	image.VerificationId==idFeedPost {
			imageFeedPost = image
		}
	}
	return imageFeedPost, nil
}

func (app *application) deleteImageVerification(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.images.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d image(s)", deleteResult.DeletedCount)
}

