package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/user/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/", app.insertUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", app.deleteUser).Methods("DELETE")

	r.HandleFunc("/api/registredUser/", app.getAllRegistredUsers).Methods("GET")
	r.HandleFunc("/api/registredUser/{id}", app.findRegistredUserByID).Methods("GET")
	r.HandleFunc("/api/registredUser/", app.insertRegistredUser).Methods("POST")
	r.HandleFunc("/api/registredUser/{id}", app.deleteRegistredUser).Methods("DELETE")

	r.HandleFunc("/api/verification/", app.getAllVerifications).Methods("GET")
	r.HandleFunc("/api/verification/{id}", app.findVerificationByID).Methods("GET")
	r.HandleFunc("/api/verification/", app.insertVerification).Methods("POST")
	r.HandleFunc("/api/verification/{id}", app.deleteVerification).Methods("DELETE")

	r.HandleFunc("/api/report/", app.getAllReports).Methods("GET")
	r.HandleFunc("/api/report/{id}", app.findReportByID).Methods("GET")
	r.HandleFunc("/api/report/", app.insertReports).Methods("POST")
	r.HandleFunc("/api/report/{id}", app.deleteReport).Methods("DELETE")


	r.HandleFunc("/api/role/", app.getAllRoles).Methods("GET")
	r.HandleFunc("/api/role/{id}", app.findRoleByID).Methods("GET")
	r.HandleFunc("/api/role/", app.insertRole).Methods("POST")
	r.HandleFunc("/api/role/{id}", app.deleteRole).Methods("DELETE")

	r.HandleFunc("/api/agent/", app.getAllAgents).Methods("GET")
	r.HandleFunc("/api/agent/{id}", app.findAgentByID).Methods("GET")
	r.HandleFunc("/api/agent/", app.insertAgent).Methods("POST")
	r.HandleFunc("/api/agent/{id}", app.deleteAgent).Methods("DELETE")

	return r
}
