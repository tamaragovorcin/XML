package domain

import (
	"github.com/google/uuid"
	_ "html/template"
)

type FollowRequest struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Following uuid.UUID
	Follower  uuid.UUID
	Accepted  bool
}

