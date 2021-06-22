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
	"strings"
	"time"
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
		Type : dto.Type,
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
		AgentId : campaing.Campaign.User,
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
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
		Type : campaign.Campaign.Type,
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
		Likes : campaign.Campaign.Likes,
		Dislikes: campaign.Campaign.Dislikes,
		Comments: campaign.Campaign.Comments,
		Type : campaign.Campaign.Type,

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
	typeString := vars["type"]

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	campaigns,err :=findNotMyOneTimeCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range campaigns {
		if campaign.Campaign.Type==typeString {
			if isTimeForExposure(campaign.Time,campaign.Date){
				if isGenderOk(userId,campaign.Campaign.TargetGroup.Gender) {
					if isDateOfBirthOk(userId,campaign.Campaign.TargetGroup.DateOne,campaign.Campaign.TargetGroup.DateTwo) {
						if isLocationOk(userId,campaign.Campaign.TargetGroup.Location) {
							contentType := app.GetFileTypeByPostId(campaign.Id)
							campaignResponse = append(campaignResponse, campaignToResponseInfluencer(campaign,contentType))
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

func isTimeForExposure(timeDate string, dateDate string) bool {
	timeTime, err := datePlusTime(dateDate, timeDate+":54.016")
	if err != nil {
		fmt.Println(err)
	}

	before5 := time.Now().UTC().Add(-5*time.Minute + 2*time.Hour)
	after5 := time.Now().UTC().Add(5*time.Minute+2*time.Hour)

	if timeTime.Before(after5) && timeTime.After(before5) {
		return true
	}
	return false
}
func datePlusTime(date, timeOfDay string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.000", date+" "+timeOfDay)
}
func (app *application) getOneTimeHomePagePromote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	typeString := vars["type"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	campaigns,err :=findNotMyOneTimeCampaigns(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range campaigns {
		if campaign.Campaign.Type==typeString {
			if isTimeForExposure(campaign.Time,campaign.Date){

				if campaignHasPartnerships(campaign.Campaign.Partnerships) {
					if isGenderOk(userId, campaign.Campaign.TargetGroup.Gender) {
						if isDateOfBirthOk(userId, campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
							if isLocationOk(userId, campaign.Campaign.TargetGroup.Location) {
								for _, partnership := range campaign.Campaign.Partnerships {
									if partnership.Approved {
										contentType := app.GetFileTypeByPostId(campaign.Id)
										campaignResponse = append(campaignResponse, campaignToResponseInfluencerHomePage(campaign, contentType, partnership.Influencer))
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

func campaignToResponseInfluencerHomePage(campaing models.OneTimeCampaign, contentType string, influencer primitive.ObjectID) dtos.CampaignDTO {
		username :=getUserUsername(influencer)
		return dtos.CampaignDTO{
		Id: campaing.Id.Hex(),
		User: campaing.Campaign.User.Hex(),
		Description: campaing.Campaign.Description,
		Time: campaing.Time,
		Date: campaing.Date,
		Link: campaing.Campaign.Link,
		ContentType: contentType,
		AgentUsername: username,
		AgentId : influencer,
	}
}

func campaignHasPartnerships(partnerships []models.Partnership) bool {
	for _, oneCampaign := range partnerships {
		if	oneCampaign.Approved {
			return true
		}
	}
	return false
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

func (app *application) likeOneTimeCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.oneTimeCampaign.FindByID(m.CampaignId.Hex())
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
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) dislikeOneTimeCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.oneTimeCampaign.FindByID(m.CampaignId.Hex())
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
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) commentOneTimeCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.oneTimeCampaign.FindByID(m.CampaignId.Hex())
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
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getLikesOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignLikes,err :=app.oneTimeCampaign.FindByID(campaignId)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := getLikesDTOS(campaignLikes.Campaign.Likes )


	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func getLikesDTOS(likes []primitive.ObjectID) []dtos.LikeDTO {
	likesDtos := []dtos.LikeDTO{}
	for _, user := range likes {

		userUsername :=getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}
	return likesDtos
}

func (app *application) getDislikesOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignLikes,err :=app.oneTimeCampaign.FindByID(campaignId)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := getDislikesDTOS(campaignLikes.Campaign.Dislikes)

	usernamesMarshaled, err := json.Marshal(likesDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func getDislikesDTOS(likes []primitive.ObjectID) []dtos.LikeDTO {
	likesDtos := []dtos.LikeDTO{}
	for _, user := range likes {

		userUsername := getUserUsername(user)
		var like = dtos.LikeDTO{
			Username: userUsername,
		}

		likesDtos = append(likesDtos, like)

	}
	return likesDtos
}

func (app *application) getCommentsOneTimeCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignId := vars["campaignId"]


	campaignComments,err :=app.oneTimeCampaign.FindByID(campaignId)

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

func getCommentDtos(comments []models.Comment) []dtos.CommentDTO {
	commentDtos :=[]dtos.CommentDTO{}
	for _, comment := range comments {
		writerUsername :=getUserUsername(comment.Writer)
		var commentDto = dtos.CommentDTO{
			Content :comment.Content,
			Writer : writerUsername,
			DateTime: strings.Split(comment.DateTime.String(), " ")[0],

		}
		commentDtos = append(commentDtos, commentDto)
	}
	return commentDtos
}

func (app *application) clickLinkOneTimeCampaign(w http.ResponseWriter, r *http.Request) {

	var m dtos.CampaignReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	campaign, err := app.oneTimeCampaign.FindByID(m.CampaignId.Hex())

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


	var oneTimeCampaign = models.OneTimeCampaign{
		Id: campaign.Id,
		Campaign:   campaignOne,
		Time : campaign.Time,
		Date: campaign.Date,
	}

	insertResult, err := app.oneTimeCampaign.Update(oneTimeCampaign)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}