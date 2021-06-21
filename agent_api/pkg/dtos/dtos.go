package dtos

import (
	"AgentApp/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Website string
	Role string
}
type ProductDTO struct {
	User        string
	Media       []string
	Price string
	Quantity    string
	Name    string
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


type CartDTO struct {
	Product          primitive.ObjectID
	User        primitive.ObjectID
	Quantity    string

}


type CartFrontDTO struct {
	Id          primitive.ObjectID
	Product     ProductResponseDTO
	User        primitive.ObjectID
	Quantity    string
	Media       [][]byte
}


type OrderFrontDTO struct {
	Id          primitive.ObjectID
	Product     ProductResponseDTO
	User        primitive.ObjectID
	Quantity    string
	Media       [][]byte
	Location  models.Location
}


type PurchaseDTO struct {
	Products       []models.CartFrontDTO
	Location    models.Location
	Buyer primitive.ObjectID
}

type DeleteImageDTO struct {
	AlbumId primitive.ObjectID
	Image string
}
type AddImagesDTO struct {
	PostId primitive.ObjectID
	Media []string
}