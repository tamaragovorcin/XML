package main

import (
	"bytes"
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
)

func (app *application) getAllCollections(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	movies, err := app.collections.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(movies)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("collections have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func findUsersCollectionByName(collections []models.Collection, idPrimitive primitive.ObjectID, Name string) (models.Collection, error){
	collectionUser := models.Collection{}

	for _, collection := range collections {
		if	collection.User.String()==idPrimitive.String() {
			if(collection.Name == Name) {
				collectionUser = collection
			}
		}
	}
	return collectionUser, nil
}
func findUsersCollectionByNameBoolean(collections []models.Collection, idPrimitive primitive.ObjectID, Name string) (bool, error){
	response := false
	for _, collection := range collections {
		if	collection.User.String()==idPrimitive.String() {
			if(collection.Name == Name) {
				response = true
			}
		}
	}
	return response, nil
}
func (app *application) findCollectionByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	idPrim, _ := primitive.ObjectIDFromHex(id)

	// Find movie by id
	m, err := app.collections.FindByID(idPrim)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("collections not found")
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

	app.infoLog.Println("Have been found a collections")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) addToFavourites(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	feedId := vars["feedId"]
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	feedIdPrimitive, _ := primitive.ObjectIDFromHex(feedId)
	allCollections, _ :=app.collections.All()
	allDataCollection,err :=findUsersCollectionByName(allCollections,userIdPrimitive,"All data")
	var savedPost = models.SavedPost{
		User: userIdPrimitive,
		FeedPost: feedIdPrimitive,
		Collection: allDataCollection,
	}


	insertResult, err := app.savedPosts.Insert(savedPost)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func (app *application) insertAllDataCollection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}


	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allCollections, _ :=app.collections.All()
	col, _ := findUsersCollectionByNameBoolean(allCollections,userIdPrimitive,"All data")
	if col == false{
	var collection = models.Collection{
		User : userIdPrimitive,
		Name: "All data",
		Posts : []models.FeedPost{},
	}

	insertResult, err := app.collections.Insert(collection)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
	}
}
func (app *application) insertCollection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	var m dtos.CollectionDTO
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	var collection = models.Collection{
		User : userIdPrimitive,
		Name : m.Name,
		Posts : []models.FeedPost{},
	}

	insertResult, err := app.collections.Insert(collection)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New highlight have been created, id=%s", insertResult.InsertedID)

}

func (app *application) deleteCollection(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.collections.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d collections(s)", deleteResult.DeletedCount)
}
func findCollectionsByUserId(collections []models.Collection, idPrimitive primitive.ObjectID) ([]models.Collection, error){
	collectionUser := []models.Collection{}

	for _, collection := range collections {
		if	collection.User.String()==idPrimitive.String() {
			collectionUser = append(collectionUser, collection)
		}
	}
	return collectionUser, nil
}
func (app *application) getUsersCollections(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allCollections, _ :=app.collections.All()
	allImages,_ := app.images.All()
	usersHighlights,err :=findCollectionsByUserId(allCollections,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	highlightsResponse := []dtos.CollectionInfoDTO{}
	for _, highlight := range usersHighlights {

		images := highlight.Posts
		highlightsResponse = append(highlightsResponse, toResponseCollections(highlight, images, allImages))
	}
	imagesMarshaled, err := json.Marshal(highlightsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func getImageByFeedPost(images []models.Image,highlightsId primitive.ObjectID) string {
	storyImage := models.Image{}

	for _, image := range images {
		if	image.PostId==highlightsId {
			storyImage = image
		}
	}
	return storyImage.Media
}
func toResponseFeedPost(storyPost models.FeedPost, image2 string) dtos.FeedPostInfoDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image, _, _ := image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}
	taggedPeople :=getTaggedPeople(storyPost.Post.Tagged)

	return dtos.FeedPostInfoDTO{
		Id: storyPost.Id,
		DateTime : strings.Split(storyPost.Post.DateTime.String(), " ")[0],
		Tagged :  taggedPeople,
		Location : locationToString(storyPost.Post.Location),
		Description : storyPost.Post.Description,
		Hashtags : hashTagsToString(storyPost.Post.Hashtags),
		Media : buffer.Bytes(),

	}
}
func toResponseCollections(highlight models.Collection, storyPosts []models.FeedPost, images []models.Image) dtos.CollectionInfoDTO {
	storiesInfoDtos := []dtos.FeedPostInfoDTO{}
	for _, storyPost := range storyPosts {
		image := getImageByFeedPost(images,storyPost.Id)
		storyPostInfoDTO :=toResponseFeedPost(storyPost,image)
		storiesInfoDtos = append(storiesInfoDtos,storyPostInfoDTO)
	}

	return dtos.CollectionInfoDTO{
		Id: highlight.Id,
		Posts: storiesInfoDtos,
		Name : highlight.Name,
	}
}

func (app *application) insetPostInCollection(w http.ResponseWriter, r *http.Request) {

	var m dtos.CollectionPostDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	collection, err := app.collections.FindByID(m.CollectionId)
	if collection == nil {
		app.infoLog.Println("Collection not found")
	}
	feedPosts, err := app.feedPosts.FindByID(m.PostId)
	if feedPosts == nil {
		app.infoLog.Println("Posts not found")
	}


	var collectionUpdate = models.Collection{
		Id: m.CollectionId,
		User:collection.User,
		Name : collection.Name,
		Posts: append(collection.Posts, *feedPosts),
	}

	insertResult, err := app.collections.Update(collectionUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
