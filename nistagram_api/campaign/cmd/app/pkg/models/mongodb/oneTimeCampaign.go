package mongodb

import (
	"context"
	"errors"

	"campaigns/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MovieModel represent a mgo database session with a movie data model.
type OneTimeCampaignModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the movies table.
func (m *OneTimeCampaignModel) All() ([]models.OneTimeCampaign, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.OneTimeCampaign{}

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
func (m *OneTimeCampaignModel) FindByID(id string) (*models.OneTimeCampaign, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find movie by id
	var movie = models.OneTimeCampaign{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&movie)
	if err != nil {
		// Checks if the movie was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &movie, nil
}

// Insert will be used to insert a new movie registry
func (m *OneTimeCampaignModel) Insert(movie models.OneTimeCampaign) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), movie)
}

// Delete will be used to delete a movie registry
func (m *OneTimeCampaignModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
func (m *OneTimeCampaignModel) Update(user models.OneTimeCampaign) (*mongo.UpdateResult, error) {

	return m.C.UpdateOne(context.TODO(),bson.M{"_id":user.Id},bson.D{{"$set",bson.M{
		"campaign.link":user.Campaign.Link,
		"campaign.description":user.Campaign.Description,
		"campaign.user":user.Campaign.User,
		"campaign.targetGroup":user.Campaign.TargetGroup,
		"campaign.statistics":user.Campaign.Statistic,
		"campaign.partnerships":user.Campaign.Partnerships,
		"campaign.likes":user.Campaign.Likes,
		"campaign.dislikes":user.Campaign.Dislikes,
		"campaign.comments":user.Campaign.Comments,
		"date":user.Date,
		"time":user.Time,
		}}})
}