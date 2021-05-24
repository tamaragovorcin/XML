package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Post struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Media string `bson:"media,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	DateTime time.Time `bson:"dateTime,omitempty"`
	Tagged []uuid.UUID `bson:"tagged,omitempty"`
	Location uuid.UUID `bson:"location,omitempty"`
	Description string `bson:"description,omitempty"`
	Blocked bool `bson:"blocked,omitempty"`
	Hashtags []string `bson:"hashtags,omitempty"`

}
type Comment struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Content string  `bson:"content,omitempty"`
	Writer uuid.UUID `bson:"writer,omitempty"`
	DateTime time.Time `bson:"dateTime,omitempty"`
}

type FeedPost struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Post uuid.UUID `bson:"post,omitempty"`
	Likes []uuid.UUID `bson:"likes,omitempty"`
	Dislikes []uuid.UUID `bson:"dislikes,omitempty"`
	Comments []uuid.UUID `bson:"comments,omitempty"`
}

type Location struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Country string `bson:"country,omitempty"`
	Town string `bson:"town,omitempty"`
	Street string `bson:"street,omitempty"`
	Number int `bson:"number,omitempty"`
	PostalCode int `bson:"postalCode,omitempty"`
}

type AlbumFeed struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Posts []string   `bson:"posts,omitempty"`
	Likes []uuid.UUID `bson:"likes,omitempty"`
	Dislikes []uuid.UUID `bson:"dislikes,omitempty"`
	Comments []uuid.UUID `bson:"comments,omitempty"`
}
type Collection struct {
	Id  uuid.UUID   `bson:"_id,omitempty"`
	User uuid.UUID  `bson:"user,omitempty"`
	Name string  `bson:"name,omitempty"`
	SavedPosts []uuid.UUID  `bson:"savedPosts,omitempty"`
}
