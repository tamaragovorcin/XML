package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"storyPosts/pkg/dtos"
	"storyPosts/pkg/models"
	"strings"
	"time"
)

func (app *application) getAllStoryPosts(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.storyPosts.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Contents have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}



func (app *application) insertStoryPost(w http.ResponseWriter, req *http.Request) {


	vars := mux.Vars(req)
	userId := vars["userId"]
	var m dtos.StoryPostDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	var post = models.Post{
		User : userIdPrimitive,
		DateTime : time.Now(),
		Tagged : m.Tagged,
		Description: m.Description,
		Hashtags: parseHashTags(m.Hashtags),
		Location : m.Location,
		Blocked : false,
	}
	var storyPost = models.StoryPost{
		OnlyCloseFriends : m.OnlyCloseFriends,
		Post : post,
	}


	insertResult, err := app.storyPosts.Insert(storyPost)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)

}
func parseHashTags(hashtags string) []string {
	a := strings.Split(hashtags, "#")
	a = a[1:]
	return a
}
func (app *application) deleteStoryPost(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.storyPosts.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}


func (app *application) getUsersStories(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages,_ := app.images.All()
	allStories, _ :=app.storyPosts.All()
	usersStoryPosts,err :=findStoriesByUserId(allStories,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	storyPostResponse := []dtos.StoryPostInfoDTO{}
	for _, storyPost := range usersStoryPosts {

		images, err := findImageByPostId(allImages,storyPost.Id)
		if err != nil {
			app.serverError(w, err)
		}

		storyPostResponse = append(storyPostResponse, toResponseStoryPost(storyPost, images.Media))

	}

	imagesMarshaled, err := json.Marshal(storyPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}


func findImageByPostId(images []models.Image, id primitive.ObjectID) (models.Image, error) {
	imageStoryPost := models.Image{}

	for _, image := range images {
		if	image.PostId==id {
			imageStoryPost = image
		}
	}
	return imageStoryPost, nil
}

func findStoriesByUserId(stories []models.StoryPost, idPrimitive primitive.ObjectID) ([]models.StoryPost,error) {
	storyPostsUser := []models.StoryPost{}

	for _, storyPost := range stories {
		if	storyPost.Post.User.String()==idPrimitive.String() {
			storyPostsUser = append(storyPostsUser, storyPost)
		}
	}
	return storyPostsUser, nil
}
func toResponseStoryPost(storyPost models.StoryPost, image2 string) dtos.StoryPostInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image, _, _ := image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}

	return dtos.StoryPostInfoDTO{
		Id: storyPost.Id,
		DateTime : strings.Split(storyPost.Post.DateTime.String(), " ")[0],
		Tagged :  storyPost.Post.Tagged,
		Location : locationToString(storyPost.Post.Location),
		Description : storyPost.Post.Description,
		Hashtags : hashTagsToString(storyPost.Post.Hashtags),
		Media : buffer.Bytes(),

	}
}

func locationToString(location models.Location) string {
	if location.Country=="" {
		return ""
	}else if location.Country!="" && location.Town=="" {
		return "Location: " +location.Country
	} else if location.Country!="" && location.Town!="" && location.Street==""{
		return "Location: " + location.Country + ", " + location.Town
	}
	return "Location: " + location.Country + ", " + location.Town + ", " + location.Street

}

func hashTagsToString(hashtags []string) string {
	hashTagString :=""
	for _, hash := range hashtags {
		hashTagString+="#"+hash
	}
	return hashTagString
}

func (app *application) getStoriesForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.storyPosts.All()
	storiesForHomePage,err :=findStoryPostsForHomePage(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	storyPostsResponse := []dtos.StoryPostInfoHomePageDTO{}
	for _, storyPost := range storiesForHomePage {

		images, err := findImageByPostId(allImages,storyPost.Id)
		if err != nil {
			app.serverError(w, err)
		}
		userInList :=getIndexInListOfUsersStories(userIdPrimitive,storyPostsResponse)
		if userInList==-1 {
			userUsername := getUserUsername(storyPost.Post.User)
			userId := storyPost.Post.User
			stories := []dtos.StoryPostInfoDTO{}
			var dto = dtos.StoryPostInfoHomePageDTO{
				UserId:       userId,
				UserUsername: userUsername,
				Stories:      append(stories, toResponseStoryPost(storyPost, images.Media)),
			}
			storyPostsResponse = append(storyPostsResponse, dto)
		}else if userInList!=-1 {
			existingDto :=storyPostsResponse[userInList]
			existingDto.Stories = append(existingDto.Stories, toResponseStoryPost(storyPost, images.Media))
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
func getIndexInListOfUsersStories(idPrimitive primitive.ObjectID, listStories []dtos.StoryPostInfoHomePageDTO) int {
	for num, story := range listStories {
		if story.UserId.String()==idPrimitive.String() {
			return num
		}
	}
	return -1
}
func findStoryPostsForHomePage(posts []models.StoryPost, idPrimitive primitive.ObjectID) ([]models.StoryPost, error) {
	storyPostsUser := []models.StoryPost{}

	for _, storyPost := range posts {

		if	storyPost.Post.User.String()!=idPrimitive.String() && checkIfStoryIsInLast24h(storyPost){
			storyPostsUser = append(storyPostsUser, storyPost)
		}
	}
	//dodati uslov za pracenje!!!!!!!!!!!
	return storyPostsUser, nil
}

func checkIfStoryIsInLast24h(post models.StoryPost) bool {
	yesterday := time.Now().Add(-24*time.Hour)
	check := post.Post.DateTime.After(yesterday)
	return check
}

func getUserUsername(user primitive.ObjectID) string {

	stringObjectID := user.Hex()
	resp, err := http.Get("http://localhost:4006/api/user/username/"+stringObjectID)
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