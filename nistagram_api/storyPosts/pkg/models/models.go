package models

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"


type StoryPost struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Post uuid.UUID `bson:"_id,omitempty"`
	CloseFriends bool `bson:"_id,omitempty"`
}

type HighLight struct {
	Id uuid.UUID        `bson:"_id,omitempty"`
	Stories []StoryPost `bson:"stories,omitempty"`
	Title string `bson:"title,omitempty"`
}