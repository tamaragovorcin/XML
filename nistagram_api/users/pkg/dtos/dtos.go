package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gender int
const (
	Male Gender = iota
	Female
	Other
)
type UserRequest struct {
	Name  string
	LastName string
	Email string
	Username string
	Password string
	PhoneNumber string
	Gender string
	DateOfBirth string
	Private bool
	Biography string
}
type UserUpdateRequest struct {
	Id string `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`
	LastName string `bson:"lastName,omitempty"`
	Email string `bson:"email,omitempty"`
	Username string `bson:"username,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
	Gender string `bson:"gender,omitempty"`
	DateOfBirth string `bson:"dateOfBirth,omitempty"`
	Private bool `bson:"private,omitempty"`
	Biography string `bson:"biography,omitempty"`
}

type LoginRequest struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
type UserTokenState struct {
	AccessToken string
	ExpiresIn int64
	Roles string
	UserId primitive.ObjectID
}



