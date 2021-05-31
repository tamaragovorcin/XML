package dtos

import (
	"feedPosts/pkg/models"
)

type FeedPostDTO struct {
	User int
	Media string
	Tagged []int
	Location models.Location
	Description string
	Hashtags []string
}