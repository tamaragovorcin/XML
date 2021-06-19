package dtos

import (
	"campaigns/pkg/models"
)

type OneTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	Date string
	Time string
	PartnershipsRequests []string
	Description string
}
type OneTimeCampaignUpdateDTO struct {
	Id string
	Link string
	Date string
	Time string
	Description string
}

type MultipleTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	StartTime string
	EndTime string
	DesiredNumber int
	PartnershipsRequests []string
	Description string
}

type CampaignDTO struct {
	Id string
	User string
	TargetGroup models.TargetGroup
	Link string
	Date string
	Time string
	Description string
	ContentType string
}