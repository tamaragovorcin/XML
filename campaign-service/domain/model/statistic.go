package model

import (
	"github.com/google/uuid"
)
type Statistic struct {

	Id uuid.UUID `gorm:"primaryKey"`
	Influencer uuid.UUID
	NumberOfClicks int
	Post uuid.UUID

}