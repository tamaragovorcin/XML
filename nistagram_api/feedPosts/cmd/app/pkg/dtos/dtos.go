package dtos

import (
	models2 "feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedPostDTO struct {
	User        string
	Media       string
	Tagged      []primitive.ObjectID
	Location    models2.Location
	Description string
	Hashtags    string
}
type IdDTO struct {
	User primitive.ObjectID
}

type HashTagDTO struct {
	 HashTags string
}

type PostInfoDTO struct {
	DateTime    string
	Tagged      []primitive.ObjectID
	Location    models2.Location
	Description string
	Hashtags    []string
	Media       string
}

type FeedPostInfoDTO struct {
	Id          primitive.ObjectID
	Likes       []primitive.ObjectID
	Dislikes    []primitive.ObjectID
	Comments    []primitive.ObjectID
	DateTime    string
	Tagged      []primitive.ObjectID
	Location    string
	Description string
	Hashtags    string
	Media       []byte
}
type FeedAlbumInfoDTO struct {
	Id          primitive.ObjectID
	Likes       []primitive.ObjectID
	Dislikes    []primitive.ObjectID
	Comments    []primitive.ObjectID
	DateTime    string
	Tagged      []primitive.ObjectID
	Location    string
	Description string
	Hashtags    string
	Media       [][]byte
}