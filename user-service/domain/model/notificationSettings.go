package model

import "github.com/google/uuid"

type NotificationSettings struct {
	AllowTags bool
	AcceptMessages bool
	User uuid.UUID `gorm:"one2one:user_notifications;" json:"user"`
	NotificationsProfiles []User `gorm:"many2many:user_notifications;" json:"-"`
	NotificationsMessages []User `gorm:"one2one:user_messages;" json:"-"`
	NotificationsPosts []uuid.UUID
	NotificationsStories []uuid.UUID
	NotificationsComments []uuid.UUID
}