package model

import (
	"github.com/google/uuid"
	_ "html/template"
)
type Verification struct {
	Id uuid.UUID `gorm:"primaryKey"`
	User uuid.UUID
	Approved bool
	Document string
	Category Category
}
