package main

import (
	"campaigns/pkg/dtos"
	"campaigns/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
	fmt.Println(dto.Link)
	fmt.Println(dto.Description)
	fmt.Println(dto.Id)

	var campaign = models.Campaign{
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
	fmt.Println(dto.User)
	fmt.Println(dto.Link)
	fmt.Println(dto.Description)
	fmt.Println(dto.PartnershipsRequests)
	fmt.Println(dto.TargetGroup.Gender)

	var campaign = models.Campaign{
		User : userIdPrimitive,
		TargetGroup : dto.TargetGroup,
		Statistic  :[]models.Statistic{},
		Link : dto.Link,
		Description :dto.Description,
		Partnerships :getPartnerships(dto.PartnershipsRequests),
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
