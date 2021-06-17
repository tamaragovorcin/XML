package dtos

import (
	"time"
)

type OneTimeCampaignDTO struct {
	User string
	TargetGroup []string
	Link string `bson:"link"`
	Time time.Time
}

type MultipleTimeCampaignDTO struct {
	User string
	TargetGroup []string
	Link string `bson:"link"`
	StartTime time.Time
	EndTime time.Time
	DesiredNumber int `bson:"desiredNumber"`
}