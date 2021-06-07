package dtos

import (
	models2 "feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedPostDTO struct {
	User        string
	Media       string
	Tagged      []string
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
	Tagged      string
	Location    models2.Location
	Description string
	Hashtags    []string
	Media       string
}

type FeedPostInfoDTO struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Media       []byte
	Username    string
}
type FeedPostInfoDTO1 struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Username    string
	ContentType string
}
type MediaDTO struct {
	 Type string
	 URL string
}

type CollectionInfoDTO struct {
	Id  primitive.ObjectID
	Posts []FeedPostInfoDTO1
	Name string
}
type CollectionAlbumInfoDTO struct {
	Id  primitive.ObjectID
	Albums []FeedAlbumInfoDTO
	Name string
}
type SavedPostDTO struct {
	User       string
	FeedPost   string
}
type UserCollectionsDTO struct {
	Id         primitive.ObjectID
	Name       string
	SavedPosts []models2.SavedPost
}
type FeedAlbumInfoDTO struct {
	Id          primitive.ObjectID
	DateTime    string
	Tagged      string
	Location    string
	Description string
	Hashtags    string
	Media       [][]byte
	Username    string
}
type CollectionDTO struct {
	Name string
}
type CollectionPostDTO struct {
	PostId primitive.ObjectID
	CollectionId primitive.ObjectID
}

type PostReactionDTO struct {
	PostId primitive.ObjectID
	UserId primitive.ObjectID
}
type CommentReactionDTO struct {
	PostId primitive.ObjectID
	UserId primitive.ObjectID
	Content string
}

type CommentDTO struct {
	Content string
	Writer string
	DateTime string
}
type LikeDTO struct {
	Username string
}