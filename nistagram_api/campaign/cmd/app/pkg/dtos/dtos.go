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

