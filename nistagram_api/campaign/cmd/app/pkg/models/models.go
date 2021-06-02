package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Campaign struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	TargetGroup []string `bson:"targetGroup,omitempty"`
	Statistic []uuid.UUID `bson:"statistics,omitempty"`
	Link string `bson:"link,omitempty"`
	FeedPosts []uuid.UUID `bson:"feedPosts,omitempty"`
	StoryPosts []uuid.UUID `bson:"storyPosts,omitempty"`
}

type MultipleTimeCampaign struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Campaign uuid.UUID `bson:"campaign,omitempty"`
	StartTime time.Time `bson:"startTime,omitempty"`
	EndTime time.Time `bson:"endTime,omitempty"`
	DesiredNumber int `bson:"desiredNumber,omitempty"`
	ModifiedTime time.Time `bson:"modifiedTime,omitempty"`
	TimesShown int `bson:"timesShown,omitempty"`
}

type OneTimeCampaign struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Campaign uuid.UUID `bson:"campaign,omitempty"`
	Time time.Time `bson:"time,omitempty"`
}

type Statistic struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Influencer uuid.UUID `bson:"influencer,omitempty"`
	NumberOfClicks int  `bson:"numberOfClicks,omitempty"`
	FeedPost uuid.UUID  `bson:"feedPost,omitempty"`
	StoryPost uuid.UUID  `bson:"storyPost,omitempty"`
}

type Partnership struct {
	ID uuid.UUID `bson:"_id,omitempty"`
	Agent uuid.UUID `bson:"agent,omitempty"`
	Influencer  uuid.UUID `bson:"influencer,omitempty"`
	Approved bool `bson:"approved,omitempty"`
}