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

	IdPrimitive, _ := primitive.ObjectIDFromHex(dto.Id)

	allCampaigns, _ :=app.multipleTimeCampaign.All()
	for _, camp := range allCampaigns {
		if camp.Id ==  IdPrimitive{
			if(camp.ModifiedTime.Add(24*time.Hour).Before(time.Now())){
				userPrimitive, _ := primitive.ObjectIDFromHex(dto.User)
				var number,_ = strconv.Atoi(dto.DesiredNumber)
				var campaign = models.Campaign{
					User : userPrimitive,
					Link : dto.Link,
					Description :dto.Description,
				}
				var oneTimeCampaign = models.MultipleTimeCampaign{
					Id : IdPrimitive,
					Campaign:   campaign,
					StartTime: dto.StartTime,
					EndTime : dto.EndTime,
					DesiredNumber:   number,
					ModifiedTime: time.Now(),
				}

				insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
				if err != nil {
					app.serverError(w, err)
				}
				app.infoLog.Printf("New content have been created, id=%s", insertResult.UpsertedID)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

			}
			app.infoLog.Printf("Campaign can not be modified, 24 hours have not elapsed since the last modification")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
		}

	}





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
		Type : dto.Type,
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
		AgentId : campaing.Campaign.User,
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
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
	}
	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		ModifiedTime: campaign.ModifiedTime,
		TimesShown: campaign.TimesShown,
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
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		TimesShown: campaign.TimesShown,
		ModifiedTime: campaign.ModifiedTime,
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
	typeString := vars["type"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.multipleTimeCampaign.All()
	campaigns,err :=findNotMyMultipleCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignMultipleDTO{}
	for _, campaign := range campaigns {
		if campaign.Campaign.Type==typeString {
			if isTimeForExposureMultiple(app,campaign,"agent") {
				if iAmFollowingThisUser(userId,campaign.Campaign.User.Hex()) {
					contentType := app.GetFileTypeByPostId(campaign.Id)
					campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultiple(campaign, contentType))
				}else {
					if isGenderOk(userId, campaign.Campaign.TargetGroup.Gender) {
						if isDateOfBirthOk(userId, campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
							if isLocationOk(userId, campaign.Campaign.TargetGroup.Location) {
								contentType := app.GetFileTypeByPostId(campaign.Id)
								campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultiple(campaign, contentType))
							}
						}
					}
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

func isTimeForExposureMultiple(app *application,campaign models.MultipleTimeCampaign, typeOfUser string) bool {
	stringDateOne := campaign.StartTime+"T11:45:26.371Z"
	stringDateTwo := campaign.EndTime+"T00:00:01.371Z"
	layout := "2006-01-02T15:04:05.000Z"

	timeDateOne, err := time.Parse(layout, stringDateOne)
	timeDateTwo, err := time.Parse(layout, stringDateTwo)

	if err != nil {
		fmt.Println(err)
	}
	timeNow :=time.Now().UTC().Add(2*time.Hour)

	if timeDateOne.Before(timeNow) && timeDateTwo.After(timeNow) {
		n:=(24*60)/campaign.DesiredNumber
		num:= n*(campaign.TimesShown+1)
		fmt.Println(n)

		month := timeNow.Month()
		year :=timeNow.Year()
		date := timeNow.Day()
		monthInt := int(month)
		monthString :=strconv.Itoa(monthInt)
		if monthInt<10 {
			monthString = "0"+monthString
		}

		dateString := strconv.Itoa(date)
		if date<10 {
			dateString = "0"+dateString
		}
		stringToday := strconv.Itoa(year)+"-"+monthString+"-"+dateString

		stringTodayDateTime:= stringToday+"T00:00:01.371Z"
		timeDateTimeToday, err := time.Parse(layout, stringTodayDateTime)

		if err != nil {
			fmt.Println(err)
		}
		var timeTime = timeDateTimeToday.Add(time.Duration(num)*time.Minute)

		before5 := time.Now().UTC().Add(-5*time.Minute + 2*time.Hour)
		after5 := time.Now().UTC().Add(5*time.Minute+2*time.Hour)

		if timeTime.Before(after5) && timeTime.After(before5) {
			if typeOfUser=="agent" {
				updateMultipleCampaignTimesShown(app,campaign.Id)
			}
			return true
		}
		return false
	}

	return false
}

func updateMultipleCampaignTimesShown(app *application, id primitive.ObjectID) {
	campaign, _ := app.multipleTimeCampaign.FindByID(id.Hex())
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: campaign.Campaign.Partnerships,
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		TimesShown: campaign.TimesShown +1,
		ModifiedTime: campaign.ModifiedTime,
	}
	_, _ = app.multipleTimeCampaign.Update(oneTimeCampaign)

}
func (app *application) getMultipleHomePagePromote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	typeString := vars["type"]

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.multipleTimeCampaign.All()
	campaigns,err :=findNotMyMultipleCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignMultipleDTO{}
	for _, campaign := range campaigns {
		if campaign.Campaign.Type==typeString {

			if isTimeForExposureMultiple(app,campaign,"promote") {

				if campaignHasPartnerships(campaign.Campaign.Partnerships) {
					for _, partnership := range campaign.Campaign.Partnerships {
						if partnership.Approved {
							if iAmFollowingThisUser(userId,partnership.Influencer.Hex()) {
								contentType := app.GetFileTypeByPostId(campaign.Id)
								campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultipleHomePage(campaign, contentType, partnership.Influencer))
							}else {
								if isGenderOk(userId, campaign.Campaign.TargetGroup.Gender) {
									if isDateOfBirthOk(userId, campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
										if isLocationOk(userId, campaign.Campaign.TargetGroup.Location) {

											contentType := app.GetFileTypeByPostId(campaign.Id)
											campaignResponse = append(campaignResponse, campaignToResponseInfluencerMultipleHomePage(campaign, contentType, partnership.Influencer))

										}
									}
								}
							}
						}
					}
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

func campaignToResponseInfluencerMultipleHomePage(campaing models.MultipleTimeCampaign, contentType string, influencer primitive.ObjectID) dtos.CampaignMultipleDTO {
	username :=getUserUsername(influencer)
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
		AgentId : influencer,
	}
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

func (app *application) likeMultipleCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.multipleTimeCampaign.FindByID(m.CampaignId.Hex())
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: campaign.Campaign.Partnerships,
		Likes : append(campaign.Campaign.Likes,m.UserId),
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		TimesShown: campaign.TimesShown +1,
		ModifiedTime: campaign.ModifiedTime,
	}

	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) dislikeMultipleCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.multipleTimeCampaign.FindByID(m.CampaignId.Hex())
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: campaign.Campaign.Partnerships,
		Likes : campaign.Campaign.Likes,
		Dislikes: append(campaign.Campaign.Dislikes,m.UserId),
		Comments: campaign.Campaign.Comments,
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		TimesShown: campaign.TimesShown +1,
		ModifiedTime: campaign.ModifiedTime,
	}

	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) commentMultipleCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.multipleTimeCampaign.FindByID(m.CampaignId.Hex())
	var comment = models.Comment{
		DateTime : time.Now(),
		Content : m.Content,
		Writer: m.UserId,
	}
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  : campaign.Campaign.Statistic,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: campaign.Campaign.Partnerships,
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: append(campaign.Campaign.Comments,comment),
	}


	var oneTimeCampaign = models.MultipleTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		StartTime: campaign.StartTime,
		EndTime : campaign.EndTime,
		DesiredNumber: campaign.DesiredNumber,
		TimesShown: campaign.TimesShown +1,
		ModifiedTime: campaign.ModifiedTime,
	}

	insertResult, err := app.multipleTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getLikesMultipleCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignLikes,err :=app.multipleTimeCampaign.FindByID(campaignId)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range campaignLikes.Campaign.Likes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}

	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getDislikesMultipleCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignLikes,err :=app.multipleTimeCampaign.FindByID(campaignId)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range campaignLikes.Campaign.Dislikes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}

	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getCommentsMultipleCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignComments,err :=app.multipleTimeCampaign.FindByID(campaignId)

	if err != nil {
		app.serverError(w, err)
	}

	commentsDtos :=getCommentDtos(campaignComments.Campaign.Comments)


	usernamesMarshaled, err := json.Marshal(commentsDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}


