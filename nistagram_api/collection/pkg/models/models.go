package models

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type Collection struct {
	ID  uuid.UUID   `bson:"_id,omitempty"`
	User uuid.UUID  `bson:"user,omitempty"`
	Name string  `bson:"name,omitempty"`
	SavedPosts []uuid.UUID  `bson:"savedPosts,omitempty"`
}



