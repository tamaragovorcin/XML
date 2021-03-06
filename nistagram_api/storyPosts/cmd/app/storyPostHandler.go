package main

import (
	"bytes"
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
	"storyPosts/pkg/dtos"
	"storyPosts/pkg/models"
	"strconv"
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
	listTagged := taggedUsersToPrimitiveObject(m)

	var post = models.Post{
		User : userIdPrimitive,
		DateTime : time.Now(),
		Tagged : listTagged,
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
	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationPost/"+"Story Post"+"/"+userId)
	fmt.Println(resp)
	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)

}
func taggedUsersToPrimitiveObject(m dtos.StoryPostDTO) []primitive.ObjectID {
	listTagged := []primitive.ObjectID{}
	for _, tag := range m.Tagged {
		primitiveTag, _ := primitive.ObjectIDFromHex(tag)

		listTagged = append(listTagged, primitiveTag)
	}
	return listTagged
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
	allStories, _ :=app.storyPosts.All()
	usersStoryPosts,err :=findStoriesByUserId(allStories,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	storyPostResponse := []dtos.StoryPostInfoDTO{}
	for _, storyPost := range usersStoryPosts {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(storyPost.Id)
		storyPostResponse = append(storyPostResponse, toResponseStoryPost1(storyPost,contentType))

	}

	imagesMarshaled, err := json.Marshal(storyPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func(app *application) GetFileByPostId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	feedId := vars["storyId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	images, err := findImageByPostId(allImages,feedIdPrim)

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
func(app *application) GetFileTypeByPostId(feedId primitive.ObjectID) string {
	allImages,_ := app.images.All()
	images, _ := findImageByPostId(allImages,feedId)

	file, _:=os.Open(images.Media)

	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)


	return ContentType


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
	taggedPeople :=getTaggedPeople(storyPost.Post.Tagged)

	file1, _:=os.Open(image2)
	FileHeader:=make([]byte,512)
	file1.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)

	return dtos.StoryPostInfoDTO{
		Id: storyPost.Id,
		DateTime : strings.Split(storyPost.Post.DateTime.String(), " ")[0],
		Tagged : taggedPeople,
		Location : locationToString(storyPost.Post.Location),
		Description : storyPost.Post.Description,
		Hashtags : hashTagsToString(storyPost.Post.Hashtags),
		ContentType: ContentType,
	}
}
func toResponseStoryPost1(storyPost models.StoryPost, contentType string) dtos.StoryPostInfoDTO {
	taggedPeople :=getTaggedPeople(storyPost.Post.Tagged)


	return dtos.StoryPostInfoDTO{
		Id: storyPost.Id,
		DateTime : strings.Split(storyPost.Post.DateTime.String(), " ")[0],
		Tagged : taggedPeople,
		Location : locationToString(storyPost.Post.Location),
		Description : storyPost.Post.Description,
		Hashtags : hashTagsToString(storyPost.Post.Hashtags),
		ContentType:  contentType,

	}
}

func getTaggedPeople(tagged []primitive.ObjectID) string {
	tagsString  := "Tagged: "
	for _, tag := range tagged {
		username :=getUserUsername(tag)
		tagsString+=username
		tagsString+=", "
	}
	return tagsString
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

		if	storyPost.Post.User.String()!=idPrimitive.String() && checkIfStoryIsInLast24h(storyPost.Post.DateTime){
			storyPostsUser = append(storyPostsUser, storyPost)
		}
	}
	//dodati uslov za pracenje!!!!!!!!!!!
	return storyPostsUser, nil
}

func checkIfStoryIsInLast24h(dateTime time.Time) bool {
	yesterday := time.Now().Add(-24*time.Hour)
	check := dateTime.After(yesterday)
	return check
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

func (app *application) getUsersStoryAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages, _ := app.images.All()
	allAlbums, _ := app.albumStories.All()
	usersStoryAlbums, err := findStoryAlbumsByUserId(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	storyAlbumResponse := []dtos.StoryAlbumInfoDTO{}
	for _, album := range usersStoryAlbums {

		images, err := findAlbumByPostId(allImages,album.Id)
		if err != nil {
			app.serverError(w, err)
		}

		storyAlbumResponse = append(storyAlbumResponse, toResponseAlbum(album, images))

	}

	imagesMarshaled, err := json.Marshal(storyAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findStoryAlbumsByUserId(albums []models.AlbumStory, idPrimitive primitive.ObjectID) ([]models.AlbumStory, error){
	feedAlbumsUser := []models.AlbumStory{}

	for _, album := range albums {
		if	album.Post.User.String()==idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}
func findAlbumByPostId(images []models.Image, idFeedAlbum primitive.ObjectID) ([]string, error) {
	imageAlbumPost := []string{}

	for _, image := range images {

		if	image.PostId==idFeedAlbum {
			imageAlbumPost= append(imageAlbumPost, image.Media)
		}
	}
	return imageAlbumPost, nil
}

func toResponseAlbum(feedAlbum models.AlbumStory, imageList []string) dtos.StoryAlbumInfoDTO {
	imagesBuffered := [][]byte{}
	for _, image2 := range imageList {
		f, _ := os.Open(image2)
		defer f.Close()
		image, _, _ := image.Decode(f)
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, image, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBuffered :=buffer.Bytes()
		imagesBuffered= append(imagesBuffered, imageBuffered)
	}

	taggedPeople :=getTaggedPeople(feedAlbum.Post.Tagged)

	return dtos.StoryAlbumInfoDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
		Location : locationToString(feedAlbum.Post.Location),
		Description : feedAlbum.Post.Description,
		Hashtags : hashTagsToString(feedAlbum.Post.Hashtags),
		Media : imagesBuffered,
		Username : "",

	}
}



func (app *application) getStoriesForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.storyPosts.All()
	storiesForHomePage,err :=findStoryPostsForHomePage(allPosts,userIdPrimitive)
	app.infoLog.Println("111")
	if err != nil {
		app.serverError(w, err)
	}
	storyPostsResponse := []dtos.StoryPostInfoHomePageDTO{}
	for _, storyPost := range storiesForHomePage {

		if iAmFollowingThisUser(userId, storyPost.Post.User.Hex()) {
			app.infoLog.Println("22")

			if storyCanBeSeen(storyPost, userIdPrimitive) == true {
				app.infoLog.Println("333")

				if !iBlockedThisUser(userId, storyPost.Post.User.Hex()) {
					if !iMutedThisUser(userId, storyPost.Post.User.Hex()) {

						images, err := findImageByPostId(allImages, storyPost.Id)
						if err != nil {
							app.serverError(w, err)
						}
						userInList := getIndexInListOfUsersStories(userIdPrimitive, storyPostsResponse)
						if userInList == -1 {
							userUsername := getUserUsername(storyPost.Post.User)

							stories := []dtos.StoryPostInfoDTO{}
							var dto = dtos.StoryPostInfoHomePageDTO{
								Id : storyPost.Id,
								UserId:       storyPost.Post.User,
								UserUsername: userUsername,
								Stories:      append(stories, toResponseStoryPost2(storyPost, images.Media)),
								CloseFriends: storyPost.OnlyCloseFriends,
							}
							storyPostsResponse = append(storyPostsResponse, dto)
						} else if userInList != -1 {
							existingDto := storyPostsResponse[userInList]
							existingDto.Stories = append(existingDto.Stories, toResponseStoryPost2(storyPost, images.Media))
						}
					}
					images, err := findImageByPostId(allImages, storyPost.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userInList := getIndexInListOfUsersStories(userIdPrimitive, storyPostsResponse)
					if userInList == -1 {
						userUsername := getUserUsername(storyPost.Post.User)

						stories := []dtos.StoryPostInfoDTO{}
						var dto = dtos.StoryPostInfoHomePageDTO{
							UserId:       storyPost.Post.User,
							UserUsername: userUsername,
							Stories:      append(stories, toResponseStoryPost2(storyPost, images.Media)),
							CloseFriends: storyPost.OnlyCloseFriends,
						}
						storyPostsResponse = append(storyPostsResponse, dto)
					} else if userInList != -1 {
						existingDto := storyPostsResponse[userInList]
						existingDto.Stories = append(existingDto.Stories, toResponseStoryPost2(storyPost, images.Media))
					}
				}
			}
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
func(app *application) getUsername(w http.ResponseWriter, r *http.Request){
	fmt.Println("")
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allFeeds,_ := app.storyPosts.All()
	feedPost, _ := findStoryByPostId(allFeeds,feedIdPrim)
	username := getUserUsername(feedPost.Post.User)

	imagesMarshaled, err := json.Marshal(username)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)



}
func getUsersPrivacy(user primitive.ObjectID) string {

	stringObjectID := user.Hex()
	resp, err := http.Get("http://localhost:80/api/users/api/user/privacy/" + stringObjectID)
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
	fmt.Println("PRIVACY:")
	fmt.Println(sb)
	return sb
}
func(app *application) GetFileMessageByPostId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	userId := vars["userId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	allFeeds, _:= app.storyPosts.All()
	feedPost, _ := findStoryByPostId(allFeeds,feedIdPrim)
	privacy := getUsersPrivacy(feedPost.Post.User)
	if iAmFollowingThisUser(userId,feedPost.Post.User.Hex()) || privacy == "public" {
		if time.Now().Before(feedPost.Post.DateTime.Add(24*time.Hour)){
			fmt.Println("VRIJEME STORY POSTA")
			images, err := findImageByPostId(allImages, feedIdPrim)

			file, err := os.Open(images.Media)
			if err != nil {
				http.Error(w, "file not found", 404)
				return
			}

			FileHeader := make([]byte, 512)
			file.Read(FileHeader)
			ContentType := http.DetectContentType(FileHeader)
			FileStat, _ := file.Stat()
			FileSize := strconv.FormatInt(FileStat.Size(), 10)
			w.Header().Set("Content-Disposition", "attachment; filename="+images.Media)
			w.Header().Set("Content-Type", ContentType)
			w.Header().Set("Content-Length", FileSize)

			file.Seek(0, 0)
			io.Copy(w, file)
			return
		}
		}
		response := dtos.ResponseDTO{
			Message: "You can not see this picture because you don't follow this user",
		}
		b, _ := json.Marshal(response)
		w.WriteHeader(http.StatusForbidden)
		w.Write(b)

		return




}
func findStoryByPostId(feeds []models.StoryPost, idFeedPost primitive.ObjectID) (models.StoryPost, error) {
	feedPost := models.StoryPost{}

	for _, feed := range feeds {
		if	feed.Id==idFeedPost {
			feedPost = feed
		}
	}
	return feedPost, nil
}
func toResponseStoryPost2(storyPost models.StoryPost, image2 string) dtos.StoryPostInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image,_,_:= image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}
	taggedPeople :=getTaggedPeople(storyPost.Post.Tagged)




	return dtos.StoryPostInfoDTO{
		Id: storyPost.Id,
		DateTime : strings.Split(storyPost.Post.DateTime.String(), " ")[0],
		Tagged : taggedPeople,
		Location : locationToString(storyPost.Post.Location),
		Description : storyPost.Post.Description,
		Hashtags : hashTagsToString(storyPost.Post.Hashtags),
		Media : buffer.Bytes(),

	}
}
func getListCloseFriends(id string) []string { //id usera ciji je stori

	resp, err := http.Get("http://localhost:80/api/users/api/user/closeFriends/"+id)
	log.Println("unable to encode image.", resp)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var listStrings []string
	sb := string(body)
	if sb!="" {
		listStrings =strings.Split(sb, ",")
	}
	return listStrings
}
func getListOfBlockedUsers(id string) []string { //id usera ciji je stori

	resp, err := http.Get("http://localhost:4006/api/user/blockedUsers/"+id)
	log.Println("unable to encode image.", resp)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var listStrings []string
	sb := string(body)
	if sb!="" {
		listStrings =strings.Split(sb, ",")
	}
	return listStrings
}
func getListOfMutedUsers(id string) []string { //id usera ciji je stori

	resp, err := http.Get("http://localhost:4006/api/user/mutedUsers/"+id)
	log.Println("unable to encode image.", resp)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var listStrings []string
	sb := string(body)
	if sb!="" {
		listStrings =strings.Split(sb, ",")
	}
	return listStrings
}



func userIsCloseFriends(user2 string, ids []string) bool { // svoj id
	for index, id := range ids {
		if index == 0 {
			id = id[1:]
		}
		if index == len(ids)-1 {
			id = id[:len(id)-1]
		}
		if strings.ToLower(strings.Trim(id," \r\n")) == strings.ToLower(strings.Trim(user2," \r\n")) {
			return true
		}
	}
	return false
}
func userIsBlocked(user2 string, ids []string) bool { // svoj id
	for index, id := range ids {
		if index == 0 {
			id = id[1:]
		}
		if index == len(ids)-1 {
			id = id[:len(id)-1]
		}
		if strings.ToLower(strings.Trim(id," \r\n")) == strings.ToLower(strings.Trim(user2," \r\n")) {
			return true
		}
	}
	return false
}
func userIsMuted(user2 string, ids []string) bool { // svoj id
	for index, id := range ids {
		if index == 0 {
			id = id[1:]
		}
		if index == len(ids)-1 {
			id = id[:len(id)-1]
		}
		if strings.ToLower(strings.Trim(id," \r\n")) == strings.ToLower(strings.Trim(user2," \r\n")) {
			return true
		}
	}
	return false
}

func storyCanBeSeen(post models.StoryPost, idPrimitive primitive.ObjectID) bool {


	if post.OnlyCloseFriends==false {return true}
	userId := post.Post.User

	closeFriends := getListCloseFriends(userId.Hex())
	if userIsCloseFriends(idPrimitive.Hex(), closeFriends){
		return true
	}
	return false
}
func iBlockedThisUser(logged string, userWithPost string) bool {
	userId := userWithPost

	blockedUsers := getListOfBlockedUsers(logged)
	if userIsBlocked(userId, blockedUsers){
		return true
	}
	return false
}
func iMutedThisUser(logged string, userWithPost string) bool {
	fmt.Println("POGODIO")
	userId := userWithPost

	blockedUsers := getListOfMutedUsers(logged)
	for _, s := range blockedUsers {
fmt.Println("lalalala")
		fmt.Println(s)

	}

	if userIsMuted(userId, blockedUsers){
		return true
	}
	return false
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