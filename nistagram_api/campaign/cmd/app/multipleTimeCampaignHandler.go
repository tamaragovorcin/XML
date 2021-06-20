package main

import (
	"campaigns/pkg/dtos"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"

	"campaigns/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) getAllMultipleTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.multipleTimeCampaign.All()
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
func (app *application) updateMultipleTimeCampaign(w http.ResponseWriter, req *http.Request) {
	var dto dtos.MultipleTimeCampaignUpdateDTO

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
	var number,_ = strconv.Atoi(dto.DesiredNumber)
	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id : IdPrimitive,
		Campaign:   campaign,
		StartTime: dto.StartTime,
		EndTime : dto.EndTime,
		DesiredNumber:   number,

	}

	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (app *application) findByIDMultipleTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.multipleTimeCampaign.FindByID(id)
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

func (app *application) insertMultipleTimeCampaign(w http.ResponseWriter, req *http.Request) {
	var dto dtos.MultipleTimeCampaignDTO

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
	number,_ :=strconv.Atoi(dto.DesiredNumber)
	var multipleTimeCampaign = models.MultipleTimeCampaign{
		Campaign:   campaign,
		StartTime: dto.StartTime,
		EndTime : dto.EndTime,
		DesiredNumber: number,
		ModifiedTime: time.Now(),

	}

	insertResult, err := app.multipleTimeCampaign.Insert(multipleTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func (app *application) deleteMultipleTimeCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.multipleTimeCampaign.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}
func (app *application) getPartnershipRequestsMultiple(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.multipleTimeCampaign.All()
	usersCampaigns,err :=findPartnershipRequestsByUserIdMultiple(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignMultipleDTO{}
	for _, campaign := range usersCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultiple(campaign,contentType))

	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func (app *application) getInfluecnersMultipleCampaigns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.multipleTimeCampaign.All()
	usersCampaigns,err :=findPartnershipsByUserIdMultiple(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignMultipleDTO{}
	for _, campaign := range usersCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultiple(campaign,contentType))

	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func campaignToResponseInfluencerMultiple(campaing models.MultipleTimeCampaign, contentType string) dtos.CampaignMultipleDTO {
	username :=getUserUsername(campaing.Campaign.User)
	return dtos.CampaignMultipleDTO{
		Id: campaing.Id.Hex(),
		User: campaing.Campaign.User.Hex(),
		Description: campaing.Campaign.Description,
		StartTime: campaing.StartTime,
		EndTime: campaing.EndTime,
		DesiredNumber: strconv.Itoa(campaing.DesiredNumber),
		Link: campaing.Campaign.Link,
		ContentType: contentType,
		AgentUsername: username,
	}
}
func (app *application) acceptPartnershipRequestMultiple(w http.ResponseWriter, req *http.Request) {
	var dto dtos.PartnershipDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	campaign, err := app.multipleTimeCampaign.FindByID(dto.CampaignId.Hex())
	partensrhipsUpdated := handleUpdatedPartnerships(campaign.Campaign.Partnerships,dto.UserId)
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: partensrhipsUpdated,
	}
	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
	}


	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}
func (app *application) deletePartnershipRequestMultiple(w http.ResponseWriter, req *http.Request) {
	var dto dtos.PartnershipDTO

	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		app.serverError(w, err)
	}
	campaign, err := app.multipleTimeCampaign.FindByID(dto.CampaignId.Hex())
	partensrhipsUpdated := handleDeletePartnerships(campaign.Campaign.Partnerships,dto.UserId)
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: partensrhipsUpdated,
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
	}
	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.UpsertedID)
	w.Write(idMarshaled)
}

func findPartnershipRequestsByUserIdMultiple(posts []models.MultipleTimeCampaign, idPrimitive primitive.ObjectID) ([]models.MultipleTimeCampaign, error) {
	campaignsUser := []models.MultipleTimeCampaign{}

	for _, campaign := range posts {
		if	userInPartnershipRequests(campaign.Campaign.Partnerships,idPrimitive) {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}

func findPartnershipsByUserIdMultiple(posts []models.MultipleTimeCampaign, idPrimitive primitive.ObjectID) ([]models.MultipleTimeCampaign, error) {
	campaignsUser := []models.MultipleTimeCampaign{}

	for _, campaign := range posts {
		if	userInPartnership(campaign.Campaign.Partnerships,idPrimitive) {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}


func (app *application) getMultipleHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.multipleTimeCampaign.All()
	campaigns,err :=findNotMyMultipleCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignMultipleDTO{}
	for _, campaign := range campaigns {
		if isGenderOk(userId,campaign.Campaign.TargetGroup.Gender) {
			if isDateOfBirthOk(userId,campaign.Campaign.TargetGroup.DateOne,campaign.Campaign.TargetGroup.DateTwo) {
				if isLocationOk(userId,campaign.Campaign.TargetGroup.Location) {
					contentType := app.GetFileTypeByPostId(campaign.Id)
					campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultiple(campaign,contentType))
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
func findNotMyMultipleCampaigns(oneTimeCampaigns []models.MultipleTimeCampaign, idPrimitive primitive.ObjectID) ([]models.MultipleTimeCampaign, error) {
	campaigns := []models.MultipleTimeCampaign{}

	for _, oneCampaign := range oneTimeCampaigns {

		if	oneCampaign.Campaign.User.Hex()!=idPrimitive.Hex() {
			campaigns = append(campaigns, oneCampaign)
		}
	}
	return campaigns, nil
}