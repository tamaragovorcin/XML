package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/", app.getAllCampaign).Methods("GET")
	r.HandleFunc("/{id}", app.findByIDCampaign).Methods("GET")
	r.HandleFunc("/", app.insertCampaign).Methods("POST")
	r.HandleFunc("/{id}", app.deleteCampaign).Methods("DELETE")

	r.HandleFunc("/multipleTimeCampaign/", app.getAllMultipleTimeCampaign).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/{id}", app.findByIDMultipleTimeCampaign).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/", app.insertMultipleTimeCampaign).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/{id}", app.deleteMultipleTimeCampaign).Methods("DELETE")

	r.HandleFunc("/oneTimeCampaign/", app.getAllOneTimeCampaign).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/{id}", app.findByIDOneTimeCampaign).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/", app.insertOneTimeCampaign).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/{id}", app.deleteOneTimeCampaign).Methods("DELETE")

	r.HandleFunc("/statistic/", app.getAllStatistic).Methods("GET")
	r.HandleFunc("/statistic/{id}", app.findByIDStatistic).Methods("GET")
	r.HandleFunc("/statistic/", app.insertStatistic).Methods("POST")
	r.HandleFunc("/statistic/{id}", app.deleteStatistic).Methods("DELETE")

	r.HandleFunc("/partnership/", app.getAllPartnerships).Methods("GET")
	r.HandleFunc("/partnership/{id}", app.findPartnershipByID).Methods("GET")
	r.HandleFunc("/partnership/", app.insertPartnership).Methods("POST")
	r.HandleFunc("/partnership/{id}", app.deletePartnership).Methods("DELETE")

	r.HandleFunc("/api/image/{userIdd}/{campaignId}", app.saveImage).Methods("POST")

	return r
}
