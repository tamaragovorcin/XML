package main

import (
	"campaigns/pkg/dtos"
	"campaigns/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (app *application) getAllCampaign(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.campaign.All()
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

func findCampaignByUserId(posts []models.OneTimeCampaign, idPrimitive primitive.ObjectID) ([]models.OneTimeCampaign, error){
	campaignsUser := []models.OneTimeCampaign{}

	for _, campaign := range posts {
		if	campaign.Campaign.User==idPrimitive {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}
func findMultipleTimeCampaignByUserId(posts []models.MultipleTimeCampaign, idPrimitive primitive.ObjectID) ([]models.MultipleTimeCampaign, error){
	campaignsUser := []models.MultipleTimeCampaign{}

	for _, campaign := range posts {
		if	campaign.Campaign.User==idPrimitive {
			campaignsUser = append(campaignsUser, campaign)
		}
	}
	return campaignsUser, nil
}
func(app *application) GetFileTypeByPostId(feedId primitive.ObjectID) string {
	allImages,_ := app.images.All()
	images, _ := findImageByCampaignId(allImages,feedId)

	file, _:=os.Open(images.Media)

	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)

	return ContentType


}
func findImageByCampaignId(images []models.Image, idFeedPost primitive.ObjectID) (models.Image, error) {
	imageFeedPost := models.Image{}

	for _, image := range images {
		if	image.CampaignId==idFeedPost {
			imageFeedPost = image
		}
	}
	return imageFeedPost, nil
}


func userInPartnershipRequests(partnerships []models.Partnership, idPrimitive primitive.ObjectID) bool {
	for _, partnership := range partnerships {

		if partnership.Influencer.Hex()==idPrimitive.Hex(){
			if partnership.Approved==false {
				return true
			}
		}
	}
	return false
}
func userInPartnership(partnerships []models.Partnership, idPrimitive primitive.ObjectID) bool {
	for _, partnership := range partnerships {

		if partnership.Influencer.Hex()==idPrimitive.Hex(){
			if partnership.Approved==true {
				return true
			}
		}
	}
	return false
}
func (app *application) getUsersCampaigns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.oneTimeCampaign.All()
	usersCampaigns,err :=findCampaignByUserId(allPosts,userIdPrimitive)
	allMultiple, _ :=app.multipleTimeCampaign.All()
	usersMultipeCampaigns,_ :=findMultipleTimeCampaignByUserId(allMultiple,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	campaignResponse := []dtos.CampaignDTO{}
	for _, campaign := range usersCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignToResponse(campaign,contentType))

	}
	for _, campaign := range usersMultipeCampaigns {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(campaign.Id)
		campaignResponse = append(campaignResponse, campaignMultipleToResponse(campaign,contentType))

	}

	imagesMarshaled, err := json.Marshal(campaignResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func(app *application) GetFileByCampaignId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	feedId := vars["campaignId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	images, err := findImageByCampaignId(allImages,feedIdPrim)

	file, err:=os.Open(images.Media)
	if err!=nil{
		http.Error(w,"file not found",404)
		return
	}


	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)
	FileStat,_:= file.Stat()
	FileSize:= strconv.FormatInt(FileStat.Size(),10)
	w.Header().Set("Content-Disposition", "attachment; filename="+images.Media)
	w.Header().Set("Content-Type", ContentType)
	w.Header().Set("Content-Length", FileSize)

	file.Seek(0,0)
	io.Copy(w,file)
	return



}
func campaignToResponse(campaing models.OneTimeCampaign, contentType string) dtos.CampaignDTO {
	return dtos.CampaignDTO{
		Id: campaing.Id.Hex(),
		User: campaing.Campaign.User.Hex(),
		Description: campaing.Campaign.Description,
		Time: campaing.Time,
		Date: campaing.Date,
		Link: campaing.Campaign.Link,
		ContentType: contentType,
		CampaignType:  "oneTime",
	}
}
func campaignMultipleToResponse(campaing models.MultipleTimeCampaign, contentType string) dtos.CampaignDTO {
	return dtos.CampaignDTO{
		Id: campaing.Id.Hex(),
		User: campaing.Campaign.User.Hex(),
		Description: campaing.Campaign.Description,
		StartTime: campaing.StartTime,
		EndTime: campaing.EndTime,
		Link: campaing.Campaign.Link,
		ContentType: contentType,
		CampaignType:  "multiple",
		DesiredNumber:  campaing.DesiredNumber,
	}
}
func (app *application) findByIDCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.campaign.FindByID(id)
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

func (app *application) insertCampaign(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Campaign
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.campaign.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteCampaign(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.campaign.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}
