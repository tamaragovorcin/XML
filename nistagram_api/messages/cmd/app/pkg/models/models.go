package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Chat struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User1 primitive.ObjectID `bson:"user1"`
	User2 primitive.ObjectID `bson:"user2"`
	Messages []Message `bson:"messages"`
}

type DisposableImage struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Opened bool `bson:"opened"`
	Media string `bson:"media"`
}

type Message struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	DateTime  time.Time `bson:"time"`
	Text string `bson:"text"`
	FeedPost primitive.ObjectID `bson:"feedPost"`
	StoryPost primitive.ObjectID`bson:"storyPost"`
	AlbumPost primitive.ObjectID`bson:"albumPost"`
	DisposableImage primitive.ObjectID `bson:"disposableImage"`
	Deleted bool `bson:"deleted"`
	Sender primitive.ObjectID`bson:"sender"`
}