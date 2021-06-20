package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type TargetGroup struct {
	Gender string
	DateOne string
	DateTwo string
	Location  Location
}
type Campaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user"`
	TargetGroup TargetGroup `bson:"targetGroup"`
	Statistic []Statistic `bson:"statistics"`
	Link string `bson:"link"`
	Description string `bson:"description"`
	Partnerships []Partnership `bson:"partnerships"`
	Likes    []primitive.ObjectID `bson:"likes"`
	Dislikes []primitive.ObjectID `bson:"dislikes"`
	Comments []Comment `bson:"comments"`
}

type MultipleTimeCampaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Campaign Campaign `bson:"campaign"`
	StartTime string `bson:"startTime"`
	EndTime string `bson:"endTime"`
	DesiredNumber int `bson:"desiredNumber"`
	ModifiedTime time.Time `bson:"modifiedTime"`
	TimesShown int `bson:"timesShown"`
}

type OneTimeCampaign struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Campaign Campaign `bson:"campaign"`
	Time string `bson:"time"`
	Date string `bson:"date"`
}

type Statistic struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Influencer primitive.ObjectID `bson:"influencer"`
	NumberOfClicks int  `bson:"numberOfClicks"`
}

type Partnership struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
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

type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country"`
	Town string `bson:"town"`
	Street string `bson:"street"`
	Number int `bson:"number"`
	PostalCode int `bson:"postalCode"`
}
type Comment struct {
	Id primitive.ObjectID`bson:"_id,omitempty"`
	Content string  `bson:"content"`
	Writer primitive.ObjectID `bson:"writer"`
	DateTime time.Time `bson:"dateTime"`
}