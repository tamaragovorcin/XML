package model



import (
	"github.com/google/uuid"
	_ "html/template"
	"os/user"
)
type Comment struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Content string
	Writer user.User `gorm:"many2one:comments_writer;" json:"-"`
}