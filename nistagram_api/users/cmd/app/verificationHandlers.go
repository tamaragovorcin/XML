package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"users/pkg/models"
)

func (app *application) getAllVerifications(w http.ResponseWriter, r *http.Request) {
	chats, err := app.verifications.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Verifications have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findVerificationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.verifications.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Verification not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a verification")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertVerification(w http.ResponseWriter, r *http.Request) {
	var m models.Verification
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.verifications.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New verification have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteVerification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.verifications.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d verifications(s)", deleteResult.DeletedCount)
}
