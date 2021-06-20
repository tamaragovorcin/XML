package dtos

import (
	"campaigns/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OneTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	Date string
	Time string
	PartnershipsRequests []string
	Description string
}
type OneTimeCampaignUpdateDTO struct {
	Id string
	Link string
	Date string
	Time string
	Description string
	User string
}
type MultipleTimeCampaignUpdateDTO struct {
	Id string
	Link string
	Date string
	StartTime string
	EndTime string
	DesiredNumber string
	Description string
	User string
}

type MultipleTimeCampaignDTO struct {
	User string
	TargetGroup models.TargetGroup
	Link string
	StartTime string
	EndTime string
	DesiredNumber string
	PartnershipsRequests []string
	Description string
}

type CampaignDTO struct {
	Id string
	User string
	TargetGroup models.TargetGroup
	Link string
	Date string
	Time string
	Description string
	ContentType string
	AgentUsername string
	AgentId primitive.ObjectID
	DesiredNumber int
	CampaignType string
	StartTime string
	EndTime string
}

type CampaignMultipleDTO struct {
	Id string
	User string
	TargetGroup models.TargetGroup
	Link string
	StartTime string
	EndTime string
	DesiredNumber string
	Description string
	ContentType string
	AgentUsername string
	AgentId primitive.ObjectID
}

type PartnershipDTO struct {
	CampaignId primitive.ObjectID
	UserId primitive.ObjectID
}

type CampaignReactionDTO struct {
	CampaignId primitive.ObjectID
	UserId primitive.ObjectID
	Content string
}
type LikeDTO struct {
	Username string
}

type CommentDTO struct {
	Content string
	Writer string
	DateTime string
}

type BestStatisticsDTO struct {
	UserId primitive.ObjectID
	NumberOfClicks int
	NumberOfPartnerships int
	Username string
}
