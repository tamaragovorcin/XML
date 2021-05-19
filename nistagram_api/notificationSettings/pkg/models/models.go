package models

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type NotificationSettings struct {
	AllowTags bool `bson:"allowTags,omitempty"`
	AcceptMessages bool `bson:"acceptMessages,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	NotificationsProfiles []uuid.UUID `bson:"notificationsProfiles,omitempty"`
	NotificationsMessages []uuid.UUID `bson:"notificationsMessages,omitempty"`
	NotificationsPosts []uuid.UUID `bson:"notificationsPosts [,omitempty"`
	NotificationsStories []uuid.UUID `bson:"notificationsStories,omitempty"`
	NotificationsComments []uuid.UUID `bson:"notificationsComments,omitempty"`
}
