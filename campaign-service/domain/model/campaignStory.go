package model
import (
	"github.com/google/uuid"
)

type CampaignStory struct {
	Campaign uuid.UUID `gorm:"one2many:campaign_campaignPost;" json:"-"`
	Story uuid.UUID `gorm:"one2many:post_tagged;" json:"-"`
}