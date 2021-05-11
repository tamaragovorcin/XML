package model

import (
	"time"
	"github.com/google/uuid"
)

type OneTimeCampaign struct {
	Campaign uuid.UUID `gorm:"many2many:campaign_ads;"`
	Time time.Time
}