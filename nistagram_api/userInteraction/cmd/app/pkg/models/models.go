package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type FollowRequest struct {
	Id uuid.UUID  `bson:"_id,omitempty"`
	Following  uuid.UUID `bson:"following,omitempty"`
	Follower   uuid.UUID `bson:"follower,omitempty"`
	Approved  bool `bson:"approved,omitempty"`
	DateTime time.Time `bson:"dateTime,omitempty"`
}

type Report struct {
	Id uuid.UUID `bson:"_id,omitempty"`
	ComplainingUser uuid.UUID `bson:"complainingUser,omitempty"`
	ReportedUser uuid.UUID `bson:"reportedUser,omitempty"`
	FeedPost uuid.UUID `bson:"feedPost,omitempty"`
	StoryPost uuid.UUID `bson:"storyPost,omitempty"`
}
type MovieResult struct {
	Movie `json:"movie"`
}

type VoteResult struct {
	Updates int `json:"updates"`
}

type Movie struct {
	Released int64    `json:"released"`
	Title    string   `json:"title,omitempty"`
	Tagline  string   `json:"tagline,omitempty"`
	Votes    int64    `json:"votes,omitempty"`
	Cast     []Person `json:"cast,omitempty"`
}

type Person struct {
	Job  string   `json:"job"`
	Role []string `json:"role"`
	Name string   `json:"name"`
}

type D3Response struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

type Node struct {
	Title string `json:"title"`
	Label string `json:"label"`
}

type Link struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}
