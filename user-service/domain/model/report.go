package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type Report struct {
	Id uuid.UUID `gorm:"primaryKey"`
	ComplainingUser uuid.UUID
	ReportedUser uuid.UUID
	Post uuid.UUID
}
