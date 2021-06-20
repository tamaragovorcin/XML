package main

import (
	"campaigns/pkg/dtos"
	"campaigns/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
)

func (app *application) getAllOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.oneTimeCampaign.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Movies have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByIDOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.oneTimeCampaign.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Movie not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert movie to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a movie")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) updateOneTimeCampaign(w http.ResponseWriter, req *http.Request) {
	var dto dtos.OneTimeCampaignUpdateDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Println("User:")
	fmt.Println(dto.User)
	userPrimitive, _ := primitive.ObjectIDFromHex(dto.User)

	var campaign = models.Campaign{
		User : userPrimitive,
		Link : dto.Link,
		Description :dto.Description,
	}
	IdPrimitive, _ := primitive.ObjectIDFromHex(dto.Id)

	var oneTimeCampaign = models.OneTimeCampaign{
		Id : IdPrimitive,
		Campaign:   campaign,
		Time: dto.Time,
		Date : dto.Date,

	}

	insertResult, err := app.oneTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}

func (app *application) insertOneTimeCampaign(w http.ResponseWriter, req *http.Request) {
	var dto dtos.OneTimeCampaignDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(dto.User)


	var campaign = models.Campaign{
		User : userIdPrimitive,
		TargetGroup : dto.TargetGroup,
		Statistic  :[]models.Statistic{},
		Link : dto.Link,
		Description :dto.Description,
		Partnerships :getPartnerships(dto.PartnershipsRequests),
	}
	var oneTimeCampaign = models.OneTimeCampaign{
		Campaign:   campaign,
		Time: dto.Time,
		Date : dto.Date,

	}

	insertResult, err := app.oneTimeCampaign.Insert(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func getPartnerships(requests []string) []models.Partnership {
	partnerships := []models.Partnership{}
	for _, request := range requests {
		primitiveRequest, _ := primitive.ObjectIDFromHex(request)
		var partnership = models.Partnership{
			Influencer : primitiveRequest,
			Approved: false,
		}
		partnerships = append(partnerships, partnership)
	}
	return partnerships
}

func (app *application) deleteOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.oneTimeCampaign.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}

func (app *application) getPartnershipRequestsOneTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	usersCampaigns,err :=findPartnershipRequestsByUserIdOneTime(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range usersCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignToResponseInfluencer(campaign,contentType))

	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}


func (app *application) getInfluencersOneTimeCampaigns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	usersCampaigns,err :=findPartnershipByUserIdOneTime(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range usersCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignToResponseInfluencer(campaign,contentType))

	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func campaignToResponseInfluencer(campaing models.OneTimeCampaign, contentType string) dtos.CampaignDTO {
	username :=getUserUsername(campaing.Campaign.User)
	return dtos.CampaignDTO{
		Id: campaing.Id.Hex(),
		User: campaing.Campaign.User.Hex(),
		Description: campaing.Campaign.Description,
		Time: campaing.Time,
		Date: campaing.Date,
		Link: campaing.Campaign.Link,
		ContentType: contentType,
		AgentUsername: username,
	}
}
func getUserUsername(user primitive.ObjectID) string {

	stringObjectID := user.Hex()
	resp, err := http.Get("http://localhost:80/api/users/api/user/username/"+stringObjectID)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	sb = sb[1:]
	sb = sb[:len(sb)-1]
	return sb
}
func findPartnershipRequestsByUserIdOneTime(posts []models.OneTimeCampaign, idPrimitive primitive.ObjectID) ([]models.OneTimeCampaign, error) {
	campaignsUser := []models.OneTimeCampaign{}

	for _, campaign := range posts {
		if	userInPartnershipRequests(campaign.Campaign.Partnerships,idPrimitive) {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}
func findPartnershipByUserIdOneTime(posts []models.OneTimeCampaign, idPrimitive primitive.ObjectID) ([]models.OneTimeCampaign, error) {
	campaignsUser := []models.OneTimeCampaign{}

	for _, campaign := range posts {
		if	userInPartnership(campaign.Campaign.Partnerships,idPrimitive) {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}

func (app *application) acceptPartnershipRequestOneTime(w http.ResponseWriter, req *http.Request) {
	var dto dtos.PartnershipDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	campaign, err := app.oneTimeCampaign.FindByID(dto.CampaignId.Hex())
	partensrhipsUpdated := handleUpdatedPartnerships(campaign.Campaign.Partnerships,dto.UserId)
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: partensrhipsUpdated,
	}
	var oneTimeCampaign = models.OneTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		Time: campaign.Time,
		Date : campaign.Date,
	}


	insertResult, err := app.oneTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}

func handleUpdatedPartnerships(partnerships []models.Partnership, id primitive.ObjectID) []models.Partnership {
	updated :=[]models.Partnership{}
	for _, partnership := range partnerships {
		if partnership.Influencer.Hex()==id.Hex() {
			var partnershipOne = models.Partnership{
				ID : partnership.ID,
				Influencer: id,
				Approved: true,
			}
			updated = append(updated,partnershipOne)

		} else {
			updated = append(updated,partnership)
		}
	}
	return updated
}

func (app *application) deletePartnershipRequestOneTime(w http.ResponseWriter, req *http.Request) {
	var dto dtos.PartnershipDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	campaign, err := app.oneTimeCampaign.FindByID(dto.CampaignId.Hex())
	partensrhipsUpdated := handleDeletePartnerships(campaign.Campaign.Partnerships,dto.UserId)
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: partensrhipsUpdated,
	}


	var oneTimeCampaign = models.OneTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		Time: campaign.Time,
		Date : campaign.Date,
	}

	insertResult, err := app.oneTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}
func handleDeletePartnerships(partnerships []models.Partnership, id primitive.ObjectID) []models.Partnership {
	updated :=[]models.Partnership{}
	for _, partnership := range partnerships {

		if partnership.Influencer.Hex()!=id.Hex() {

			updated = append(updated,partnership)
		}
	}
	return updated
}

func (app *application) getOneTimeHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	campaigns,err :=findNotMyOneTimeCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range campaigns {

		if isGenderOk(userId,campaign.Campaign.TargetGroup.Gender) {
			if isDateOfBirthOk(userId,campaign.Campaign.TargetGroup.DateOne,campaign.Campaign.TargetGroup.DateTwo) {
				if isLocationOk(userId,campaign.Campaign.TargetGroup.Location) {
					contentType := app.GetFileTypeByPostId(campaign.Id)
					campaignResponse = append(campaignResponse, campaignToResponseInfluencer(campaign,contentType))
				}
			}
		}


	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findNotMyOneTimeCampaigns(oneTimeCampaigns []models.OneTimeCampaign, idPrimitive primitive.ObjectID) ([]models.OneTimeCampaign, error) {
	campaigns := []models.OneTimeCampaign{}

	for _, oneCampaign := range oneTimeCampaigns {

		if	oneCampaign.Campaign.User.Hex()!=idPrimitive.Hex() {
			campaigns = append(campaigns, oneCampaign)
		}
	}
	return campaigns, nil
}