package main

import (
	"bytes"
	"context"
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"feedPosts/tracer"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"users/pkg/models"
)
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			fmt.Println("No Token Found")

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}


		authStringHeader := r.Header.Get("Authorization")
		if authStringHeader == "" {
			fmt.Errorf("Neki eror za auth")
		}
		authHeader := strings.Split(authStringHeader, "Bearer ")
		jwtToken := authHeader[1]

		token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("luna") , nil
		})

		if err != nil {
			fmt.Println("Your Token has been expired.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}



		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rolesString, _ := claims["roles"].(string)
			fmt.Println(rolesString)
			var tokenRoles []models.Role

			if err := json.Unmarshal([]byte(rolesString), &tokenRoles); err != nil {
				fmt.Println("Usercccc.")
			}



		} else{
			fmt.Println("User authorize fail.")
		}
	}


}
func (app *application) getAllFeedPosts(w http.ResponseWriter, r *http.Request) {
	bookings, err := app.feedPosts.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Contents have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findFeedPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idd, _ := primitive.ObjectIDFromHex(id)

	// Find booking by id
	m, err := app.feedPosts.FindByID(idd)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Booking not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a booking")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertFeedPost(w http.ResponseWriter, req *http.Request) {
	span := tracer.StartSpanFromRequest("insertUser", app.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling insert user at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	rt, err := decodeBody(ctx, req.Body)
	if err != nil {
		tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}




	vars := mux.Vars(req)
	userId := vars["userIdd"]
	//var m dtos.FeedPostDTO
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	/*err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}*/
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	id, _ := createPost(ctx, rt, userIdPrimitive, *app)



	//app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationPost/"+"Feed Post"+"/"+userId)
	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(id)
	w.Write(idMarshaled)
}



func parseHashTags(hashtags string) []string {
	if hashtags=="" {
		return nil
	}
	a := strings.Split(hashtags, "#")
	a = a[1:]
	return a
}

func (app *application) deleteFeedPost(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	reportId := vars["reportId"]
	// Delete booking by id
	postIdPrimitive, _ := primitive.ObjectIDFromHex(id)

	deleteResult, err := app.feedPosts.Delete(id)
	_, err2 := app.reports.Delete(reportId)

	if err != nil {
		app.serverError(w, err)
	}
	if err2 != nil {
		app.serverError(w, err)
	}
	allCollections, _ := app.collections.All()
	for _, collection := range allCollections {
		for _, post := range collection.Posts {
			if post.Id.Hex()==postIdPrimitive.Hex() {
				updateThisCollectionPost(collection,post.Id,app)
			}
		}
	}
	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}

func updateThisCollectionPost(collection models.Collection, id primitive.ObjectID, app *application) {
	var collectionUpdate = models.Collection{
		Id: collection.Id,
		User:collection.User,
		Name : collection.Name,
		Posts: getPostsAfterRemovingFeedPost(collection.Posts,id),
	}

	insertResult, _ := app.collections.Update(collectionUpdate)
	app.infoLog.Printf("Have been eliminated %d albumFeeds(s)", insertResult.UpsertedID)
}
func getPostsAfterRemovingFeedPost(posts []models.FeedPost, id primitive.ObjectID) []models.FeedPost {
	newPosts := []models.FeedPost{}
	for _, post := range posts {
		if post.Id.Hex()!=id.Hex() {
			newPosts=append(newPosts,post)
		}
	}
	return newPosts
}
func (app *application) getUsersFeedPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allPosts, _ :=app.feedPosts.All()
	usersFeedPosts,err :=findFeedPostsByUserId(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range usersFeedPosts {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(feedPost.Id)
		feedPostResponse = append(feedPostResponse, toResponse(feedPost,contentType))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
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

func(app *application) GetFileByPostId(w http.ResponseWriter, r *http.Request){
	fmt.Println("")
	vars := mux.Vars(r)
	feedId := vars["feedId"]
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
func findFeedByPostId(feeds []models.FeedPost, idFeedPost primitive.ObjectID) (models.FeedPost, error) {
	feedPost := models.FeedPost{}

	for _, feed := range feeds {
		if	feed.Id==idFeedPost {
			feedPost = feed
		}
	}
	return feedPost, nil
}
func(app *application) GetFileMessageByPostId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	userId := vars["userId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	allFeeds, _:= app.feedPosts.All()
	feedPost, _ := findFeedByPostId(allFeeds,feedIdPrim)
	privacy := getUsersPrivacy(feedPost.Post.User)
	if iAmFollowingThisUser(userId,feedPost.Post.User.Hex()) || privacy == "public" {
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
	response := dtos.ResponseDTO{
		Message: "You can not see this picture because you don't follow this user",
	}
	b, _ := json.Marshal(response)
	w.WriteHeader(http.StatusForbidden)
	w.Write(b)

	return




}
func toResponseImageMedia(feed models.Image) dtos.ImageMediaDTO {

	return dtos.ImageMediaDTO{
		Id: feed.Id,
		Media: feed.Media,
	}
}
func(app *application) getUsername(w http.ResponseWriter, r *http.Request){
	fmt.Println("")
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allFeeds,_ := app.feedPosts.All()
	feedPost, _ := findFeedByPostId(allFeeds,feedIdPrim)
	username := getUserUsername(feedPost.Post.User)

	imagesMarshaled, err := json.Marshal(username)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)



}
func(app *application) GetImagesByAlbumId(w http.ResponseWriter, r *http.Request){
	fmt.Println("")
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	images, _ := findImagesByPostId(allImages,feedIdPrim)

	imagesResponse := []dtos.ImageMediaDTO{}
	for _, img := range images {
		imagesResponse = append(imagesResponse, toResponseImageMedia(img))

	}
	imagesMarshaled, err := json.Marshal(imagesResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)



}
func(app *application) GetFilesByAlbumPostId(w http.ResponseWriter, r *http.Request){
	fmt.Println("")
	vars := mux.Vars(r)
	feedId := vars["feedId"]
	feedIdPrim, _ := primitive.ObjectIDFromHex(feedId)

	allImages,_ := app.images.All()
	for _, image := range allImages {
		if  image.Id == feedIdPrim{
			file, err:=os.Open(image.Media)
			if err!=nil{
				http.Error(w,"file not found",404)
				return
			}
			FileHeader:=make([]byte,512)
			file.Read(FileHeader)
			ContentType:= http.DetectContentType(FileHeader)
			FileStat,_:= file.Stat()
			FileSize:= strconv.FormatInt(FileStat.Size(),10)
			w.Header().Set("Content-Disposition", "attachment; filename="+image.Media)
			w.Header().Set("Content-Type", ContentType)
			w.Header().Set("Content-Length", FileSize)

			file.Seek(0,0)
			io.Copy(w,file)
			return

		}

	}







}
func toResponse(feedPost models.FeedPost, contentType string) dtos.FeedPostInfoDTO1 {
	taggedPeople :=getTaggedPeople(feedPost.Post.Tagged)
	return dtos.FeedPostInfoDTO1{
		Id: feedPost.Id,
		DateTime : strings.Split(feedPost.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
		Location : locationToString(feedPost.Post.Location),
		Description : feedPost.Post.Description,
		Hashtags : hashTagsToString(feedPost.Post.Hashtags),
		Username : "",
		ContentType: contentType,
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

func findFeedPostsByUserId(posts []models.FeedPost, idPrimitive primitive.ObjectID) ([]models.FeedPost, error){
	feedPostUser := []models.FeedPost{}

	for _, feedPost := range posts {
		if	feedPost.Post.User.String()==idPrimitive.String() {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}

func (app *application) getFeedPostsByLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	city :=vars["city"]
	street :=vars["street"]
	locationFeedPosts, _ :=app.feedPosts.All()

	if country!="n" || city!="n" || street!="n" {
		locationFeedPosts,_ =findFeedPostsByLocation(locationFeedPosts,country,city,street)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range locationFeedPosts {
		contentType := app.GetFileTypeByPostId(feedPost.Id)
		feedPostResponse = append(feedPostResponse, toResponse(feedPost,contentType))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func (app *application) getFeedPostByTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(user)
	tagsFeedAlbums, _ :=app.feedPosts.All()

	allImages,_ := app.images.All()
	tagsFeedAlbums =findFeedPostsByTags(tagsFeedAlbums,userIdPrimitive)

	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range tagsFeedAlbums {

		images, err := findImageByPostId(allImages,feedPost.Id)
		if err != nil {
			app.serverError(w, err)
		}
		userUsername :=getUserUsername(feedPost.Post.User)

		feedPostResponse = append(feedPostResponse, toResponseHomePage(feedPost, images.Media, userUsername))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func findFeedPostsByTags(albums []models.FeedPost, idPrimitive primitive.ObjectID) []models.FeedPost {
	listAlbums:=[]models.FeedPost{}
	for _, album := range albums {
		if userIsPublic(album.Post.User)==true {

			for _, tag := range album.Post.Tagged {
				if tag.String() == idPrimitive.String() {
					listAlbums = append(listAlbums, album)
				}
			}
		}
	}
	return listAlbums
}
func findFeedPostsByLocation(posts []models.FeedPost, country string, city string, street string)([]models.FeedPost, error) {
	feedPostsLocation := []models.FeedPost{}

	for _, feedPost := range posts {
		if userIsPublic(feedPost.Post.User)==true {
			if	feedPost.Post.Location.Country==country {
				if city=="n" {
					feedPostsLocation = append(feedPostsLocation, feedPost)
				} else if feedPost.Post.Location.Town==city {
					if street== "n" {
						feedPostsLocation = append(feedPostsLocation, feedPost)
					} else if feedPost.Post.Location.Street==street {
						feedPostsLocation = append(feedPostsLocation, feedPost)
					}
				}
			}
		}
		
	}
	return feedPostsLocation, nil
}

func userIsPublic(user primitive.ObjectID) bool {

	stringObjectID := user.Hex()
	resp, err := http.Get("http://localhost:80/api/users/api/user/privacy/"+stringObjectID)
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
	if sb == "public" {
		return true
	}

	return false
}

func (app *application) getFeedPostsByHashTags(w http.ResponseWriter, r *http.Request) {
	var hashtags dtos.HashTagDTO
	err := json.NewDecoder(r.Body).Decode(&hashtags)
	if err != nil {
		app.serverError(w, err)
	}
	hashTagsFeedPosts, _ :=app.feedPosts.All()

	if hashtags.HashTags!="n" {
		hashTagsFeedPosts,_ =findFeedPostsByHashTags(hashTagsFeedPosts,parseHashTags(hashtags.HashTags))
	}

	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range hashTagsFeedPosts {

		if err != nil {
			app.serverError(w, err)
		}
		contentType := app.GetFileTypeByPostId(feedPost.Id)
		feedPostResponse = append(feedPostResponse, toResponse(feedPost,contentType))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findFeedPostsByHashTags(posts []models.FeedPost, hashtags []string) ([]models.FeedPost, error) {
	feedPostsHashTags := []models.FeedPost{}

	for _, feedPost := range posts {
		if userIsPublic(feedPost.Post.User)==true {
			feedPostsHashTagsList := feedPost.Post.Hashtags
			if feedPostsHashTagsList != nil {
				if postContainsAllHashTags(feedPostsHashTagsList, hashtags) {
					feedPostsHashTags = append(feedPostsHashTags, feedPost)
				}
			}
		}

	}
	return feedPostsHashTags, nil
}

func postContainsAllHashTags(list []string, hashtags []string) bool {

	for _, hash := range hashtags {
		found :=false
		for _, itemInList := range list {
			if hash == itemInList {
				found= true
			}
		}
		if found==false {
			return false
		}
	}
	return true
}

func (app *application) getPhototsForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.feedPosts.All()
	postsForHomePage,err :=findFeedPostsForHomePage(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range postsForHomePage {
		if iAmFollowingThisUser(userId,feedPost.Post.User.Hex()) {
			if !iBlockedThisUser(userId, feedPost.Post.User.Hex()) {
				if (!iMutedThisUser(userId, feedPost.Post.User.Hex())) {

					images, err := findImageByPostId(allImages, feedPost.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(feedPost.Post.User)
					feedPostResponse = append(feedPostResponse, toResponseHomePage(feedPost, images.Media, userUsername))
				}
			}
		}
	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
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
func getUsersPrivacy(user primitive.ObjectID) string {

	stringObjectID := user.Hex()
	resp, err := http.Get("http://localhost:80/api/users/api/user/privacy/"+stringObjectID)
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
func findFeedPostsForHomePage(posts []models.FeedPost, idPrimitive primitive.ObjectID) ([]models.FeedPost, error) {
	feedPostUser := []models.FeedPost{}

	for _, feedPost := range posts {

		if	feedPost.Post.User.String()!=idPrimitive.String() {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}
func toResponseHomePage(feedPost models.FeedPost, image2 string, username string) dtos.FeedPostInfoDTO1 {
	taggedPeople :=getTaggedPeople(feedPost.Post.Tagged)
	file1, _:=os.Open(image2)
	FileHeader:=make([]byte,512)
	file1.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)
	return dtos.FeedPostInfoDTO1{
		Id: feedPost.Id,
		DateTime : strings.Split(feedPost.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
		Location : locationToString(feedPost.Post.Location),
		Description : feedPost.Post.Description,
		Hashtags : hashTagsToString(feedPost.Post.Hashtags),
		Username : username,
		ContentType: ContentType,
	}
}

func (app *application) likeTheFeedPost(w http.ResponseWriter, r *http.Request) {

	var m dtos.PostReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.feedPosts.FindByID(m.PostId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedPost.Post.User,
		DateTime : feedPost.Post.DateTime,
		Tagged : feedPost.Post.Tagged,
		Description: feedPost.Post.Description,
		Hashtags: feedPost.Post.Hashtags,
		Location : feedPost.Post.Location,
		Blocked : feedPost.Post.Blocked,
	}
	var feedPostUpdate = models.FeedPost{
		Id: feedPost.Id,
		Dislikes:feedPost.Dislikes,
		Comments : feedPost.Comments,
		Post : post,
		Likes: append(feedPost.Likes, m.UserId),
	}

	insertResult, err := app.feedPosts.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) dislikeTheFeedPost(w http.ResponseWriter, r *http.Request) {

	var m dtos.PostReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.feedPosts.FindByID(m.PostId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedPost.Post.User,
		DateTime : feedPost.Post.DateTime,
		Tagged : feedPost.Post.Tagged,
		Description: feedPost.Post.Description,
		Hashtags: feedPost.Post.Hashtags,
		Location : feedPost.Post.Location,
		Blocked : feedPost.Post.Blocked,
	}
	var feedPostUpdate = models.FeedPost{
		Id: feedPost.Id,
		Dislikes:append(feedPost.Dislikes, m.UserId),
		Comments : feedPost.Comments,
		Post : post,
		Likes: feedPost.Likes,
	}

	insertResult, err := app.feedPosts.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) commentTheFeedPost(w http.ResponseWriter, r *http.Request) {

	var m dtos.CommentReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.feedPosts.FindByID(m.PostId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}
	var post = models.Post{
		User : feedPost.Post.User,
		DateTime : feedPost.Post.DateTime,
		Tagged : feedPost.Post.Tagged,
		Description: feedPost.Post.Description,
		Hashtags: feedPost.Post.Hashtags,
		Location : feedPost.Post.Location,
		Blocked : feedPost.Post.Blocked,
	}
	var comment = models.Comment{
		DateTime : time.Now(),
		Content : m.Content,
		Writer: m.UserId,
	}
	var feedPostUpdate = models.FeedPost{
		Id: feedPost.Id,
		Dislikes:feedPost.Dislikes,
		Comments : append(feedPost.Comments, comment),
		Post : post,
		Likes: feedPost.Likes,
	}

	insertResult, err := app.feedPosts.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	resp, err := http.Get("http://localhost:80/api/users/api/sendNotificationComment/"+m.UserId.Hex()+"/"+post.User.Hex()+"/"+m.Content)
	fmt.Println(resp)
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getlikesFeedPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range likesForPost.Likes {

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

func (app *application) getdislikesFeedPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	likesDtos := []dtos.LikeDTO{}
	for _, user := range likesForPost.Dislikes {

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

func (app *application) getcommentsFeedPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postIdPrimitive, _ := primitive.ObjectIDFromHex(postId)


	likesForPost,err :=app.feedPosts.FindByID(postIdPrimitive)

	if err != nil {
		app.serverError(w, err)
	}

	commentsDtos :=getCommentDtos(likesForPost.Comments)


	usernamesMarshaled, err := json.Marshal(commentsDtos)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usernamesMarshaled)
}

func (app *application) getLikedPhotos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.feedPosts.All()
	postsForHomePage,err :=findLikedFeedPostsByUser(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range postsForHomePage {
		if iAmFollowingThisUser(userId,feedPost.Post.User.Hex()) {
			if !iBlockedThisUser(userId, feedPost.Post.User.Hex()) {

					images, err := findImageByPostId(allImages, feedPost.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(feedPost.Post.User)
					feedPostResponse = append(feedPostResponse, toResponseHomePage(feedPost, images.Media, userUsername))
				}
			}
		}


	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}


func (app *application) getDislikedPhotos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.feedPosts.All()
	postsForHomePage,err :=findDisLikedFeedPostsByUser(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO1{}
	for _, feedPost := range postsForHomePage {
		if iAmFollowingThisUser(userId,feedPost.Post.User.Hex()) {
			if !iBlockedThisUser(userId, feedPost.Post.User.Hex()) {

					images, err := findImageByPostId(allImages, feedPost.Id)
					if err != nil {
						app.serverError(w, err)
					}
					userUsername := getUserUsername(feedPost.Post.User)
					feedPostResponse = append(feedPostResponse, toResponseHomePage(feedPost, images.Media, userUsername))
				}
			}
	}


	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func findLikedFeedPostsByUser(posts []models.FeedPost, idPrimitive primitive.ObjectID) ([]models.FeedPost, error){
	feedPostUser := []models.FeedPost{}

	for _, feedPost := range posts {

		if	userLikedThePhoto(feedPost.Likes,idPrimitive) {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}

func userLikedThePhoto(ids []primitive.ObjectID, idPrimitive primitive.ObjectID) bool {
	for _, id := range ids {
		if id.Hex()==idPrimitive.Hex() {
			return true
		}
	}
	return false

}
func findDisLikedFeedPostsByUser(posts []models.FeedPost, idPrimitive primitive.ObjectID) ([]models.FeedPost, error){
	feedPostUser := []models.FeedPost{}

	for _, feedPost := range posts {

		if	userLikedThePhoto(feedPost.Dislikes,idPrimitive) {
			feedPostUser = append(feedPostUser, feedPost)
		}
	}
	return feedPostUser, nil
}