package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
	var m models.Agent
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.agents.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New Agent have been created, id=%s", insertResult.InsertedID)
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
