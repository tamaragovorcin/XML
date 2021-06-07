package main

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"storyPosts/pkg/dtos"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"storyPosts/pkg/models"
)

func (app *application) getAllStory(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ad, err := app.albumStories.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(ad)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("albumStory have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertAlbumStory(w http.ResponseWriter, req *http.Request) {

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
	var storyPost = models.AlbumStory{
		Post : post,
		OnlyCloseFriends : m.OnlyCloseFriends,

	}


	insertResult, err := app.albumStories.Insert(storyPost)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)

}

func (app *application) deleteStory(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.albumStories.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d albumStory(s)", deleteResult.DeletedCount)
}


func (app *application) getStoryAlbumsForHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allImages,_ := app.images.All()
	allPosts, _ :=app.albumStories.All()
	storiesForHomePage,err :=findStoryAlbumsForHomePage(allPosts,userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	storyAlbumsResponse := []dtos.StoryAlbumInfoHomePageDTO{}
	for _, albumPost := range storiesForHomePage {
		if iAmFollowingThisUser(userId,albumPost.Post.User.Hex()) {

			images, err := findAlbumByPostId(allImages, albumPost.Id)
			if err != nil {
				app.serverError(w, err)
			}
			userInList := getIndexInListOfUsersStoryAlbums(userIdPrimitive, storyAlbumsResponse)
			if userInList == -1 {
				userUsername := getUserUsername(albumPost.Post.User)
				userId := albumPost.Post.User
				albums := []dtos.StoryAlbumInfoDTO{}
				var dto = dtos.StoryAlbumInfoHomePageDTO{
					UserId:       userId,
					UserUsername: userUsername,
					Albums:       append(albums, toResponseAlbum(albumPost, images)),
				}
				storyAlbumsResponse = append(storyAlbumsResponse, dto)
			} else if userInList != -1 {
				existingDto := storyAlbumsResponse[userInList]
				existingDto.Albums = append(existingDto.Albums, toResponseAlbum(albumPost, images))
			}
		}
	}

	imagesMarshaled, err := json.Marshal(storyAlbumsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func getIndexInListOfUsersStoryAlbums(idPrimitive primitive.ObjectID, listAlbums []dtos.StoryAlbumInfoHomePageDTO) int {
	for num, story := range listAlbums {
		if story.UserId.String()==idPrimitive.String() {
			return num
		}
	}
	return -1
}
func findStoryAlbumsForHomePage(posts []models.AlbumStory, idPrimitive primitive.ObjectID) ([]models.AlbumStory, error) {
	storyPostsUser := []models.AlbumStory{}

	for _, storyPost := range posts {

		if	storyPost.Post.User.String()!=idPrimitive.String() && checkIfStoryIsInLast24h(storyPost.Post.DateTime){
			storyPostsUser = append(storyPostsUser, storyPost)
		}
	}
	//dodati uslov za pracenje!!!!!!!!!!!
	return storyPostsUser, nil
}