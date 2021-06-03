package main

import (
	"bytes"
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

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

	// Find booking by id
	m, err := app.feedPosts.FindByID(id)
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
	vars := mux.Vars(req)
	userId := vars["userIdd"]
	var m dtos.FeedPostDTO
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}
	fmt.Println(res1)
	fmt.Println(res1)
	fmt.Println(userId)
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
	var feedPost = models.FeedPost{
		Post : post,
		Likes : nil,
		Dislikes: nil,
		Comments: nil,
	}


	insertResult, err := app.feedPosts.Insert(feedPost)
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

	// Delete booking by id
	deleteResult, err := app.feedPosts.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}
func (app *application) getUsersFeedPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allImages,_ := app.images.All()
	allPosts, _ :=app.feedPosts.All()
	usersFeedPosts,err :=findFeedPostsByUserId(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO{}
	for _, feedPost := range usersFeedPosts {

		images, err := findImageByPostId(allImages,feedPost.Id)
		if err != nil {
			app.serverError(w, err)
		}

		feedPostResponse = append(feedPostResponse, toResponse(feedPost, images.Media))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func toResponse(feedPost models.FeedPost, image2 string) dtos.FeedPostInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image, _, _ := image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}

	return dtos.FeedPostInfoDTO{
		Id: feedPost.Id,
		Comments: feedPost.Comments,
		Likes: feedPost.Likes,
		Dislikes: feedPost.Dislikes,

			DateTime : strings.Split(feedPost.Post.DateTime.String(), " ")[0],
			Tagged :feedPost.Post.Tagged,
			Location : feedPost.Post.Location,
			Description : feedPost.Post.Description,
			Hashtags : feedPost.Post.Hashtags,
			Media : buffer.Bytes(),

	}
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
	allImages,_ := app.images.All()
	locationFeedPosts, _ :=app.feedPosts.All()

	if country!="n" || city!="n" || street!="n" {
		locationFeedPosts,_ =findFeedPostsByLocation(locationFeedPosts,country,city,street)
	}
	feedPostResponse := []dtos.FeedPostInfoDTO{}
	for _, feedPost := range locationFeedPosts {

		images, err := findImageByPostId(allImages,feedPost.Id)
		if err != nil {
			app.serverError(w, err)
		}

		feedPostResponse = append(feedPostResponse, toResponse(feedPost, images.Media))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func findFeedPostsByLocation(posts []models.FeedPost, country string, city string, street string)([]models.FeedPost, error) {
	feedPostsLocation := []models.FeedPost{}

	for _, feedPost := range posts {
		print(feedPost.Post.Location.Country)
		print(country)
		print(city)
		print(street)
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
	return feedPostsLocation, nil
}

func (app *application) getFeedPostsByHashTags(w http.ResponseWriter, r *http.Request) {
	var hashtags dtos.HashTagDTO
	err := json.NewDecoder(r.Body).Decode(&hashtags)
	if err != nil {
		app.serverError(w, err)
	}
	allImages,_ := app.images.All()
	hashTagsFeedPosts, _ :=app.feedPosts.All()

	if hashtags.HashTags!="n" {
		hashTagsFeedPosts,_ =findFeedPostsByHashTags(hashTagsFeedPosts,parseHashTags(hashtags.HashTags))
	}

	feedPostResponse := []dtos.FeedPostInfoDTO{}
	for _, feedPost := range hashTagsFeedPosts {

		images, err := findImageByPostId(allImages,feedPost.Id)
		if err != nil {
			app.serverError(w, err)
		}

		feedPostResponse = append(feedPostResponse, toResponse(feedPost, images.Media))

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
		feedPostsHashTagsList := feedPost.Post.Hashtags
		if feedPostsHashTagsList !=nil {
			if postContainsAllHashTags(feedPostsHashTagsList,hashtags) {
				feedPostsHashTags = append(feedPostsHashTags, feedPost)
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