package model

import (
	"github.com/google/uuid"
	_ "html/template"
	"os/user"
)
type FeedPost struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Post Post `gorm:"one2one:feedPost_post;" json:"-"`
	Likes []user.User `gorm:"many2many:feedPost_likes;" json:"-"`
	Dislikes []user.User `gorm:"many2many:feedPost_dislikes;" json:"-"`
	Comments []Comment `gorm:"one2many:feedPost_comments;" json:"-"`


}