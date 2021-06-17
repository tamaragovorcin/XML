package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Campaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	TargetGroup []string `bson:"targetGroup"`
	Statistic []primitive.ObjectID `bson:"statistics"`
	Link string `bson:"link"`
	FeedPosts []primitive.ObjectID `bson:"feedPosts"`
	StoryPosts []primitive.ObjectID `bson:"storyPosts"`
}

type MultipleTimeCampaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Campaign Campaign `bson:"campaign"`
	StartTime time.Time `bson:"startTime"`
	EndTime time.Time `bson:"endTime"`
	DesiredNumber int `bson:"desiredNumber"`
	ModifiedTime time.Time `bson:"modifiedTime"`
	TimesShown int `bson:"timesShown"`
}

type OneTimeCampaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Campaign Campaign `bson:"campaign"`
	Time time.Time `bson:"time"`
}

type Statistic struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Influencer primitive.ObjectID `bson:"influencer"`
	NumberOfClicks int  `bson:"numberOfClicks"`
	FeedPost primitive.ObjectID  `bson:"feedPost"`
	StoryPost primitive.ObjectID  `bson:"storyPost"`
}

type Partnership struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Agent primitive.ObjectID `bson:"agent"`
	Influencer  primitive.ObjectID `bson:"influencer"`
	Approved bool `bson:"approved"`
}

type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	CampaignId primitive.ObjectID `bson:"campaignId"`
}
type Video struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	CampaignId primitive.ObjectID `bson:"campaignId"`
}