package main

import (
	"campaigns/pkg/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"strings"
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

func (app *application) saveImageWithToken(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	token := vars["token"]
	campaignId := vars["campaignId"]
	userId :=getUserIdWithToken(token)
	if userId=="not" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	defer file.Close()
	//var path = "/var/lib/campaigns/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	campaignIdPrimitive, _ :=primitive.ObjectIDFromHex(campaignId)
	var image =models.Image {
		Media : path,
		UserId : userIdPrimitive,
		CampaignId : campaignIdPrimitive,
	}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}
func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("POGODIO UPIS SLIKE")
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	campaignId := vars["campaignId"]
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	defer file.Close()
	//var path = "/var/lib/campaigns/data/"+hander.Filename
	var path = "files/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)

	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	campaignIdPrimitive, _ :=primitive.ObjectIDFromHex(campaignId)
	var image =models.Image {
		Media : path,
		UserId : userIdPrimitive,
		CampaignId : campaignIdPrimitive,
	}

	insertResult, err  := app.images.Insert(image)


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New image has been created, id=%s", insertResult.InsertedID)
}
