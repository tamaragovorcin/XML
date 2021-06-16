package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"users/pkg/dtos"
	"users/pkg/models"
)

func (app *application) getAllVerifications(w http.ResponseWriter, r *http.Request) {
	chats, err := app.verifications.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Verifications have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findVerificationByID(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	var i dtos.VerificationReactionDTO
	m, err := app.verifications.FindByID(i.RequestId)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Verification not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a verification")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

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



func (app *application) refreshVerification(w http.ResponseWriter, r *http.Request) {

	var m dtos.VerificationReactionDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	feedPost, err := app.verifications.FindByID(m.RequestId)
	if feedPost == nil {
		app.infoLog.Println("Feed Post not found")
	}

	var feedPostUpdate = models.Verification{
		Id: feedPost.Id,
		Name:feedPost.Name,
		LastName : feedPost.LastName,
		Category : feedPost.Category,
		Approved: true,
	}

	insertResult, err := app.verifications.Update(feedPostUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) getAllRequestVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("**********************")
	fmt.Println("LALALALLAL")
	allRequests, err  :=app.verifications.GetAll()
	fmt.Println(app.verifications.GetAll())
	if err != nil {
		app.serverError(w, err)
	}
	feedPostResponse := []dtos.RequestDTO{}
	for _, request := range allRequests {
		if err != nil {
			app.serverError(w, err)
		}
		fmt.Println(request.User)
		feedPostResponse = append(feedPostResponse, verificationRequestToResponse(request))

	}

	imagesMarshaled, err := json.Marshal(feedPostResponse)

	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func verificationRequestToResponse(collection models.Verification) dtos.RequestDTO {
	return dtos.RequestDTO{
		Id: collection.Id,
		Name: collection.Name,
		LastName: collection.LastName,
		Category: collection.Category,
	}
}
func(app *application) GetFileTypeByPostId(feedId primitive.ObjectID) string {
	/*resp, err := http.Get("http://localhost:4006/api/user/closeFriends/")
	allImages,_ := app.images.All()
	images, _ := findImageByPostId(allImages,feedId)

	file, _:=os.Open(images.Media)

	FileHeader:=make([]byte,512)
	file.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)

	return ContentType*/
return ""

}
