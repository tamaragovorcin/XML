package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageDTO struct {
	Id primitive.ObjectID
	DateTime  time.Time
	Text string
	FeedPost primitive.ObjectID
	AlbumPost primitive.ObjectID
	StoryPost primitive.ObjectID
	DisposableImage string
	Deleted bool
	Sender primitive.ObjectID
	Receiver primitive.ObjectID
}
type MessagePostDTO struct {
	Text string
	FeedPost primitive.ObjectID
	AlbumPost primitive.ObjectID
	StoryPost primitive.ObjectID
	DisposableImage primitive.ObjectID
	Deleted bool
	Sender primitive.ObjectID
	Receivers []primitive.ObjectID
}

type UsernameDTO struct {
	Username string

}
type DeletedChatDTO struct {
	Deleted bool
	ForUser primitive.ObjectID
}
type MessageFrontDTO struct {
	Id primitive.ObjectID
	DateTime  time.Time
	Text string
	FeedPost primitive.ObjectID
	AlbumPost primitive.ObjectID
	StoryPost primitive.ObjectID
	DisposableImage string
	OpenedDisposable bool
	DisposableImageId primitive.ObjectID
	Sender string
	Receiver string
}