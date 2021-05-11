package model

import (
	"github.com/google/uuid"
	_ "html/template"
	"os/user"
)
type AlbumFeed struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Posts []user.User `gorm:"one2many:albumFeed_posts;" json:"-"`
	Likes []user.User `gorm:"many2many:albumFeed_likes;" json:"-"`
	Dislikes []user.User `gorm:"many2many:albumFeed_dislikes;" json:"-"`
	Comments []Comment `gorm:"one2many:albumFeed_comments;" json:"-"`
}
