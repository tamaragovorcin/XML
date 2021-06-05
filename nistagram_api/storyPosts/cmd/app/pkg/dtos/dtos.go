package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"storyPosts/pkg/models"
)

type StoryPostDTO struct {
	User string
	Media string
	Tagged []primitive.ObjectID
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
	Tagged      []primitive.ObjectID
	Location    string
	Description string
	Hashtags    string
	Media       []byte
}
type HighlightDTO struct {
	Name string
}
type HighlightStoryDTO struct {
	StoryId primitive.ObjectID
	HighlightId primitive.ObjectID
}

type HighlightsInfoDTO struct {
	Id  primitive.ObjectID
	Stories []StoryPostInfoDTO
	Name string
}


