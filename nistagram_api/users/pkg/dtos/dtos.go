package dtos

import "time"

type Gender int
const (
	MALE Gender = iota
	FEMALE
)
type UserRequest struct {
	Name  string
	LastName string
	Email string
	Username string
	Password string
	PhoneNumber int
	Gender Gender
	DateOfBirth time.Time
	Private bool
	Biography string
}

type LoginRequest struct {
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
type UserTokenState struct {
	AccessToken string
	ExpiresIn int64
	Roles []string
}



