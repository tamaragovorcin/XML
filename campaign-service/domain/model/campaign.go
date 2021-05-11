package model

import (
	"github.com/google/uuid"
)

type Campaign struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Ads []uuid.UUID `gorm:"many2many:campaign_ads;"`
	TargetGroup []string
	Statistic Statistic
	User uuid.UUID `gorm:"many2many:campaign_ads;"`
}