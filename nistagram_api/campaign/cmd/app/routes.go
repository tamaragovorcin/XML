package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

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
	r.HandleFunc("/api/getUsersCampaigns/{userId}", app.getUsersCampaigns).Methods("GET")
	r.HandleFunc("/api/file/{campaignId}", app.GetFileByCampaignId).Methods("GET")
	r.HandleFunc("/api/campaign/update", app.updateOneTimeCampaign).Methods("POST")
	r.HandleFunc("/api/campaign/multiple/update", app.updateMultipleTimeCampaign).Methods("POST")
	r.HandleFunc("/api/campaign/delete/{id}", app.deleteOneTimeCampaign).Methods("GET")
	r.HandleFunc("/api/campaign/id/{id}", app.findByIDOneTimeCampaign).Methods("GET")
	r.HandleFunc("/api/campaign/multiple/id/{id}", app.findByIDMultipleTimeCampaign).Methods("GET")
	r.HandleFunc("/api/campaign/delete/multiple/{id}", app.deleteMultipleTimeCampaign).Methods("GET")



	r.HandleFunc("/partnershipRequests/{userId}", app.getPartnershipRequestsOneTime).Methods("GET")
	r.HandleFunc("/acceptPartnership", app.acceptPartnershipRequestOneTime).Methods("POST")
	r.HandleFunc("/deletePartnership", app.deletePartnershipRequestOneTime).Methods("POST")

	r.HandleFunc("/partnershipRequestsMultiple/{userId}", app.getPartnershipRequestsMultiple).Methods("GET")
	r.HandleFunc("/acceptPartnershipMultiple", app.acceptPartnershipRequestMultiple).Methods("POST")
	r.HandleFunc("/deletePartnershipMultiple", app.deletePartnershipRequestMultiple).Methods("POST")

	r.HandleFunc("/promoteMultiple/{userId}", app.getInfluecnersMultipleCampaigns).Methods("GET")
	r.HandleFunc("/promoteOneTime/{userId}", app.getInfluencersOneTimeCampaigns).Methods("GET")

	r.HandleFunc("/multipleHomePage/{userId}/{type}", app.getMultipleHomePage).Methods("GET")
	r.HandleFunc("/oneTimeHomePage/{userId}/{type}", app.getOneTimeHomePage).Methods("GET")

	r.HandleFunc("/multipleHomePage/promote/{userId}/{type}", app.getMultipleHomePagePromote).Methods("GET")
	r.HandleFunc("/oneTimeHomePage/promote/{userId}/{type}", app.getOneTimeHomePagePromote).Methods("GET")

	r.HandleFunc("/oneTimeCampaign/like/", app.likeOneTimeCampaign).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/dislike/", app.dislikeOneTimeCampaign).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/comment/", app.commentOneTimeCampaign).Methods("POST")

	r.HandleFunc("/multipleTimeCampaign/like/", app.likeMultipleCampaign).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/dislike/", app.dislikeMultipleCampaign).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/comment/", app.commentMultipleCampaign).Methods("POST")

	r.HandleFunc("/oneTimeCampaign/likes/{campaignId}", app.getLikesOneTimeCampaign).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/dislikes/{campaignId}", app.getDislikesOneTimeCampaign).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/comments/{campaignId}", app.getCommentsOneTimeCampaign).Methods("GET")

	r.HandleFunc("/multipleTimeCampaign/likes/{campaignId}", app.getLikesMultipleCampaign).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/dislikes/{campaignId}", app.getDislikesMultipleCampaign).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/comments/{campaignId}", app.getCommentsMultipleCampaign).Methods("GET")


	r.HandleFunc("/multipleTimeCampaign/clickLink/", app.clickLinkMultipleCampaign).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/clickLink/", app.clickLinkOneTimeCampaign).Methods("POST")

	r.HandleFunc("/bestPromoters/", app.getBestInfluencers).Methods("GET")
	r.HandleFunc("/storyCampaigns/{userId}", app.getStoryCampaignsForHomePage).Methods("GET")

	return r
}
