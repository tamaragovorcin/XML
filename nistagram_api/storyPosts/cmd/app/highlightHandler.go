package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"storyPosts/pkg/dtos"
	"storyPosts/pkg/models"
	"strings"
)

func (app *application) getAllHighlights(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.highlights.All()
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


func (app *application) insertHighlight(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	userId := vars["userId"]
	var m dtos.HighlightDTO
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	var highlight = models.HighLight{
		User : userIdPrimitive,
		Name : m.Name,
		Stories : []models.StoryPost{},
	}

	insertResult, err := app.highlights.Insert(highlight)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New highlight have been created, id=%s", insertResult.InsertedID)

}

func (app *application) deleteHighlight(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := app.highlights.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}
func (app *application) getUsersHiglights(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------")
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	allHighlights, _ :=app.highlights.All()
	allImages,_ := app.images.All()
	usersHighlights,err :=findHighlightsByUserId(allHighlights,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	highlightsResponse := []dtos.HighlightsInfoDTO{}
	for _, highlight := range usersHighlights {

		images := highlight.Stories
		highlightsResponse = append(highlightsResponse, toResponseHighLights(highlight, images,allImages))
	}
	imagesMarshaled, err := json.Marshal(highlightsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func toResponseHighLights(highlight models.HighLight, storyPosts []models.StoryPost, images []models.Image) dtos.HighlightsInfoDTO {
	storiesInfoDtos := []dtos.StoryPostInfoDTO{}
	for _, storyPost := range storyPosts {
		image := getImageByStoryPost(images,storyPost.Id)
		storyPostInfoDTO :=toResponseStoryPost(storyPost,image)
		storiesInfoDtos = append(storiesInfoDtos,storyPostInfoDTO)
	}

	return dtos.HighlightsInfoDTO{
		Id: highlight.Id,
		Stories: storiesInfoDtos,
		Name : highlight.Name,
	}
}

func getImageByStoryPost(images []models.Image,highlightsId primitive.ObjectID) string {
	storyImage := models.Image{}

	for _, image := range images {
		if	image.PostId==highlightsId {
			storyImage = image
		}
	}
	return storyImage.Media
}

func findHighlightsByUserId(highlights []models.HighLight, idPrimitive primitive.ObjectID) ([]models.HighLight, error){
	highlightsUser := []models.HighLight{}

	for _, highlight := range highlights {
		if	highlight.User.String()==idPrimitive.String() {
			highlightsUser = append(highlightsUser, highlight)
		}
	}
	return highlightsUser, nil
}
func (app *application) insetStoryInHighlight(w http.ResponseWriter, r *http.Request) {

		var m dtos.HighlightStoryDTO
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			app.serverError(w, err)
		}

		highlight, err := app.highlights.FindByID(m.HighlightId)
		if highlight == nil {
			app.infoLog.Println("Hihglight not found")
		}
		storyPost, err := app.storyPosts.FindByID(m.StoryId)
		if storyPost == nil {
			app.infoLog.Println("Hihglight not found")
		}


		var highlightUpdate = models.HighLight{
			Id: m.HighlightId,
			User:highlight.User,
			Name : highlight.Name,
			Stories: append(highlight.Stories, *storyPost),
		}

		insertResult, err := app.highlights.Update(highlightUpdate)
		if err != nil {
			app.serverError(w, err)
		}
		app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) insetStoryAlbumInHighlight(w http.ResponseWriter, r *http.Request) {

	var m dtos.HighlightStoryAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	highlightAlbum, err := app.highlightsAlbum.FindByID(m.HighlightId)
	if highlightAlbum == nil {
		app.infoLog.Println("Hihglight not found")
	}
	storyAlbum, err := app.storyPosts.FindByID(m.StoryId)
	if storyAlbum == nil {
		app.infoLog.Println("Hihglight not found")
	}


	var highlightUpdate = models.HighLightAlbum{
		Id: m.HighlightId,
		User:highlightAlbum.User,
		Name : highlightAlbum.Name,
		Stories: append(highlightAlbum.Stories, *storyAlbum),
	}

	insertResult, err := app.highlightsAlbum.Update(highlightUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
