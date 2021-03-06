package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"storyPosts/pkg/models"
)

type StoryPostDTO struct {
	User string
	Media string
	Tagged []string
	Location models.Location
	Description string
	Hashtags string
	OnlyCloseFriends bool
}
type IdDTO struct {
	User primitive.ObjectID
}
type StoryPostInfoDTO struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Media       []byte
	Type 		string
	ContentType string
}
type StoryPostInfoDTO1 struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Type 		string
	ContentType string
}
type HighlightDTO struct {
	Name string
}
type HighlightStoryDTO struct {
	StoryId primitive.ObjectID
	HighlightId primitive.ObjectID
}
type HighlightStoryAlbumDTO struct {
	StoryId primitive.ObjectID
	HighlightId primitive.ObjectID
}
type HighlightsInfoDTO struct {
	Id  primitive.ObjectID
	Stories []StoryPostInfoDTO
	Name string
}
type HighlightsAlbumInfoDTO struct {
	Id  primitive.ObjectID
	Albums []StoryAlbumInfoDTO
	Name string
}

type StoryPostInfoHomePageDTO struct {
	UserId          primitive.ObjectID
	UserUsername    string
	Stories    []StoryPostInfoDTO
	CloseFriends bool
	Id primitive.ObjectID
}
type ResponseDTO struct {
	Message string
	Username string
}
type StoryAlbumInfoDTO struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Media       [][]byte
	Username    string
}
type StoryAlbumInfoHomePageDTO struct {
	UserId          primitive.ObjectID
	UserUsername    string
	Albums    []StoryAlbumInfoDTO
}
