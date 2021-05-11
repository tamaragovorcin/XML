package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type Complaint struct {
	Id uuid.UUID `gorm:"primaryKey"`
	ComplainingUser uuid.UUID
	AccusedUser uuid.UUID
	Post uuid.UUID
}
