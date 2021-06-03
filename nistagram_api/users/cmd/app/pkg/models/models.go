package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastName,omitempty"`
	Email string                `validate:"required,email" bson:"email,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Roles []Role                `bson:"roles,omitempty"`
	PhoneNumber string          `bson:"phoneNumber,omitempty"`
	Gender string               `bson:"gender,omitempty"`
	DateOfBirth string          `bson:"dateOfBirth,omitempty"`
}


type Role struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`

}
type Agent struct {
	Id  uuid.UUID `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Website string `bson:"website,omitempty"`
	ApprovedAgent bool `bson:"approvedAgent"`
}

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfileInformation ProfileInformation `bson:"profileInformation,omitempty"`
	Private            bool               `bson:"private"`
	Website            string             `bson:"webSite"`
	Biography          string             `bson:"biography"`
	Verified           bool               `bson:"verified"`
	Category           Category           `bson:"category"`
	LikedPosts         []uuid.UUID        `bson:"likedPosts"`
	DislikedPosts      []uuid.UUID        `bson:"disliked"`
}

type Verification struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	User     uuid.UUID `bson:"user"`
	Name     string    `bson:"name"`
	LastName string    `bson:"lastname"`
	Approved bool      `bson:"approved"`
	Document string    `bson:"document"`
	Category Category  `bson:"category"`
}
type ProfileImage struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media,omitempty"`
	UserId primitive.ObjectID `bson:"userId"`
}