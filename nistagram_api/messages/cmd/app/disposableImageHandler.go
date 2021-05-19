package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gomod/pkg/models"
	"net/http"
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
