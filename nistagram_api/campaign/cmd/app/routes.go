package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/multipleTimeCampaign/", IsAuthorized(app.getAllMultipleTimeCampaign)).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/{id}", IsAuthorized(app.findByIDMultipleTimeCampaign)).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/", IsAuthorized(app.insertMultipleTimeCampaign)).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/{id}", IsAuthorized(app.deleteMultipleTimeCampaign)).Methods("DELETE")
	r.HandleFunc("/multipleTimeCampaign/{token}", IsAuthorized(app.insertMultipleTimeCampaignWithToken)).Methods("POST")

	r.HandleFunc("/oneTimeCampaign/", IsAuthorized(app.getAllOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/{id}", IsAuthorized(app.findByIDOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/", IsAuthorized(app.insertOneTimeCampaign)).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/{id}", IsAuthorized(app.deleteOneTimeCampaign)).Methods("DELETE")
	r.HandleFunc("/oneTimeCampaign/{token}", IsAuthorized(app.insertOneTimeCampaignWithToken)).Methods("POST")

	r.HandleFunc("/statistic/", IsAuthorized(app.getAllStatistic)).Methods("GET")
	r.HandleFunc("/statistic/{id}", IsAuthorized(app.findByIDStatistic)).Methods("GET")
	r.HandleFunc("/statistic/", IsAuthorized(app.insertStatistic)).Methods("POST")
	r.HandleFunc("/statistic/{id}", IsAuthorized(app.deleteStatistic)).Methods("DELETE")

	r.HandleFunc("/partnership/", IsAuthorized(app.getAllPartnerships)).Methods("GET")
	r.HandleFunc("/partnership/{id}", IsAuthorized(app.findPartnershipByID)).Methods("GET")
	r.HandleFunc("/partnership/", IsAuthorized(app.insertPartnership)).Methods("POST")
	r.HandleFunc("/partnership/{id}", IsAuthorized(app.deletePartnership)).Methods("DELETE")

	r.HandleFunc("/api/image/{userIdd}/{campaignId}", IsAuthorized(app.saveImage)).Methods("POST")
	r.HandleFunc("/api/image/{token}/{campaignId}", IsAuthorized(app.saveImageWithToken)).Methods("POST")

	r.HandleFunc("/api/getUsersCampaigns/{userId}", IsAuthorized(app.getUsersCampaigns)).Methods("GET")
	r.HandleFunc("/api/file/{campaignId}", app.GetFileByCampaignId).Methods("GET")
	r.HandleFunc("/api/campaign/update", IsAuthorized(app.updateOneTimeCampaign)).Methods("POST")
	r.HandleFunc("/api/campaign/multiple/update", IsAuthorized(app.updateMultipleTimeCampaign)).Methods("POST")
	r.HandleFunc("/api/campaign/delete/{id}", IsAuthorized(app.deleteOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/api/campaign/id/{id}", IsAuthorized(app.findByIDOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/api/campaign/multiple/id/{id}", IsAuthorized(app.findByIDMultipleTimeCampaign)).Methods("GET")
	r.HandleFunc("/api/campaign/delete/multiple/{id}", IsAuthorized(app.deleteMultipleTimeCampaign)).Methods("GET")



	r.HandleFunc("/partnershipRequests/{userId}", IsAuthorized(app.getPartnershipRequestsOneTime)).Methods("GET")
	r.HandleFunc("/acceptPartnership", IsAuthorized(app.acceptPartnershipRequestOneTime)).Methods("POST")
	r.HandleFunc("/deletePartnership", IsAuthorized(app.deletePartnershipRequestOneTime)).Methods("POST")

	r.HandleFunc("/partnershipRequestsMultiple/{userId}", IsAuthorized(app.getPartnershipRequestsMultiple)).Methods("GET")
	r.HandleFunc("/acceptPartnershipMultiple", IsAuthorized(app.acceptPartnershipRequestMultiple)).Methods("POST")
	r.HandleFunc("/deletePartnershipMultiple", IsAuthorized(app.deletePartnershipRequestMultiple)).Methods("POST")

	r.HandleFunc("/promoteMultiple/{userId}", IsAuthorized(app.getInfluecnersMultipleCampaigns)).Methods("GET")
	r.HandleFunc("/promoteOneTime/{userId}", IsAuthorized(app.getInfluencersOneTimeCampaigns)).Methods("GET")

	r.HandleFunc("/multipleHomePage/{userId}/{type}", IsAuthorized(app.getMultipleHomePage)).Methods("GET")
	r.HandleFunc("/oneTimeHomePage/{userId}/{type}", IsAuthorized(app.getOneTimeHomePage)).Methods("GET")

	r.HandleFunc("/multipleHomePage/promote/{userId}/{type}", IsAuthorized(app.getMultipleHomePagePromote)).Methods("GET")
	r.HandleFunc("/oneTimeHomePage/promote/{userId}/{type}", IsAuthorized(app.getOneTimeHomePagePromote)).Methods("GET")

	r.HandleFunc("/oneTimeCampaign/like/", IsAuthorized(app.likeOneTimeCampaign)).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/dislike/", IsAuthorized(app.dislikeOneTimeCampaign)).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/comment/", IsAuthorized(app.commentOneTimeCampaign)).Methods("POST")

	r.HandleFunc("/multipleTimeCampaign/like/", IsAuthorized(app.likeMultipleCampaign)).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/dislike/", IsAuthorized(app.dislikeMultipleCampaign)).Methods("POST")
	r.HandleFunc("/multipleTimeCampaign/comment/", IsAuthorized(app.commentMultipleCampaign)).Methods("POST")

	r.HandleFunc("/oneTimeCampaign/likes/{campaignId}", IsAuthorized(app.getLikesOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/dislikes/{campaignId}", IsAuthorized(app.getDislikesOneTimeCampaign)).Methods("GET")
	r.HandleFunc("/oneTimeCampaign/comments/{campaignId}", IsAuthorized(app.getCommentsOneTimeCampaign)).Methods("GET")

	r.HandleFunc("/multipleTimeCampaign/likes/{campaignId}", IsAuthorized(app.getLikesMultipleCampaign)).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/dislikes/{campaignId}", IsAuthorized(app.getDislikesMultipleCampaign)).Methods("GET")
	r.HandleFunc("/multipleTimeCampaign/comments/{campaignId}", IsAuthorized(app.getCommentsMultipleCampaign)).Methods("GET")


	r.HandleFunc("/multipleTimeCampaign/clickLink/", IsAuthorized(app.clickLinkMultipleCampaign)).Methods("POST")
	r.HandleFunc("/oneTimeCampaign/clickLink/", IsAuthorized(app.clickLinkOneTimeCampaign)).Methods("POST")

	r.HandleFunc("/bestPromoters/", IsAuthorized(app.getBestInfluencers)).Methods("GET")
	r.HandleFunc("/storyCampaigns/{userId}", IsAuthorized(app.getStoryCampaignsForHomePage)).Methods("GET")

	r.HandleFunc("/bestCampaigns/{token}", IsAuthorized(app.getBestUsersCampaign)).Methods("GET")

	return r
}
