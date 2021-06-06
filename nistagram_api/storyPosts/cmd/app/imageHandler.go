package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"storyPosts/pkg/models"
	"strings"
)

func (app *application) getAllImages(w http.ResponseWriter, r *http.Request) {
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

func (app *application) findImageByID(w http.ResponseWriter, r *http.Request) {
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
func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {
		vars := mux.Vars(r)
		userId := vars["userId"]
		feedId := vars["storyId"]
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
		var path = "images/"+hander.Filename
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer f.Close()
		io.Copy(f, file)

		userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
		postIdPrimitive, _ :=primitive.ObjectIDFromHex(feedId)
		var image =models.Image {
			Media : path,
			UserId : userIdPrimitive,
			PostId : postIdPrimitive,
		}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}

func imgPath(carID int) string {
	return fmt.Sprintf("../../images/story/%v/", carID)
}

func imagePath(carID int) (string, error) {
	carPath := imgPath(carID)
	err := os.Mkdir(carPath, 0755)
	if err != nil {
		return "", err
	}
	return carPath, nil
}


func (app *application) insertImage(w http.ResponseWriter, r *http.Request) {
	// Define booking model
	var m models.Image
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new booking
	insertResult, err := app.images.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteImage(w http.ResponseWriter, r *http.Request) {
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
