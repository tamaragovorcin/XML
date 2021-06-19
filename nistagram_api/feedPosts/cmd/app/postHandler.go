package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)


func (app *application) findIfLocationIsOk(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	country := vars["country"]
	city := vars["city"]
	street := vars["street"]

	idUserPrimitive, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}
	writing := ""
	if doesUserHavePostsWithThisAddress(idUserPrimitive,country,city,street,app){
		writing="locationOk"
	}else {
		writing = "notOk"
	}
	b, err := json.Marshal(writing)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func doesUserHavePostsWithThisAddress(userPrimitive primitive.ObjectID, country string, city string, street string, app *application) bool {
	allFeedPosts, _ := app.feedPosts.All()
	allAlbums, _ := app.albumFeeds.All()

	usersFeedPosts, _ := findFeedPostsByUserId(allFeedPosts, userPrimitive)
	usersAlbums, _ := findFeedAlbumsByUserId(allAlbums, userPrimitive)

	for _, feedPost := range usersFeedPosts {
		if	feedPost.Post.Location.Country==country {
			if city=="n" {
				return true
			} else if feedPost.Post.Location.Town==city {
				if street== "n" {
					return true
				} else if feedPost.Post.Location.Street==street {
					return true
				}
			}
		}
	}
	for _, feedPost := range usersAlbums {
		if	feedPost.Post.Location.Country==country {
			if city=="n" {
				return true
			} else if feedPost.Post.Location.Town==city {
				if street== "n" {
					return true
				} else if feedPost.Post.Location.Street==street {
					return true
				}
			}
		}
	}
	return false
}
