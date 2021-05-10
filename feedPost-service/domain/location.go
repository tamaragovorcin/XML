package domain

import (
	"github.com/google/uuid"
	_ "html/template"
)
type Location struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Country string
	Town string
	Street string
	Number int
	PostalCode int
}
