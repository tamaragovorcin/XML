package domain

import (
	"github.com/google/uuid"
	_ "html/template"
	"os/user"
	"time"
)
type Post struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Content Content
	User user.User `gorm:"one2one:post_user;" json:"-"`
	DateTime time.Time
	Tagged []user.User `gorm:"one2many:post_tagged;" json:"-"`
	Location Location
	Description string
	Blocked bool
}
