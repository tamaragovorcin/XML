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


type Purchase struct {
	Id  uuid.UUID `bson:"_id,omitempty"`
	ChosenProducts []uuid.UUID  `bson:"chosenProducts,omitempty"`
	Buyer uuid.UUID `bson:"buyer,omitempty"`
	Address uuid.UUID `bson:"address,omitempty"`
}
type Product struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Price float64 `bson:"price,omitempty"`
	AvailableQuantity int `bson:"availableQuantity,omitempty"`
	Media []uuid.UUID `bson:"media,omitempty"`
}

type Content struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Media string `bson:"media,omitempty"`
}

type ChosenProduct struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	Product uuid.UUID `bson:"product,omitempty"`
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