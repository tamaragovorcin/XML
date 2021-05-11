package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type Product struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Price float64
	AvailableQuantity int
	Picture string
}
