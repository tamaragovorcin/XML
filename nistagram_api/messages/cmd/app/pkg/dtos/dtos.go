package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageDTO struct {
	Id primitive.ObjectID
	DateTime  time.Time
	Text string
	FeedPost primitive.ObjectID
	StoryPost primitive.ObjectID
	DisposableImage primitive.ObjectID
	Deleted bool
	Sender primitive.ObjectID
	Receiver primitive.ObjectID
}