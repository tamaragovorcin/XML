package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"users/pkg/dtos"
	"users/pkg/models"
)

func (app *application) getAllAgents(w http.ResponseWriter, r *http.Request) {
	chats, err := app.agents.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Agents have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findAgentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.agents.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Agent not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a Agent")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var m dtos.AgentRequest
	var able = true
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}


	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	for i, s := range users {
		fmt.Println(i, s)
		if(s.ProfileInformation.Username == m.Username){
			app.infoLog.Printf("This username is already taken")

			able = false


			var e = errors.New("This username is already taken")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			app.serverError(w, errors.New("This username is already taken"))
			break

		}

		if( s.ProfileInformation.Email == m.Email){
			app.infoLog.Printf("User with this email already exists")
			able = false

			var e = errors.New("User with this email already exists")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusConflict)
			w.Write(b)
			app.serverError(w, errors.New("User with this email already exists"))
			break

		}

	}

	if able  {
		hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)

		var profileInformation = models.ProfileInformation{
			Name: m.Name, LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			Password:    hashAndSalt,
			Roles:       []models.Role{{Name: "USER"}},
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}

		var agent = models.Agent{
			ProfileInformation: profileInformation,
			Biography:          m.Biography,
			Private:            m.Private,
			Verified:           false,
			Website: m.Website,
			ApprovedAgent : false,
		}

		insertResult, err := app.agents.Insert(agent)
		if err != nil {
			app.serverError(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idMarshaled, err := json.Marshal(insertResult.InsertedID)

		w.Write(idMarshaled)
	}
}

func (app *application) deleteAgent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.agents.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d Agent(s)", deleteResult.DeletedCount)
}
