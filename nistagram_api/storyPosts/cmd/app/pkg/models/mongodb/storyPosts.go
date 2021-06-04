package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"storyPosts/pkg/models"
)

// UserModel represent a mgo database session with a user model data.
type StoryPostModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *StoryPostModel) All() ([]models.StoryPost, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.StoryPost{}

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

func (m *StoryPostModel) FindByID(id primitive.ObjectID) (*models.StoryPost, error) {
	var story = models.StoryPost{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&story)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &story, nil
}

// Insert will be used to insert a new user
func (m *StoryPostModel) Insert(user models.StoryPost) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *StoryPostModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
