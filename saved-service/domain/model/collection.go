package model

import (
	"github.com/google/uuid"
	"os/user"
)

type Collection struct {
	Id uuid.UUID `gorm:"primaryKey"`
	User user.User `gorm:"one2one:post_user;" json:"-"`
	Name string
	SavedPosts []uuid.UUID `gorm:"many2many:collection_savedposts;"`
}