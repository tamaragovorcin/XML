package model

import (
	"github.com/google/uuid"
	_ "html/template"
)
type HighLights struct {
	Id uuid.UUID        `gorm:"primaryKey"`
	Stories []StoryPost `gorm:"many2many:highlights_stories;" json:"-"`
	Title string
}
