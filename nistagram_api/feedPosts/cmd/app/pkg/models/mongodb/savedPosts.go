package mongodb

import (
	"context"
	"errors"
	"feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SavedPostsModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *SavedPostsModel) All() ([]models.SavedPost, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.SavedPost{}

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
func (m *SavedPostsModel) FindByID(id string) (*models.SavedPost, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find user by id
	var user = models.SavedPost{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

// Insert will be used to insert a new user
func (m *SavedPostsModel) Insert(user models.SavedPost) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *SavedPostsModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})

}
