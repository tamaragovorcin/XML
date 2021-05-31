package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)
//ENUMS
type Gender int
const (
	Male Gender = iota
	Female
	Other
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
type ProfileInformation struct {
	Id       int `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastName,omitempty"`
	Email string `validate:"required,email" bson:"email,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Roles []Role `bson:"roles,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
	Gender string `bson:"gender,omitempty"`
	DateOfBirth string `bson:"dateOfBirth,omitempty"`
}


type Role struct {
	Id int `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`

}
type Agent struct {
	Id  uuid.UUID `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Website string `bson:"website,omitempty"`
	ApprovedAgent bool `bson:"approvedAgent,omitempty"`
}

type User struct {
	Id int `bson:"_id,omitempty"`
	ProfileInformation ProfileInformation `bson:"profileInformation,omitempty"`
	Private bool `bson:"private,omitempty"`
	Website string `bson:"webSite,omitempty"`
	Biography string `bson:"biography,omitempty"`
	Verified bool `bson:"verified,omitempty"`
	Category Category `bson:"category,omitempty"`
	LikedPosts []uuid.UUID `bson:"likedPosts,omitempty"`
	DislikedPosts []uuid.UUID `bson:"disliked,omitempty"`
}

type Verification struct {
	Id uuid.UUID  `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastname,omitempty"`
	Approved bool `bson:"approved,omitempty"`
	Document string `bson:"document,omitempty"`
	Category Category `bson:"category,omitempty"`
}