func (app *application) clickLinkMultipleCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.multipleTimeCampaign.FindByID(m.CampaignId.Hex())

	newStatistics :=newListStatistics(campaign.Campaign.Statistic,m)
	var campaignOne = models.Campaign{
		User : campaign.Campaign.User,
		TargetGroup : campaign.Campaign.TargetGroup,
		Statistic  :newStatistics,
		Link : campaign.Campaign.Link,
		Description :campaign.Campaign.Description,
		Partnerships: campaign.Campaign.Partnerships,
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
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
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func newListStatistics(statistics []models.Statistic, m dtos.CampaignReactionDTO)  []models.Statistic {
	statisticsNewList := []models.Statistic{}
	if statisticsForThisUserExists(statistics,m.UserId) {
		for _, statisticOne := range statistics {
			if statisticOne.Influencer.Hex()!=m.UserId.Hex() {
				statisticsNewList = append(statisticsNewList,statisticOne)
			} else {
				statisticOne.NumberOfClicks+=1
				statisticsNewList = append(statisticsNewList,statisticOne)
			}
		}
		return statisticsNewList
	} else {
		var statisticOne = models.Statistic{
			Influencer:     m.UserId,
			NumberOfClicks: 1,
		}
		statisticsNewList = append(statistics, statisticOne)

	}
	return statisticsNewList
}


func statisticsForThisUserExists(statistics []models.Statistic, id primitive.ObjectID) bool {
	for _, statisticOne := range statistics {
		if statisticOne.Influencer.Hex()==id.Hex() {
			return true
		}
	}
	return false
}