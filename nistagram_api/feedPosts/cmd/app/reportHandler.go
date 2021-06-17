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
)

func (app *application) getAllFeedReports(w http.ResponseWriter, r *http.Request) {
	allReports, _ :=app.reports.GetAll()
	allImages,_ := app.images.All()

	feedPostResponse := []dtos.FeedPostInfoReportDTO{}
	for _, report := range allReports {
		if report.Type=="post" {
			feedPost, err := app.feedPosts.FindByID(report.Post)

			images, err := findImageByPostId(allImages, feedPost.Id)
			if err != nil {
				app.serverError(w, err)
			}

			userUsername := getUserUsername(feedPost.Post.User)
			feedPostResponse = append(feedPostResponse, toResponseFeedReport(feedPost, images.Media, userUsername, report.Id))
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

func (app *application) getAllAlbumReports(w http.ResponseWriter, r *http.Request) {
	allReports, _ :=app.reports.GetAll()
	allImages,_ := app.images.All()

	feedAlbumsResponse := []dtos.FeedAlbumInfoReportDTO{}
	for _, report := range allReports {
		if report.Type=="album" {
			album, err := app.albumFeeds.FindByID(report.Post)

			images, err := findAlbumByPostId(allImages,album.Id)
			if err != nil {
				app.serverError(w, err)
			}

			userUsername := getUserUsername(album.Post.User)
			feedAlbumsResponse = append(feedAlbumsResponse, toResponseAlbumReport(album, images,userUsername,report.Id))
		}
	}
	imagesMarshaled, err := json.Marshal(feedAlbumsResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)

}
func toResponseFeedReport(feedPost *models.FeedPost, image2 string, username string, reportId primitive.ObjectID) dtos.FeedPostInfoReportDTO {
	taggedPeople :=getTaggedPeople(feedPost.Post.Tagged)
	file1, _:=os.Open(image2)
	FileHeader:=make([]byte,512)
	file1.Read(FileHeader)
	ContentType:= http.DetectContentType(FileHeader)
	return dtos.FeedPostInfoReportDTO{
		Id: feedPost.Id,
		UserId : feedPost.Post.User,
		DateTime : strings.Split(feedPost.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
		Location : locationToString(feedPost.Post.Location),
		Description : feedPost.Post.Description,
		Hashtags : hashTagsToString(feedPost.Post.Hashtags),
		Username : username,
		ContentType: ContentType,
		ReportId : reportId,
	}
}


func toResponseAlbumReport(feedAlbum *models.AlbumFeed, imageList []string, username string, reportId primitive.ObjectID) dtos.FeedAlbumInfoReportDTO {
	imagesBuffered := [][]byte{}
	for _, image2 := range imageList {
		f, _ := os.Open(image2)
		defer f.Close()
		image, _, _ := image.Decode(f)
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, image, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBuffered :=buffer.Bytes()
		imagesBuffered= append(imagesBuffered, imageBuffered)
	}

	taggedPeople :=getTaggedPeople(feedAlbum.Post.Tagged)

	return dtos.FeedAlbumInfoReportDTO{
		Id: feedAlbum.Id,
		UserId : feedAlbum.Post.User,
		DateTime : strings.Split(feedAlbum.Post.DateTime.String(), " ")[0],
		Tagged :taggedPeople,
		Location : locationToString(feedAlbum.Post.Location),
		Description : feedAlbum.Post.Description,
		Hashtags : hashTagsToString(feedAlbum.Post.Hashtags),
		Media : imagesBuffered,
		Username : username,
		ReportId : reportId,
	}
}
func (app *application) reportFeedPost(w http.ResponseWriter, req *http.Request)  {
	var m dtos.ReportDTO

	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	var report = models.Report{
		ComplainingUser: m.UserId,
		Post: m.PostId,
		Type : m.Type,
	}
	insertResult, err := app.reports.Insert(report)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}



func (app *application) deleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	deleteResult, err := app.reports.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d content(s)", deleteResult.DeletedCount)
}
func (app *application) removeEverythingFromUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(id)
	app.infoLog.Printf(id)

	removeFromFeedPosts(userIdPrimitive,app)
	removeFromFeedAlbums(userIdPrimitive,app)

	removeFromCollection(userIdPrimitive,app)
	removeFromCollectionAlbum(userIdPrimitive,app)

	removeFromOthersUsersCollection(userIdPrimitive,app)
	removeFromOthersUsersCollectionAlbum(userIdPrimitive,app)

	removeFromLikes(userIdPrimitive,app)
	removeFromDislikes(userIdPrimitive,app)
	removeFromComments(userIdPrimitive,app)

	removeFromLikesAlbum(userIdPrimitive,app)
	removeFromDislikesAlbum(userIdPrimitive,app)
	removeFromCommentsAlbum(userIdPrimitive,app)

}

func removeFromOthersUsersCollectionAlbum(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.collectionAlbums.All()
	for _,post := range allPosts {
		for _,singlePost := range post.Albums {

			if singlePost.Post.User.Hex()==idPrimitive.Hex() {
				newListAlbums :=removeAlbumFromCollection(post.Albums,singlePost.Id)
				var collectionUpdate = models.CollectionAlbum{
					Id: post.Id,
					User:post.User,
					Name : post.Name,
					Albums: newListAlbums,
				}
				_,_= app.collectionAlbums.Update(collectionUpdate)

			}
		}
	}
}

func removeAlbumFromCollection(albums []models.AlbumFeed, id primitive.ObjectID) []models.AlbumFeed {
	listNew :=[]models.AlbumFeed{}
	for _, album := range albums{
		if album.Id.Hex()!=id.Hex() {
			listNew = append(listNew, album)
		}
	}
	return listNew
}

func removeFromOthersUsersCollection(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.collections.All()
	for _,post := range allPosts {
		for _,singlePost := range post.Posts {

			if singlePost.Post.User.Hex()==idPrimitive.Hex() {
				newListPosts :=removePostFromCollection(post.Posts,singlePost.Id)
				var collectionUpdate = models.Collection{
					Id: post.Id,
					User:post.User,
					Name : post.Name,
					Posts: newListPosts,
				}
				_,_= app.collections.Update(collectionUpdate)

			}
		}
	}
}

func removePostFromCollection(posts []models.FeedPost, id primitive.ObjectID) []models.FeedPost {
	listNew :=[]models.FeedPost{}
	for _, album := range posts{
		if album.Id.Hex()!=id.Hex() {
			listNew = append(listNew, album)
		}
	}
	return listNew
}

func removeFromCommentsAlbum(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.albumFeeds.All()
	for _,post := range allPosts {
		for _,comment := range post.Comments {
			if comment.Writer.Hex()==idPrimitive.Hex() {
				newCommentsList :=removeComment(post.Comments,comment.Writer)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.AlbumFeed{
					Id: post.Id,
					Dislikes:post.Dislikes,
					Comments : newCommentsList,
					Post : postFeed,
					Likes: post.Likes,
				}

				_, _ = app.albumFeeds.Update(feedPostUpdate)
			}
		}
	}
}

func removeFromDislikesAlbum(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.albumFeeds.All()
	for _,post := range allPosts {
		for _,id := range post.Dislikes {
			if id.Hex()==idPrimitive.Hex() {
				newDisLikesList :=removeDislike(post.Dislikes,id)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.AlbumFeed{
					Id: post.Id,
					Dislikes:newDisLikesList,
					Comments : post.Comments,
					Post : postFeed,
					Likes: post.Likes,
				}

				_, _ = app.albumFeeds.Update(feedPostUpdate)
			}
		}
	}
}

func removeFromLikesAlbum(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.albumFeeds.All()
	for _,post := range allPosts {
		for _,id := range post.Likes {
			if id.Hex()==idPrimitive.Hex() {
				newLikesList :=removeLike(post.Likes,id)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.AlbumFeed{
					Id: post.Id,
					Dislikes:post.Dislikes,
					Comments : post.Comments,
					Post : postFeed,
					Likes: newLikesList,
				}

				_, _ = app.albumFeeds.Update(feedPostUpdate)
			}
		}
	}
}

func removeFromComments(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.feedPosts.All()
	for _,post := range allPosts {
		for _,comment := range post.Comments {
			if comment.Writer.Hex()==idPrimitive.Hex() {
				newCommentsList :=removeComment(post.Comments,comment.Writer)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.FeedPost{
					Id: post.Id,
					Dislikes:post.Dislikes,
					Comments : newCommentsList,
					Post : postFeed,
					Likes: post.Likes,
				}

				_, _ = app.feedPosts.Update(feedPostUpdate)
			}
		}
	}
}

func removeComment(comments []models.Comment, writer primitive.ObjectID) (ids []models.Comment) {
	listNew :=[]models.Comment{}
	for _, comment := range comments{
		if comment.Writer.Hex()!=writer.Hex() {
			listNew = append(listNew, comment)
		}
	}
	return listNew
}

func removeFromDislikes(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.feedPosts.All()
	for _,post := range allPosts {
		for _,id := range post.Dislikes {
			if id.Hex()==idPrimitive.Hex() {
				newDisLikesList :=removeDislike(post.Dislikes,id)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.FeedPost{
					Id: post.Id,
					Dislikes:newDisLikesList,
					Comments : post.Comments,
					Post : postFeed,
					Likes: post.Likes,
				}

				_, _ = app.feedPosts.Update(feedPostUpdate)
			}
		}
	}
}

func removeDislike(dislikes []primitive.ObjectID, id primitive.ObjectID) (ids []primitive.ObjectID) {
	listNew :=[]primitive.ObjectID{}
	for _, dislike := range dislikes {
		if dislike.Hex()!=id.Hex() {
			listNew = append(listNew, dislike)
		}
	}
	return listNew
}

func removeFromLikes(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.feedPosts.All()
	for _,post := range allPosts {
		for _,id := range post.Likes {
			if id.Hex()==idPrimitive.Hex() {
				newLikesList :=removeLike(post.Likes,id)
				var postFeed = models.Post{
					User : post.Post.User,
					DateTime : post.Post.DateTime,
					Tagged : post.Post.Tagged,
					Description: post.Post.Description,
					Hashtags: post.Post.Hashtags,
					Location : post.Post.Location,
					Blocked : post.Post.Blocked,
				}
				var feedPostUpdate = models.FeedPost{
					Id: post.Id,
					Dislikes:post.Dislikes,
					Comments : post.Comments,
					Post : postFeed,
					Likes: newLikesList,
				}

				_, _ = app.feedPosts.Update(feedPostUpdate)
			}
		}
	}
}

func removeLike(likes []primitive.ObjectID, id primitive.ObjectID) (ids []primitive.ObjectID) {
	listNew :=[]primitive.ObjectID{}
	for _, like := range likes {
		if like.Hex()!=id.Hex() {
			listNew = append(listNew, like)
		}
	}
	return listNew
}

func removeFromCollectionAlbum(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.collectionAlbums.All()
	for _,post := range allPosts {
		if post.User.Hex()==idPrimitive.Hex() {
			_, _ = app.collectionAlbums.Delete(post.Id.Hex())
		}
	}
}

func removeFromCollection(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.collections.All()
	for _,post := range allPosts {
		if post.User.Hex()==idPrimitive.Hex() {
			_, _ = app.collections.Delete(post.Id.Hex())
		}
	}
}

func removeFromFeedAlbums(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.albumFeeds.All()
	for _,post := range allPosts {
		if post.Post.User.Hex()==idPrimitive.Hex() {
			_, _ = app.albumFeeds.Delete(post.Id.Hex())
		}
	}
}

func removeFromFeedPosts(idPrimitive primitive.ObjectID, app *application) {
	allPosts,_ := app.feedPosts.All()
	for _,post := range allPosts {
		if post.Post.User.Hex()==idPrimitive.Hex() {
			_, _ = app.feedPosts.Delete(post.Id.Hex())
		}
	}
}
