package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"users/pkg/dtos"
	"users/pkg/models"
)

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	chats, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Users have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.users.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func HashAndSaltPasswordIfStrong(password string) (string, error) {


	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	app.infoLog.Println("Users ssssssssssssssss")
	var m dtos.UserRequest
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)
	var profileInformation = models.ProfileInformation{Name: m.Name, LastName: m.LastName,
		Email:       m.Email,
		Username:    m.Username,
		Password:    hashAndSalt,
		Roles:       []models.Role{{ Id: 1, Name: "USER"}},
		PhoneNumber: m.PhoneNumber,
		Gender: models.Gender(m.Gender),
		DateOfBirth: m.DateOfBirth,
	}


	var user = models.User{Id: 1,
		ProfileInformation: profileInformation,
		Biography: m.Biography,
		Private: m.Private,
		Verified: false,
	}


	insertResult, err := app.users.Insert(user)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d users(s)", deleteResult.DeletedCount)
}
