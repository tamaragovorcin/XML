package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type Purchase struct {
	Id uuid.UUID `gorm:"primaryKey"`
	ChosenProducts []ChosenProduct
}
