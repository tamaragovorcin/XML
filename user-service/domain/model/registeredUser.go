package model

import (
	"github.com/google/uuid"
	_ "html/template"
)
type Category int

const (

	INFLUENCER Category = iota
	SPORTS
	NEW_MEDIA
	BUSINESS
	BREND
	ORGANIZATION
)
type RegisteredUser struct {
	User User
	Private bool
	Website string
	Biography string
	Verified bool
	Category Category
	Followers []User `gorm:"many2many:user_followers;" json:"followers"`
	Following []User `gorm:"many2many:user_following;" json:"following"`
	FavoritePosts []uuid.UUID `gorm:"many2many:user_favoritePosts;" json:"favoritePosts"`                                    //post
	Collections []uuid.UUID `gorm:"many2many:user_collections;" json:"collections"`
	Posts []uuid.UUID `gorm:"many2many:user_posts;" json:"posts"`
	Stories []uuid.UUID `gorm:"many2many:user_stories;" json:"stories"`
	CloseFriends []uuid.UUID `gorm:"many2many:user_closeFriends;" json:"closeFriends"`
	Highlights []uuid.UUID `gorm:"many2many:user_highlights;" json:"highlights"`
	Muted []uuid.UUID `gorm:"many2many:user_muted;" json:"muted"`
	Blocked []uuid.UUID `gorm:"many2many:user_blocked;" json:"blocked"`
	NotificationSettings uuid.UUID `gorm:"one2one:user_notifications;" json:"notifications"`            //?????????????????????????????????
	LikedPosts []uuid.UUID `gorm:"many2many:user_liked;" json:"liked"`
	Disliked []uuid.UUID `gorm:"many2many:user_disliked;" json:"disliked"`
}
