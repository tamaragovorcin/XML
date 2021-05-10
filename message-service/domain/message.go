package domain


import (
	"github.com/google/uuid"
	_ "html/template"
	"time"
)
type Message struct {
	Id uuid.UUID `gorm:"primaryKey"`
	DateTime  time.Time
	Text string
	Post uuid.UUID
	DisposableImage DisposableImage
	Deleted bool
}
