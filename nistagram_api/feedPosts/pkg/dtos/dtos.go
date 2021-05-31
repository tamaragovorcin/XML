package dtos

import (
	"feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedPostDTO struct {
	User string
	Media string
	Tagged []int
	Location models.Location
	Description string
	Hashtags string
}
type IdDTO struct {
	User primitive.ObjectID
}