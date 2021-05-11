package model

import (
	"github.com/google/uuid"
	_ "html/template"
)
type Chat struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Receiver uuid.UUID
	Sender uuid.UUID
	Messages []Message `gorm:"one2many:messages;"`
}
