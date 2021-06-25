package dtos

import (
	"AgentApp/pkg/models"
	"encoding/xml"
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

type PurchaseResponseDTO struct {

	Id          primitive.ObjectID
	Product     []PurchDTO
	User        primitive.ObjectID
	Location  models.Location
}

type PurchDTO struct {
	Id          primitive.ObjectID
	Price string
	Quantity    string
	Name    string
	Media       [][]byte
	MediaOrig       []string
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
type CampaignDTO struct {
	Id string `xml:"id"`
	User string `xml:"user"`
	TargetGroup TargetGroup `xml:"target_group"`
	Link string `xml:"link"`
	Date string `xml:"date"`
	Time string `xml:"time"`
	Description string `xml:"description"`
	ContentType string `xml:"contentType"`
	AgentUsername string  `xml:"agentUsername"`
	AgentId primitive.ObjectID `xml:"agentId"`
	DesiredNumber int `xml:"desiredNumber"`
	CampaignType string `xml:"campaignType"`
	StartTime string `xml:"startTime"`
	EndTime string `xml:"endTime"`
	NumberOfLikes int `xml:"numberOfLikes"`
	NumberOfDislikes int `xml:"numberOfDislikes"`
	NumberOfComments int `xml:"numberOfComments"`
	Likes string `xml:"likes"`
	Dislikes string `xml:"dislikes"`
	Comments string `xml:"comments"`
	TimesShownTotal int `xml:"timesShownTotal"`
	BestInfluencer  string `xml:"bestInfluencer"`
	HiredInfluencers string `xml:"hiredInfluencers"`
}

type TargetGroup struct {
	Gender string `xml:"gender"`
	DateOne string `xml:"dateOne"`
	DateTwo string `xml:"dateTwo"`
	LocationTarget  LocationTarget `xml:"locationTarget"`
}
type LocationTarget struct {
	Country string `xml:"country"`
	Town string `xml:"town"`
	Street string `xml:"street"`
}

type BestCampaigns struct {
	XMLName        xml.Name       `xml:"bestCampaigns"`
	FirstCampaign Campaign `xml:"firstCampaign"`
	SecondCampaign Campaign `xml:"secondCampaign"`
	ThirdCampaign Campaign `xml:"thirdCampaign"`
}
type Campaign struct {
	Id string `xml:"id"`
	User string `xml:"user"`
	TargetGroup TargetGroup `xml:"target_group"`
	Link string `xml:"link"`
	Date string `xml:"date"`
	Time string `xml:"time"`
	Description string `xml:"description"`
	ContentType string `xml:"contentType"`
	AgentUsername string  `xml:"agentUsername"`
	AgentId string `xml:"agentId"`
	DesiredNumber string `xml:"desiredNumber"`
	CampaignType string `xml:"campaignType"`
	StartTime string `xml:"startTime"`
	EndTime string `xml:"endTime"`
	NumberOfLikes string `xml:"numberOfLikes"`
	NumberOfDislikes string `xml:"numberOfDislikes"`
	NumberOfComments string `xml:"numberOfComments"`
	Likes string `xml:"likes"`
	Dislikes string `xml:"dislikes"`
	Comments string `xml:"comments"`
	TimesShownTotal int `xml:"timesShownTotal"`
	BestInfluencer  string `xml:"bestInfluencer"`
	HiredInfluencers string `xml:"hiredInfluencers"`
}

type LikeDTO struct {
	Username string
}

type CommentDTO struct {
	Content string
	Writer string
	DateTime string
}