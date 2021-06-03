package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Post struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	User        primitive.ObjectID   `bson:"user"`
	DateTime    time.Time            `bson:"dateTime,omitempty"`
	Tagged      []primitive.ObjectID `bson:"tagged,omitempty"`
	Location    Location             `bson:"location,omitempty"`
	Description string               `bson:"description,omitempty"`
	Blocked     bool                 `bson:"blocked"`
	Hashtags    []string             `bson:"hashtags"`

}
type Comment struct {
	Id primitive.ObjectID`bson:"_id,omitempty"`
	Content string  `bson:"content,omitempty"`
	Writer primitive.ObjectID `bson:"writer"`
	DateTime time.Time `bson:"dateTime,omitempty"`
}

type FeedPost struct {
	Id       primitive.ObjectID   `bson:"_id,omitempty"`
	Post     Post                 `bson:"post,omitempty"`
	Likes    []primitive.ObjectID `bson:"likes"`
	Dislikes []primitive.ObjectID `bson:"dislikes"`
	Comments []primitive.ObjectID `bson:"comments"`
}

type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country,omitempty"`
	Town string `bson:"town,omitempty"`
	Street string `bson:"street,omitempty"`
	Number int `bson:"number,omitempty"`
	PostalCode int `bson:"postalCode,omitempty"`
}

type AlbumFeed struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post `bson:"post,omitempty"`
	Likes []primitive.ObjectID `bson:"likes"`
	Dislikes []primitive.ObjectID `bson:"dislikes"`
	Comments []primitive.ObjectID `bson:"comments"`
}
type Collection struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	User primitive.ObjectID  `bson:"user"`
	Name string  `bson:"name,omitempty"`
	SavedPosts []uuid.UUID  `bson:"savedPosts,omitempty"`
}
type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media,omitempty"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}
