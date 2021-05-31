package main

import (
	"encoding/json"
	"feedPosts/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
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
		app.infoLog.Printf("3131111111111111111111131313")

		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Printf("method is post")
		r.ParseMultipartForm(32 << 20)
		file, hander, err := r.FormFile("file")
		if err != nil {
			fmt.Println("**************************************")
			fmt.Println(err.Error())

		}
		defer file.Close()
		var path = "./images/feed/"+id+ "_"+hander.Filename
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
		if err != nil {
			fmt.Println("--------------------------------------")

			fmt.Println(err.Error())
		}
		defer f.Close()
		io.Copy(f, file)

		var image =models.Image {
			Media : path,
			UserId : 10,
			PostId : 453,
		}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}

func imgPath(carID int) string {
	return fmt.Sprintf("../../images/feed/%v/", carID)
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