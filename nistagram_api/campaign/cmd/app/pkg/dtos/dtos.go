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
	Type string
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
	Type string

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
	NumberOfLikes int
	NumberOfDislikes int
	NumberOfComments int
	Likes []LikeDTO
	Dislikes []LikeDTO
	Comments []CommentDTO
	TimesShownTotal int
	BestInfluencer  string
	HiredInfluencers string
	NumberOfClicks int
}
type CampaignAgentAppDTO struct {
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
	NumberOfLikes int
	NumberOfDislikes int
	NumberOfComments int
	Likes string
	Dislikes string
	Comments string
	TimesShownTotal int
	BestInfluencer  string
	HiredInfluencers string
	NumberOfClicks int
	Media []byte
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
type StoryPostInfoHomePageDTO struct {
	Link string
	Type string
	UserId          primitive.ObjectID
	UserUsername    string
	CampaignId primitive.ObjectID
	Stories    []StoryPostInfoDTO
}

type StoryPostInfoDTO struct {
	Id          primitive.ObjectID
	Media       []byte
	Type 		string
	ContentType string
	Link string
}

type CampaignStoriesDTO struct{
	CampaignId primitive.ObjectID
	UserId          primitive.ObjectID
	Media       []byte
	Type 		string
	ContentType string
	Link string
}