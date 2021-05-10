package domain

import (
	"github.com/google/uuid"
	_ "html/template"
)
type Content struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Video string
	Image string
}