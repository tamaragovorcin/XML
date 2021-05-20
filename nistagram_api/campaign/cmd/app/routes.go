package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/ad/", app.getAllAd).Methods("GET")
	r.HandleFunc("/ad/{id}", app.findByIDAd).Methods("GET")
	r.HandleFunc("/ad/", app.insertAd).Methods("POST")
	r.HandleFunc("/ad/{id}", app.deleteAd).Methods("DELETE")

	r.HandleFunc("/", app.getAllCampaign).Methods("GET")
	r.HandleFunc("/{id}", app.findByIDCampaign).Methods("GET")
	r.HandleFunc("/", app.insertCampaign).Methods("POST")
	r.HandleFunc("/{id}", app.deleteCampaign).Methods("DELETE")

	r.HandleFunc("/campaignPost/", app.getAllCampaignPost).Methods("GET")
	r.HandleFunc("/campaignPost/{id}", app.findByIDCampaignPost).Methods("GET")
	r.HandleFunc("/campaignPost/", app.insertCampaignPost).Methods("POST")
	r.HandleFunc("/campaignPost/{id}", app.deleteCampaignPost).Methods("DELETE")

	r.HandleFunc("/campaignStory/", app.getAllCampaignStory).Methods("GET")
	r.HandleFunc("/campaignStory/{id}", app.findByIDCampaignStory).Methods("GET")
	r.HandleFunc("/campaignStory/", app.insertCampaignStory).Methods("POST")
	r.HandleFunc("/campaignStory/{id}", app.deleteCampaignStory).Methods("DELETE")

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

	return r
}
