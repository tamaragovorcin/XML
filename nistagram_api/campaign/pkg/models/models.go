package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Ad struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Content uuid.UUID `bson:"content,omitempty"`
	Link string `bson:"link,omitempty"`
}

type Campaign struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Ads []uuid.UUID `bson:"ads,omitempty"`
	TargetGroup []string `bson:"targetGroup,omitempty"`
	Statistic Statistic `bson:"statistic,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
}

type CampaignPost struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Campaign uuid.UUID `bson:"campaign,omitempty"`
	Post uuid.UUID `bson:"post,omitempty"`
}

type CampaignStory struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Campaign uuid.UUID `bson:"campaign,omitempty"`
	Story uuid.UUID `bson:"story,omitempty"`
}

type MultipleTimeCampaign struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Campaign uuid.UUID `bson:"campaign,omitempty"`
	StartTime time.Time `bson:"startTime,omitempty"`
	EndTime time.Time `bson:"endTime,omitempty"`
	DesiredNumber int `bson:"desiredNumber,omitempty"`
	ModifiedTime time.Time `bson:"modifiedTime,omitempty"`
	TimesShown int `bson:"timesShown,omitempty"`
	NumberOfClicks int `bson:"numberOfClicks,omitempty"`
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
	Post uuid.UUID  `bson:"post,omitempty"`

}
