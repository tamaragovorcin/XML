package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"storyPosts/pkg/models"
)

// MovieModel represent a mgo database session with a movie data model.
type HighlightAlbumModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *HighlightAlbumModel) All() ([]models.HighLightAlbum, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.HighLightAlbum{}

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

func (m *HighlightAlbumModel) FindByID(id primitive.ObjectID) (*models.HighLightAlbum, error) {

	var highlight = models.HighLightAlbum{}
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
func (m *HighlightAlbumModel) Insert(user models.HighLightAlbum) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *HighlightAlbumModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *HighlightAlbumModel) Update(highlight models.HighLightAlbum) (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":highlight.Id},bson.D{{"$set",bson.M{"name":highlight.Name,"user":highlight.User,
		"albums":highlight.Albums}}})
}
