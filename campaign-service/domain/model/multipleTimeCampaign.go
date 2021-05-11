package model

import (
	"github.com/google/uuid"
	"time"
)

type MultipleTimeCampaign struct {
	Campaign uuid.UUID `gorm:"many2many:campaign_ads;"`
	StartTime time.Time
	EndTime time.Time
	DesiredNumber int
	ModifiedTime time.Time
	TimesShown int
	NumberOfClicks int
}