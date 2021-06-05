package mongodb

import (
"context"
"errors"

"feedPosts/pkg/models"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
)

// UserModel represent a mgo database session with a user model data.
type FeedPostModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *FeedPostModel) All() ([]models.FeedPost, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.FeedPost{}

	// Find all users
	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

// FindByID will be used to find a new user registry by id
func (m *FeedPostModel) FindByID(id primitive.ObjectID) (*models.FeedPost, error) {
	var feedPosts = models.FeedPost{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&feedPosts)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &feedPosts, nil
}

// Insert will be used to insert a new user
func (m *FeedPostModel) Insert(user models.FeedPost) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *FeedPostModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *FeedPostModel) Update(feed models.FeedPost)  (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":feed.Id},bson.D{{"$set",bson.M{"likes":feed.Likes,"dislikes":feed.Dislikes,"comments":feed.Comments,"post.user":feed.Post.User,
		"post.dateTime":feed.Post.DateTime,"post.tagged":feed.Post.Tagged,"post.location":feed.Post.Location,
		"post.description":feed.Post.Description,"post.blocked":feed.Post.Blocked,"post.hashtags":feed.Post.Hashtags}}})
}
