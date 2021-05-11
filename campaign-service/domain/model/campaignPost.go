package model

import (
	"github.com/google/uuid"
)

type CampaignPost struct {
	Campaign uuid.UUID `gorm:"one2many:campaign_campaignPost;" json:"-"`
	Post uuid.UUID `gorm:"one2many:post_tagged;" json:"-"`
}