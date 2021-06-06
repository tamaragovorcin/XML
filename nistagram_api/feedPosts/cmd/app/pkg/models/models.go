package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	User        primitive.ObjectID   `bson:"user"`
	DateTime    time.Time            `bson:"dateTime"`
	Tagged      []primitive.ObjectID `bson:"tagged"`
	Location    Location             `bson:"location"`
	Description string               `bson:"description"`
	Blocked     bool                 `bson:"blocked"`
	Hashtags    []string             `bson:"hashtags"`

}
type Comment struct {
	Id primitive.ObjectID`bson:"_id,omitempty"`
	Content string  `bson:"content"`
	Writer primitive.ObjectID `bson:"writer"`
	DateTime time.Time `bson:"dateTime"`
}

type FeedPost struct {
	Id       primitive.ObjectID   `bson:"_id,omitempty"`
	Post     Post                 `bson:"post"`
	Likes    []primitive.ObjectID `bson:"likes"`
	Dislikes []primitive.ObjectID `bson:"dislikes"`
	Comments []Comment `bson:"comments"`
}

type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country"`
	Town string `bson:"town"`
	Street string `bson:"street"`
	Number int `bson:"number"`
	PostalCode int `bson:"postalCode"`
}

type AlbumFeed struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Post Post `bson:"post"`
	Likes []primitive.ObjectID `bson:"likes"`
	Dislikes []primitive.ObjectID `bson:"dislikes"`
	Comments []Comment `bson:"comments"`
}
type Collection struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	Posts []FeedPost `bson:"posts"`
	Name string `bson:"name"`
}
type SavedPost struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	User primitive.ObjectID  `bson:"user"`
	FeedPost primitive.ObjectID `bson:"feedPost"`
	Collection Collection `bson:"collection"`
}
type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}
type Video struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}