package mongodb

import (
	"context"
	"errors"

	"feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type CollectionAlbumModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *CollectionAlbumModel) All() ([]models.CollectionAlbum, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.CollectionAlbum{}

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
func (m *CollectionAlbumModel) FindByID(id primitive.ObjectID) (*models.CollectionAlbum, error) {
	var collection = models.CollectionAlbum{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&collection)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &collection, nil
}

// Insert will be used to insert a new user
func (m *CollectionAlbumModel) Insert(user models.CollectionAlbum) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *CollectionAlbumModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
func (m *CollectionAlbumModel) Update(collection models.CollectionAlbum) (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":collection.Id},bson.D{{"$set",bson.M{"name":collection.Name,"user":collection.User,
		"albums":collection.Albums}}})
}
