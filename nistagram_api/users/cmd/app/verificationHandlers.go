package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"users/pkg/dtos"
	"users/pkg/models"
)


func (app *application) insertVerification(w http.ResponseWriter, r *http.Request) {
	var m dtos.VerifyRequest

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	sb := m.Id
	sb = sb[1:]
	sb = sb[:len(sb)-1]
	Id, err := primitive.ObjectIDFromHex(m.Id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var verification = models.Verification{
		Name:     m.Name,
		LastName: m.LastName,
		User : Id,
		Approved: false,
		Category: m.Category,
	}

	insertResult, err := app.verifications.Insert(verification)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New verification have been created, id=%s", insertResult.InsertedID)
	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func (app *application) deleteVerification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.verifications.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d verifications(s)", deleteResult.DeletedCount)
}



func (app *application) acceptVerificationRequest(w http.ResponseWriter, r *http.Request) {

	var mm dtos.VerificationReactionDTO
	err := json.NewDecoder(r.Body).Decode(&mm)
	if err != nil {
		app.serverError(w, err)
	}
	verification, err := app.verifications.FindByID(mm.RequestId)

	_, err = app.verifications.Delete(mm.RequestId.Hex())

	user, err := app.users.FindByID(mm.UserId)
	if user == nil {
		app.infoLog.Println("User not found")
	}

	var userUpdate = models.User{
		Id: user.Id,
		ProfileInformation: user.ProfileInformation,
		Biography: user.Biography,
		Private: user.Private,
		Verified: true,
		Website: user.Website,
		Category : verification.Category,
	}

	insertResult, err := app.users.Update(userUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getAllRequestVerification(w http.ResponseWriter, r *http.Request) {

	allRequests, err  :=app.verifications.GetAll()
	allImages, err  :=app.images.All()

	if err != nil {
		app.serverError(w, err)
	}
	requestsFront := []dtos.RequestDTO{}
	for _, request := range allRequests {
		if err != nil {
			app.serverError(w, err)
		}
		image,_ :=findImageByVerificationId(allImages,request.Id)
		requestsFront = append(requestsFront, verificationRequestToResponse(request,image.Media))

	}

	imagesMarshaled, err := json.Marshal(requestsFront)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func verificationRequestToResponse(verification models.Verification, image2 string) dtos.RequestDTO {
	f, _ := os.Open(image2)
	defer f.Close()
	image, _, _ := image.Decode(f)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}
	imageBuffered :=buffer.Bytes()
	return dtos.RequestDTO{
		Id: verification.Id,
		Name: verification.Name,
		LastName: verification.LastName,
		Category: verification.Category,
		Media : imageBuffered,
		UserId : verification.User,
	}
}

