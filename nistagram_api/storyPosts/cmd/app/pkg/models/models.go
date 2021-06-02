package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type StoryPost struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Post uuid.UUID `bson:"post,omitempty"`
	OnlyCloseFriends bool `bson:"closeFriends,omitempty"`
}

type HighLight struct {
	Id uuid.UUID        `bson:"_id,omitempty"`
	Stories []uuid.UUID `bson:"stories,omitempty"`
	Name string `bson:"name,omitempty"`
}
type AlbumStory struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Stories []uuid.UUID `bson:"stories,omitempty"`
}
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