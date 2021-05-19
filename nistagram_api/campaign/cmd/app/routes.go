package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/ad/", app.getAllAd).Methods("GET")
	r.HandleFunc("/api/ad/{id}", app.findByIDAd).Methods("GET")
	r.HandleFunc("/api/ad/", app.insertAd).Methods("POST")
	r.HandleFunc("/api/ad/{id}", app.deleteAd).Methods("DELETE")

	r.HandleFunc("/api/campaign/", app.getAllCampaign).Methods("GET")
	r.HandleFunc("/api/campaign/{id}", app.findByIDCampaign).Methods("GET")
	r.HandleFunc("/api/campaign/", app.insertCampaign).Methods("POST")
	r.HandleFunc("/api/campaign/{id}", app.deleteCampaign).Methods("DELETE")

	r.HandleFunc("/api/campaignPost/", app.getAllCampaignPost).Methods("GET")
	r.HandleFunc("/api/campaignPost/{id}", app.findByIDCampaignPost).Methods("GET")
	r.HandleFunc("/api/campaignPost/", app.insertCampaignPost).Methods("POST")
	r.HandleFunc("/api/campaignPost/{id}", app.deleteCampaignPost).Methods("DELETE")

	r.HandleFunc("/api/campaignStory/", app.getAllCampaignStory).Methods("GET")
	r.HandleFunc("/api/campaignStory/{id}", app.findByIDCampaignStory).Methods("GET")
	r.HandleFunc("/api/campaignStory/", app.insertCampaignStory).Methods("POST")
	r.HandleFunc("/api/campaignStory/{id}", app.deleteCampaignStory).Methods("DELETE")

	r.HandleFunc("/api/multipleTimeCampaign/", app.getAllMultipleTimeCampaign).Methods("GET")
	r.HandleFunc("/api/multipleTimeCampaign/{id}", app.findByIDMultipleTimeCampaign).Methods("GET")
	r.HandleFunc("/api/multipleTimeCampaign/", app.insertMultipleTimeCampaign).Methods("POST")
	r.HandleFunc("/api/multipleTimeCampaign/{id}", app.deleteMultipleTimeCampaign).Methods("DELETE")

	r.HandleFunc("/api/oneTimeCampaign/", app.getAllOneTimeCampaign).Methods("GET")
	r.HandleFunc("/api/oneTimeCampaign/{id}", app.findByIDOneTimeCampaign).Methods("GET")
	r.HandleFunc("/api/oneTimeCampaign/", app.insertOneTimeCampaign).Methods("POST")
	r.HandleFunc("/api/oneTimeCampaign/{id}", app.deleteOneTimeCampaign).Methods("DELETE")

	r.HandleFunc("/api/statistic/", app.getAllStatistic).Methods("GET")
	r.HandleFunc("/api/statistic/{id}", app.findByIDStatistic).Methods("GET")
	r.HandleFunc("/api/statistic/", app.insertStatistic).Methods("POST")
	r.HandleFunc("/api/statistic/{id}", app.deleteStatistic).Methods("DELETE")

	return r
}
