package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Name     string             `bson:"name"`
	LastName string             `bson:"lastName"`
	Email string                `validate:"required,email" bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Roles []Role                `bson:"roles"`
	PhoneNumber string          `bson:"phoneNumber"`
	Gender string               `bson:"gender"`
	DateOfBirth string          `bson:"dateOfBirth"`
}


type Role struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`

}
type Agent struct {
	Id  primitive.ObjectID  `bson:"_id,omitempty"`
	ProfileInformation ProfileInformation `bson:"profileInformation,omitempty"`
	Private            bool               `bson:"private"`
	Website            string             `bson:"webSite"`
	Biography          string             `bson:"biography"`
	Verified           bool               `bson:"verified"`
	Category           Category           `bson:"category"`
	ApprovedAgent bool `bson:"approvedAgent"`
}

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfileInformation ProfileInformation `bson:"profileInformation"`
	Private            bool               `bson:"private"`
	Website            string             `bson:"webSite"`
	Biography          string             `bson:"biography"`
	Verified           bool               `bson:"verified"`

	ApprovedAgent string `bson:"approvedAgent"`

	Category           string           `bson:"category"`

}

type Verification struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	User     primitive.ObjectID `bson:"user"`
	Name     string    `bson:"name"`
	LastName string    `bson:"lastname"`
	Approved bool      `bson:"approved"`

	Category string  `bson:"category"`
}
type ProfileImage struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
}
type Notifications struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	User     primitive.ObjectID `bson:"user"`
	NotificationsMessages bool `bson:"notificationsMessages"`
	NotificationsComments bool `bson:"notificationsComments"`
}
type NotificationForUser struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Subject     primitive.ObjectID `bson:"subject"`
	Object     primitive.ObjectID `bson:"object"`
	Posts bool `bson:"posts"`
	Stories bool `bson:"stories"`
}
type NotificationContent struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Subject     primitive.ObjectID `bson:"subject"`
	Object primitive.ObjectID `bson:"object"`
	Posted string `bson:"posted"`
}
type Settings struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	User     primitive.ObjectID `bson:"user"`
	AllowTags bool `bson:"allowTags"`
	AcceptMessages bool `bson:"acceptMessages"`
	Muted []primitive.ObjectID `bson:"muted"`
	Blocked []primitive.ObjectID `bson:"blocked"`
	CloseFriends []primitive.ObjectID  `bson:"closeFriends"`
}
type Image struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	VerificationId primitive.ObjectID `bson:"verificationId"`
}