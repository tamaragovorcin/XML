package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Chat struct {
	ID uuid.UUID `bson:"_id,omitempty"`
	Receiver uuid.UUID `bson:"receiver,omitempty"`
	Sender uuid.UUID `bson:"sender,omitempty"`
	Messages []Message `bson:"messages,omitempty"`
}

type DisposableImage struct {
	ID uuid.UUID `bson:"_id,omitempty"`
	Opened bool `bson:"opened,omitempty"`
	Content string `bson:"content,omitempty"`
}

type Message struct {
	ID uuid.UUID `bson:"_id,omitempty"`
	DateTime  time.Time `bson:"time,omitempty"`
	Text string `bson:"text,omitempty"`
	Post uuid.UUID `bson:"post,omitempty"`
	DisposableImage DisposableImage `bson:"disposableImage,omitempty"`
	Deleted bool `bson:"deleted,omitempty"`
}