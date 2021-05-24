package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type FollowRequest struct {
	Id uuid.UUID  `bson:"_id,omitempty"`
	Following  uuid.UUID `bson:"following,omitempty"`
	Follower   uuid.UUID `bson:"follower,omitempty"`
	Approved  bool `bson:"approved,omitempty"`
	DateTime time.Time `bson:"dateTime,omitempty"`
}

type Report struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	ComplainingUser uuid.UUID `bson:"complainingUser,omitempty"`
	ReportedUser uuid.UUID `bson:"reportedUser,omitempty"`
	FeedPost uuid.UUID `bson:"feedPost,omitempty"`
	StoryPost uuid.UUID `bson:"storyPost,omitempty"`
}