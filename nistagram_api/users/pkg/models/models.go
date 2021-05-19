package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)
//ENUMS
type Gender int
const (
	MALE Gender = iota
	FEMALE
)

type Category int
const (
	INFLUENCER Category = iota
	SPORTS
	NEW_MEDIA
	BUSINESS
	BRAND
	ORGANIZATION
)

// User is used to represent user profile data
type User struct {
	ID       uuid.UUID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastname,omitempty"`
	Email string `validate:"required,email" bson:"email,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Roles []Role `bson:"role,omitempty"`
	PhoneNumber int `bson:"phoneNumber,omitempty"`
	Gender Gender `bson:"gender,omitempty"`
	DateOfBirth time.Time `bson:"dateOfBirth,omitempty"`
}

type UserRequest struct {
	Name  string `bson:"name,omitempty"`
	Surname string `bson:"surname,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	RepeatedPassword string `bson:"repeatedPassword,omitempty"`
}

type LoginRequest struct {
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type Role struct {
	Id string `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`

}
type Agent struct {
	User User `bson:"user,omitempty"`
	Website string `bson:"website,omitempty"`
	Approved bool `bson:"approved,omitempty"`
}

type RegisteredUser struct {
	User User `bson:"user,omitempty"`
	Private bool `bson:"private,omitempty"`
	Website string `bson:"approved,omitempty"`
	Biography string `bson:"biography,omitempty"`
	Verified bool `bson:"verified,omitempty"`
	Category Category `bson:"category,omitempty"`
	Followers []User `bson:"followers,omitempty"`
	Following []User `bson:"following,omitempty"`
	FavoritePosts []uuid.UUID `bson:"favoritePosts,omitempty"`
	Collections []uuid.UUID `bson:"collections,omitempty"`
	Posts []uuid.UUID `bson:"posts,omitempty"`
	Stories []uuid.UUID `bson:"stories,omitempty"`
	CloseFriends []uuid.UUID `bson:"closeFriends,omitempty"`
	Highlights []uuid.UUID `bson:"highlights,omitempty"`
	Muted []uuid.UUID `bson:"muted,omitempty"`
	Blocked []uuid.UUID `bson:"blocked,omitempty"`
	NotificationSettings uuid.UUID `bson:"notificationSettings,omitempty"`
	LikedPosts []uuid.UUID `bson:"likedPosts,omitempty"`
	Disliked []uuid.UUID `bson:"disliked,omitempty"`
}
type Report struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	ComplainingUser uuid.UUID `bson:"complainingUser,omitempty"`
	ReportedUser uuid.UUID `bson:"reportedUser,omitempty"`
	Post uuid.UUID `bson:"post,omitempty"`
}

type Verification struct {
	Id uuid.UUID  `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Approved bool `bson:"approved,omitempty"`
	Document string `bson:"document,omitempty"`
	Category Category `bson:"category,omitempty"`
}
