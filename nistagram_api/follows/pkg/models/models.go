package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type FollowRequest struct {
	Id uuid.UUID  `bson:"_id,omitempty"`
	Following  uuid.UUID `bson:"following,omitempty"`
	Follower   uuid.UUID `bson:"follower,omitempty"`
	Accepted  bool `bson:"accepted,omitempty"`

}
