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
type HighlightModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *HighlightModel) All() ([]models.HighLight, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.HighLight{}

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

func (m *HighlightModel) FindByID(id primitive.ObjectID) (*models.HighLight, error) {

	var highlight = models.HighLight{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&highlight)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &highlight, nil
}

// Insert will be used to insert a new user
func (m *HighlightModel) Insert(user models.HighLight) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *HighlightModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *HighlightModel) Update(highlight models.HighLight) (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":highlight.Id},bson.D{{"$set",bson.M{"name":highlight.Name,"user":highlight.User,
		"stories":highlight.Stories}}})
}
