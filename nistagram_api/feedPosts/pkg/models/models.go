package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"os/user"
	"time"
)

type Post struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Content Content `bson:"content,omitempty"`
	User user.User `bson:"user,omitempty"`
	DateTime time.Time `bson:"dateTime,omitempty"`
	Tagged []user.User `bson:"tagged,omitempty"`
	Location Location `bson:"location,omitempty"`
	Description string `bson:"description,omitempty"`
	Blocked bool `bson:"blocked,omitempty"`
}
type Comment struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Content string  `bson:"content,omitempty"`
	Writer user.User `bson:"writer,omitempty"`
}

type Content struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Video string `bson:"video,omitempty"`
	Image string `bson:"image,omitempty"`
}

type FeedPost struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Post Post `bson:"post,omitempty"`
	Likes []user.User `bson:"likes,omitempty"`
	Dislikes []user.User `bson:"dislikes,omitempty"`
	Comments []Comment `bson:"comments,omitempty"`
}

type Location struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Country string `bson:"country,omitempty"`
	Town string `bson:"town,omitempty"`
	Street string `bson:"street,omitempty"`
	Number int `bson:"number,omitempty"`
	PostalCode int `bson:"postalCode,omitempty"`
}
