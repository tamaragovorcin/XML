package main

import (
	"bytes"
	"campaigns/pkg/dtos"
	"campaigns/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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
		NumberOfLikes : len(campaing.Campaign.Likes),
		NumberOfDislikes : len(campaing.Campaign.Dislikes),
		NumberOfComments : len(campaing.Campaign.Comments),
		Likes : getLikesDTOS(campaing.Campaign.Likes),
		Dislikes : getDislikesDTOS(campaing.Campaign.Dislikes),
		Comments: getCommentDtos(campaing.Campaign.Comments),
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
		NumberOfLikes : len(campaing.Campaign.Likes),
		NumberOfDislikes : len(campaing.Campaign.Dislikes),
		NumberOfComments : len(campaing.Campaign.Comments),
		Likes : getLikesDTOS(campaing.Campaign.Likes),
		Dislikes : getDislikesDTOS(campaing.Campaign.Dislikes),
		Comments: getCommentDtos(campaing.Campaign.Comments),
		TimesShownTotal : campaing.TimesShownTotal,
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



func (app *application) getStoryCampaignsForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages, _ := app.images.All()
	storiesForHomePage := findCampaignStoriesForHomePage(app, userIdPrimitive)

	storyPostsResponse := []dtos.StoryPostInfoHomePageDTO{}
	for _, storyPost := range storiesForHomePage {


						images, err := findImageByCampaignId(allImages, storyPost.CampaignId)
						if err != nil {
							app.serverError(w, err)
						}
						userInList := getIndexInListOfUsersStories(userIdPrimitive, storyPostsResponse)
						if userInList == -1 {
							userUsername := getUserUsername(storyPost.UserId)

							stories := []dtos.StoryPostInfoDTO{}
							var dto = dtos.StoryPostInfoHomePageDTO{
								Link : storyPost.Link,
								Type : storyPost.Type,
								CampaignId: storyPost.CampaignId,
								UserId: storyPost.UserId,
								UserUsername: userUsername,
								Stories:      append(stories, toResponseStoryPost2(storyPost, images.Media)),
							}
							storyPostsResponse = append(storyPostsResponse, dto)
						} else if userInList != -1 {
							existingDto := storyPostsResponse[userInList]
							existingDto.Stories = append(existingDto.Stories, toResponseStoryPost2(storyPost, images.Media))
						}
					}


	imagesMarshaled, err := json.Marshal(storyPostsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func toResponseStoryPost2(storyPost dtos.CampaignStoriesDTO, image2 string) dtos.StoryPostInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image,_,_:= image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}




	return dtos.StoryPostInfoDTO{
		Id: storyPost.CampaignId,
		Link : storyPost.Link,
		Media : buffer.Bytes(),

	}
}
func getIndexInListOfUsersStories(idPrimitive primitive.ObjectID, listStories []dtos.StoryPostInfoHomePageDTO) int {
	for num, story := range listStories {
		if story.UserId.String()==idPrimitive.String() {
			return num
		}
	}
	return -1
}
func findCampaignStoriesForHomePage(app *application, idPrimitive primitive.ObjectID) []dtos.CampaignStoriesDTO {
	oneTimeCampaigns, _ := app.oneTimeCampaign.All()
	multipleTimeCampaigns, _ := app.multipleTimeCampaign.All()
	allCampaigns :=[]dtos.CampaignStoriesDTO{}
	fmt.Println("1")
	for _, campaign := range oneTimeCampaigns {
		if	campaign.Campaign.User.Hex()!=idPrimitive.Hex() {

			if campaign.Campaign.Type == "story" {

				if isTimeForExposure(campaign.Time, campaign.Date) {

					if iAmFollowingThisUser(idPrimitive.Hex(),campaign.Campaign.User.Hex()) {

						var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
							CampaignId: campaign.Id,
							UserId:     campaign.Campaign.User,
							Type:      "oneTime",
							Link:       campaign.Campaign.Link,
						}
						allCampaigns = append(allCampaigns, CampaignStoriesDTO)
					} else {

						if isGenderOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Gender) {
							if isDateOfBirthOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
								if isLocationOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Location) {

									var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
										CampaignId: campaign.Id,
										UserId:     campaign.Campaign.User,
										Type:       "oneTime",
										Link:       campaign.Campaign.Link,
									}
									allCampaigns = append(allCampaigns, CampaignStoriesDTO)
								}
							}
						}
					}
				}
			}
		}
	}
	for _, campaign := range multipleTimeCampaigns {
		if	campaign.Campaign.User.Hex()!=idPrimitive.Hex() {

			if campaign.Campaign.Type == "story" {
				if isTimeForExposureMultiple(app,campaign,"agent") {
					if iAmFollowingThisUser(idPrimitive.Hex(), campaign.Campaign.User.Hex()) {
						var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
							CampaignId: campaign.Id,
							UserId:     campaign.Campaign.User,
							Type:       "multiple",
							Link:       campaign.Campaign.Link,
						}
						allCampaigns = append(allCampaigns, CampaignStoriesDTO)
					} else {
						if isGenderOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Gender) {
							if isDateOfBirthOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
								if isLocationOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Location) {

									var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
										CampaignId: campaign.Id,
										UserId:     campaign.Campaign.User,
										Type:       "multiple",
										Link:       campaign.Campaign.Link,
									}
									allCampaigns = append(allCampaigns, CampaignStoriesDTO)
								}
							}
						}
					}
				}

			}
		}
	}
	for _, campaign := range oneTimeCampaigns {
		if campaign.Campaign.Type=="story" {

			if isTimeForExposure(campaign.Time, campaign.Date) {

				if campaignHasPartnerships(campaign.Campaign.Partnerships) {
					for _, partnership := range campaign.Campaign.Partnerships {
						if partnership.Approved {
							if iAmFollowingThisUser(idPrimitive.Hex(), partnership.Influencer.Hex()) {
								var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
									CampaignId: campaign.Id,
									UserId:     partnership.Influencer,
									Type:      "oneTime",
									Link:       campaign.Campaign.Link,
								}
								allCampaigns = append(allCampaigns, CampaignStoriesDTO)
							} else {
								if isGenderOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Gender) {
									if isDateOfBirthOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
										if isLocationOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Location) {

											var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
												CampaignId: campaign.Id,
												UserId:     partnership.Influencer,
												Type:      "oneTime",
												Link:       campaign.Campaign.Link,
											}
											allCampaigns = append(allCampaigns, CampaignStoriesDTO)
										}
									}
								}
							}
						}
					}
				}
			}
		}
		for _, campaign := range multipleTimeCampaigns {
			if campaign.Campaign.Type == "story" {

				if isTimeForExposureMultiple(app,campaign,"promote") {

					if campaignHasPartnerships(campaign.Campaign.Partnerships) {
						for _, partnership := range campaign.Campaign.Partnerships {
							if partnership.Approved {
								if iAmFollowingThisUser(idPrimitive.Hex(), partnership.Influencer.Hex()) {
									var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
										CampaignId: campaign.Id,
										UserId:     partnership.Influencer,
										Type:       "multiple",
										Link:       campaign.Campaign.Link,
									}
									allCampaigns = append(allCampaigns, CampaignStoriesDTO)
								} else {
									if isGenderOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Gender) {
										if isDateOfBirthOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.DateOne, campaign.Campaign.TargetGroup.DateTwo) {
											if isLocationOk(idPrimitive.Hex(), campaign.Campaign.TargetGroup.Location) {

												var CampaignStoriesDTO = dtos.CampaignStoriesDTO{
													CampaignId: campaign.Id,
													UserId:     partnership.Influencer,
													Type:       "multiple",
													Link:       campaign.Campaign.Link,
												}
												allCampaigns = append(allCampaigns, CampaignStoriesDTO)
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
	}
	return allCampaigns
}

func iAmFollowingThisUser(logged string, userWithPost string) bool {

	postBody, _ := json.Marshal(map[string]string{
		"follower":  logged,
		"following": userWithPost,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://localhost:80/api/userInteraction/api/checkInteraction", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	sbbtext := strings.ToLower(strings.Trim(sb," \r\n"))

	if sbbtext=="true" {
		return true
	} else {
		return false
	}
}



