package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"users/pkg/models"
)

// MovieModel represent a mgo database session with a movie data model.
type SettingsModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the movies table.
func (m *SettingsModel) GetAll() ([]models.Settings, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.Settings{}

	// Find all movies
	movieCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = movieCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

// FindByID will be used to find a new movie registry by id
func (m *SettingsModel) FindByID(id primitive.ObjectID) (*models.Settings, error) {

	var user = models.Settings{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

// Insert will be used to insert a new movie registry
func (m *SettingsModel) Insert( notification models.Settings) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), notification)
}

// Delete will be used to delete a movie registry
func (m *SettingsModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *SettingsModel) Update(settings models.Settings)  (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":settings.Id},bson.D{{"$set",bson.M{"user":settings.User,
		"allowTags":settings.AllowTags,
		"acceptMessages":settings.AcceptMessages,"muted":settings.Muted,
		"blocked":settings.Blocked,"closeFriends":settings.CloseFriends}}})
}
