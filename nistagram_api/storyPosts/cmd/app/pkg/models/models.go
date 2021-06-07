package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StoryPost struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post `bson:"post"`
	OnlyCloseFriends bool `bson:"onlyCloseFriends"`
}

type HighLight struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	Stories []StoryPost `bson:"stories"`
	Name string `bson:"name"`
}
type AlbumStory struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post
	OnlyCloseFriends bool `bson:"onlyCloseFriends"`
}
type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country"`
	Town string `bson:"town"`
	Street string `bson:"street"`
	Number int `bson:"number"`
	PostalCode int `bson:"postalCode"`
}

type Post struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	DateTime time.Time `bson:"dateTime"`
	Tagged []primitive.ObjectID `bson:"tagged"`
	Location Location `bson:"location"`
	Description string `bson:"description"`
	Blocked bool `bson:"blocked"`
	Hashtags []string `bson:"hashtags"`

}

type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}
type HighLightAlbum struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	Albums []AlbumStory `bson:"albums"`
	Name string `bson:"name"`
}