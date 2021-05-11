package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type AlbumStory struct {
	Id uuid.UUID         `gorm:"primaryKey"`
	Stories []uuid.UUID    `gorm:"one2many:albumFeed_posts;" json:"-"`
}