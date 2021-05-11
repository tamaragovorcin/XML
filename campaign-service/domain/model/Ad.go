package model

import (
	"github.com/google/uuid"
)

type Ad struct {
	Content uuid.UUID
	Link string
}