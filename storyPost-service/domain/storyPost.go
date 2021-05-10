package domain

import (
	"github.com/google/uuid"
	_ "html/template"
)
type StoryPost struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Post []uuid.UUID
	CloseFriends bool
}
