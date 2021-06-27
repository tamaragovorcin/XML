package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "html/template"
	"time"
)

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
type User struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	ProfileInformation ProfileInformation `bson:"profileInformation,omitempty"`
	Website            string             `bson:"webSite"`
}
type Role struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`

}
type ProfileInformation struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastName,omitempty"`
	Email string                `validate:"required,email" bson:"email,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Roles string                `bson:"roles,omitempty"`
	PhoneNumber string          `bson:"phoneNumber,omitempty"`
	Gender string               `bson:"gender,omitempty"`
	DateOfBirth string          `bson:"dateOfBirth,omitempty"`
}

type Purchase struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	Buyer primitive.ObjectID `bson:"buyer,omitempty"`
	Products       []CartFrontDTO `bson:"products,omitempty"`
	Location    Location `bson:"location,omitempty"`
}

type Cart struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	ChosenProducts primitive.ObjectID  `bson:"chosenProducts,omitempty"`
	Buyer primitive.ObjectID `bson:"buyer,omitempty"`
	Quantity string `bson:"quantity,omitempty"`
}

type Product struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	User        primitive.ObjectID   `bson:"user"`
	DateTime    time.Time            `bson:"dateTime"`
	Price string `bson:"price,omitempty"`
	Quantity string `bson:"availableQuantity,omitempty"`
	Media []string `bson:"media,omitempty"`
	Name string `bson:"name,omitempty"`
}
type Image struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	Media string `bson:"media"`
	UserId primitive.ObjectID `bson:"userId"`
	PostId primitive.ObjectID `bson:"postId"`
}
type Content struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	User uuid.UUID `bson:"user,omitempty"`
	Media string `bson:"media,omitempty"`
}

type ChosenProduct struct {
	Id  primitive.ObjectID   `bson:"_id,omitempty"`
	Product uuid.UUID `bson:"product,omitempty"`
	Quantity int `bson:"quantity,omitempty"`
}
type Location struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Country string `bson:"country"`
	Town string `bson:"town"`
	Street string `bson:"street"`
	Number int `bson:"number"`
	PostalCode int `bson:"postalCode"`
}


type ProductResponseDTO struct {
	Id          primitive.ObjectID
	User        primitive.ObjectID
	Price string
	Quantity    string
	Name    string
	Media       [][]byte
	DateTime    string
	MediaOrig       []string
}
type CartFrontDTO struct {
	Id          primitive.ObjectID
	Product     ProductResponseDTO
	User        primitive.ObjectID
	Quantity    string
	Media       [][]byte
}

