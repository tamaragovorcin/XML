package model
import (
	"github.com/google/uuid"
	_ "html/template"
)
type AlbumFeed struct {
	Id uuid.UUID `gorm:"primaryKey"`
	AgentId uuid.UUID
	InfluencerId uuid.UUID
	Approved bool
}