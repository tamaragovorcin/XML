package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"storyPosts/pkg/models"
)

type StoryPostDTO struct {
	User string
	Media string
	Tagged []int
	Location models.Location
	Description string
	Hashtags string
	OnlyCloseFriends bool
}
type IdDTO struct {
	User primitive.ObjectID
}
