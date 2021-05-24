package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/", app.insertUser).Methods("POST")
	r.HandleFunc("/{id}", app.deleteUser).Methods("DELETE")

	r.HandleFunc("/profileInformation/", app.getAllProfileInformation).Methods("GET")
	r.HandleFunc("/profileInformation/{id}", app.findProfileInformationByID).Methods("GET")
	r.HandleFunc("/profileInformation/", app.insertProfileInformation).Methods("POST")
	r.HandleFunc("/profileInformation/{id}", app.deleteProfileInformation).Methods("DELETE")

	r.HandleFunc("/verification/", app.getAllVerifications).Methods("GET")
	r.HandleFunc("/verification/{id}", app.findVerificationByID).Methods("GET")
	r.HandleFunc("/verification/", app.insertVerification).Methods("POST")
	r.HandleFunc("/verification/{id}", app.deleteVerification).Methods("DELETE")

	r.HandleFunc("/role/", app.getAllRoles).Methods("GET")
	r.HandleFunc("/role/{id}", app.findRoleByID).Methods("GET")
	r.HandleFunc("/role/", app.insertRole).Methods("POST")
	r.HandleFunc("/role/{id}", app.deleteRole).Methods("DELETE")

	r.HandleFunc("/agent/", app.getAllAgents).Methods("GET")
	r.HandleFunc("/agent/{id}", app.findAgentByID).Methods("GET")
	r.HandleFunc("/agent/", app.insertAgent).Methods("POST")
	r.HandleFunc("/agent/{id}", app.deleteAgent).Methods("DELETE")

	return r
}
