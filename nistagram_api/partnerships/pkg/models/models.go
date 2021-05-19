package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Partnership struct {
	ID uuid.UUID `bson:"_id,omitempty"`
	Agent uuid.UUID `bson:"agent,omitempty"`
	Influencer  uuid.UUID `bson:"influencer,omitempty"`
	Approved bool `bson:"approved,omitempty"`
}