package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StoryPost struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post `bson:"post,omitempty"`
	OnlyCloseFriends bool `bson:"onlyCloseFriends"`
}

type HighLight struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Stories []StoryPost `bson:"stories"`
	Name string `bson:"name,omitempty"`
}
type AlbumStory struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post
	OnlyCloseFriends bool `bson:"onlyCloseFriends"`
}
type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country,omitempty"`
	Town string `bson:"town,omitempty"`
	Street string `bson:"street,omitempty"`
	Number int `bson:"number,omitempty"`
	PostalCode int `bson:"postalCode,omitempty"`
}

type Post struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	DateTime time.Time `bson:"dateTime,omitempty"`
	Tagged []int `bson:"tagged,omitempty"`
	Location Location `bson:"location,omitempty"`
	Description string `bson:"description,omitempty"`
	Blocked bool `bson:"blocked,omitempty"`
	Hashtags []string `bson:"hashtags,omitempty"`

}

type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media,omitempty"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}