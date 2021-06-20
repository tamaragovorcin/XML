package main

import (
	"campaigns/pkg/dtos"
	"campaigns/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)


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


func isLocationOk(id string, location models.Location) bool {
	if location.Country=="" {
		return true
	}
	country := location.Country
	town :=location.Town
	street :=location.Street

	if country=="" {
		country = "n"
	}
	if town=="" {
		town = "n"
	}
	if street=="" {
		street = "n"
	}
	resp, err := http.Get("http://localhost:80/api/feedPosts/locationOk/"+id+"/"+country+"/"+town+"/"+street)
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
	if sb == "locationOk" {
		return true
	}

	return false
}

func isDateOfBirthOk(id string, one string, two string) bool {
	if one=="" || two==""{
		return true
	}
	resp, err := http.Get("http://localhost:80/api/users/api/user/dateOfBirthOk/"+id+"/"+one+"/"+two)
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
	if sb == "dateOfBirthOk" {
		return true
	}

	return false
}

func isGenderOk(id string, gender string) bool {
	if gender=="" {
		return true
	}
	resp, err := http.Get("http://localhost:80/api/users/api/user/genderOk/"+id+"/"+gender)
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
	if sb == "sameGender" {
		return true
	}

	return false
}


func(app *application) getBestInfluencers(w http.ResponseWriter, r *http.Request){


	allStatistics :=getAllStatistics(app)
	bestStatisticsDTOS := getAllInfluencersInStatisticsDTO(allStatistics)
	sort.SliceStable(bestStatisticsDTOS, func(i, j int) bool {
		return bestStatisticsDTOS[i].NumberOfPartnerships+ bestStatisticsDTOS[i].NumberOfClicks >
			bestStatisticsDTOS[j].NumberOfPartnerships+ bestStatisticsDTOS[j].NumberOfClicks
	})
	usernamesMarshaled, err := json.Marshal(bestStatisticsDTOS)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func getAllInfluencersInStatisticsDTO(statistics []models.Statistic) []dtos.BestStatisticsDTO {
	bestStatisticsDTO :=[]dtos.BestStatisticsDTO{}

	for _, statisticOne := range statistics {
		if userAlreadyHasDTO(statisticOne.Influencer, bestStatisticsDTO) {
				updateUserBestStatistic(bestStatisticsDTO,statisticOne)

		} else {
			username :=getUserUsername(statisticOne.Influencer)

			var dto = dtos.BestStatisticsDTO{
				UserId: statisticOne.Influencer,
				NumberOfClicks : statisticOne.NumberOfClicks,
				NumberOfPartnerships: 1,
				Username : username,
			}
			bestStatisticsDTO = append(bestStatisticsDTO, dto)
		}

	}

	return bestStatisticsDTO
}

func updateUserBestStatistic(dto []dtos.BestStatisticsDTO, one models.Statistic) {
	for _, statisticOne := range dto {
		if statisticOne.UserId.Hex()==one.Influencer.Hex() {
			statisticOne.NumberOfPartnerships+=1
			statisticOne.NumberOfClicks+=one.NumberOfClicks
		}
	}
}

func userAlreadyHasDTO(id primitive.ObjectID, bestStatistics []dtos.BestStatisticsDTO) bool {
	for _, statisticOne := range bestStatistics {
		if statisticOne.UserId.Hex()==id.Hex() {
			return true
		}
	}
	return false
}

func getAllStatistics(app *application) []models.Statistic {
	oneTimeCampaigns, _ := app.oneTimeCampaign.All()
	multipleTimeCampaigns, _ := app.multipleTimeCampaign.All()
	allStatistics :=[]models.Statistic{}
	for _, campaign := range oneTimeCampaigns {
		for _, statistic := range campaign.Campaign.Statistic {
			allStatistics = append(allStatistics, statistic)
		}
	}
	for _, campaign := range multipleTimeCampaigns {
		for _, statistic := range campaign.Campaign.Statistic {
			allStatistics = append(allStatistics, statistic)
		}
	}
	return allStatistics
}