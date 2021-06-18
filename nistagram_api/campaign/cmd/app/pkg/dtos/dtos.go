package dtos

import (
	"campaigns/pkg/models"
	"time"
)

type OneTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	Date time.Time
	Time time.Time
	PartnershipsRequests []string
	Description string
}

type MultipleTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	StartTime time.Time
	EndTime time.Time
	DesiredNumber int
	PartnershipsRequests []string
	Description string
}

