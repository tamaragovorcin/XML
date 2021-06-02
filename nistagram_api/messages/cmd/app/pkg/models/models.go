package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Chat struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	User1 uuid.UUID `bson:"user1,omitempty"`
	User2 uuid.UUID `bson:"user2,omitempty"`
	Messages []uuid.UUID `bson:"messages,omitempty"`
}

type DisposableImage struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Opened bool `bson:"opened,omitempty"`
	Media string `bson:"media,omitempty"`
}

type Message struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	DateTime  time.Time `bson:"time,omitempty"`
	Text string `bson:"text,omitempty"`
	FeedPost uuid.UUID `bson:"feedPost,omitempty"`
	StoryPost uuid.UUID `bson:"storyPost,omitempty"`
	DisposableImage uuid.UUID `bson:"disposableImage,omitempty"`
	Deleted bool `bson:"deleted,omitempty"`
	Sender uuid.UUID `bson:"sender,omitempty"`
}