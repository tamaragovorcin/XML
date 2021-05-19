package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type AlbumFeed struct {
	Id uuid.UUID         `gorm:"primaryKey"`
	Posts []uuid.UUID    `gorm:"one2many:albumFeed_posts;" json:"-"`
	Likes []uuid.UUID    `gorm:"many2many:albumFeed_likes;" json:"-"`
	Dislikes []uuid.UUID `gorm:"many2many:albumFeed_dislikes;" json:"-"`
	Comments []uuid.UUID  `gorm:"one2many:albumFeed_comments;" json:"-"`

}
type AlbumStory struct {
	Id uuid.UUID         `gorm:"primaryKey"`
	Stories []uuid.UUID    `gorm:"one2many:albumFeed_posts;" json:"-"`
}
