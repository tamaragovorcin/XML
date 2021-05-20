package models
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
	Id uuid.UUID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastname,omitempty"`
	Email string `validate:"required,email" bson:"email,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
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

type Purchase struct {
	Id  uuid.UUID `bson:"_id,omitempty"`
	ChosenProducts []ChosenProduct  `bson:"chosenProducts,omitempty"`
	Buyer uuid.UUID `bson:"buyer,omitempty"`
}
type Product struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Price float64 `bson:"price,omitempty"`
	AvailableQuantity int `bson:"availableQuantity,omitempty"`
	Picture string `bson:"picture,omitempty"`
}

type ChosenProduct struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Product Product `bson:"product,omitempty"`
	Quantity int `bson:"quantity,omitempty"`
}
type Location struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Country string `bson:"country,omitempty"`
	Town string `bson:"town,omitempty"`
	Street string `bson:"street,omitempty"`
	Number int `bson:"number,omitempty"`
	PostalCode int `bson:"postalCode,omitempty"`
}