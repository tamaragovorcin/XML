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