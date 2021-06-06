package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/api/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/update/", app.updateUser).Methods("POST")
	r.HandleFunc("/api/", app.insertUser).Methods("POST")
	r.HandleFunc("/api/{id}", app.deleteUser).Methods("DELETE")

	r.HandleFunc("/api/search/{name}", app.search).Methods("GET")
	r.HandleFunc("/api/login", app.loginUser).Methods("POST")
	r.HandleFunc("/api/user/privacy/{userId}", app.findUserPrivacy).Methods("GET")
	r.HandleFunc("/api/user/username/{userId}", app.findUserUsername).Methods("GET")

	//r.HandleFunc("/api/getLoggedIn", app.getLoggedIn).Methods("GET")

	r.HandleFunc("/profileInformation/", app.getAllProfileInformation).Methods("GET")
	r.HandleFunc("/profileInformation/{id}", app.findProfileInformationByID).Methods("GET")
	r.HandleFunc("/profileInformation/", app.insertProfileInformation).Methods("POST")
	r.HandleFunc("/profileInformation/{id}", app.deleteProfileInformation).Methods("DELETE")

	r.HandleFunc("/verification/", app.getAllVerifications).Methods("GET")
	r.HandleFunc("/verification/{id}", app.findVerificationByID).Methods("GET")
	r.HandleFunc("/verification/", app.insertVerification).Methods("POST")
	r.HandleFunc("/verification/{id}", app.deleteVerification).Methods("DELETE")

	r.HandleFunc("/api/role/", app.getAllRoles).Methods("GET")
	r.HandleFunc("/api/role/{id}", app.findRoleByID).Methods("GET")
	r.HandleFunc("/api/role/", app.insertRole).Methods("POST")
	r.HandleFunc("/api/role/{id}", app.deleteRole).Methods("DELETE")

	r.HandleFunc("/agent/", app.getAllAgents).Methods("GET")
	r.HandleFunc("/agent/{id}", app.findAgentByID).Methods("GET")
	r.HandleFunc("/agent/", app.insertAgent).Methods("POST")
	r.HandleFunc("/agent/{id}", app.deleteAgent).Methods("DELETE")

	r.HandleFunc("/api/user/profileImage/{userId}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/user/profileImage/{userId}", app.getUsersProfileImage).Methods("GET")

	r.HandleFunc("/notification/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/notification/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/notification/", app.insertNotification).Methods("POST")
	r.HandleFunc("/notification/{id}", app.deleteNotification).Methods("DELETE")

	r.HandleFunc("/settings/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/settings/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/settings/", app.insertNotification).Methods("POST")
	r.HandleFunc("/settings/{id}", app.deleteNotification).Methods("DELETE")

	return r
}