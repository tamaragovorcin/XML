package models

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type Notifications struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	User     uuid.UUID `bson:"user,omitempty"`
	NotificationsProfiles []uuid.UUID `bson:"notificationsProfiles,omitempty"`
	NotificationsMessages []uuid.UUID `bson:"notificationsMessages,omitempty"`
	NotificationsPosts []uuid.UUID `bson:"notificationsPosts [,omitempty"`
	NotificationsStories []uuid.UUID `bson:"notificationsStories,omitempty"`
	NotificationsComments []uuid.UUID `bson:"notificationsComments,omitempty"`
}

type Settings struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	User     uuid.UUID `bson:"user,omitempty"`
	AllowTags bool `bson:"allowTags,omitempty"`
	AcceptMessages bool `bson:"acceptMessages,omitempty"`
	Muted []uuid.UUID `bson:"muted,omitempty"`
	Blocked []uuid.UUID `bson:"blocked,omitempty"`
	CloseFriends []uuid.UUID `bson:"closeFriends,omitempty"`
}


