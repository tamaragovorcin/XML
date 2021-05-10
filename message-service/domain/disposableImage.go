package domain

import (
	"github.com/google/uuid"
	_ "html/template"
)

type DisposableImage struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Opened bool
	Content string
}
