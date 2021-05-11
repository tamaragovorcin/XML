package model

import (
	"github.com/google/uuid"
	_ "html/template"
	"time"
)

type Gender int

const (

	MALE Gender = iota
	FEMALE
)

type User struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Name  string
	Email string `validate:"required,email"`
	Username string `gorm:"unique;not null"`
	Password string
	Surname string
	Roles []Role `gorm:"many2many:user_roles;"`
	PhoneNumber int
	Gender Gender
	DateOfBirth time.Time
}

type UserRequest struct {
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ActivateLinkRequest struct {
	Email string `json:"email"`
}

type ChangeNewPasswordRequest struct {
	ResetPasswordId string `json:"resetPasswordId"`
	Password string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}


