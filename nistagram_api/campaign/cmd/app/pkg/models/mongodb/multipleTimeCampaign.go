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
type MultipleTimeCampaignModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the movies table.
func (m *MultipleTimeCampaignModel) All() ([]models.MultipleTimeCampaign, error) {

	// Define variables
	ctx := context.TODO()
	mm := []models.MultipleTimeCampaign{}

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
func (m *MultipleTimeCampaignModel) FindByID(id string) (*models.MultipleTimeCampaign, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find movie by id
	var movie = models.MultipleTimeCampaign{}
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
func (m *MultipleTimeCampaignModel) Insert(movie models.MultipleTimeCampaign) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), movie)
}

// Delete will be used to delete a movie registry
func (m *MultipleTimeCampaignModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *MultipleTimeCampaignModel) Update(campaign models.MultipleTimeCampaign) (*mongo.UpdateResult, error) {

	return m.C.UpdateOne(context.TODO(),bson.M{"_id":campaign.Id},bson.D{{"$set",bson.M{
		"campaign.link":campaign.Campaign.Link,
		"campaign.description":campaign.Campaign.Description,
		"campaign.user":campaign.Campaign.User,
		"campaign.targetGroup":campaign.Campaign.TargetGroup,
		"campaign.statistics":campaign.Campaign.Statistic,
		"campaign.partnerships":campaign.Campaign.Partnerships,
		"campaign.likes":campaign.Campaign.Likes,
		"campaign.dislikes":campaign.Campaign.Dislikes,
		"campaign.comments":campaign.Campaign.Comments,
		"startTime":campaign.StartTime,
		"endTime":campaign.EndTime,
		"desiredNumber":campaign.DesiredNumber,
		"modifiedTime":campaign.ModifiedTime,
		"timesShown":campaign.TimesShown,
		"campaign.type" : campaign.Campaign.Type,

	}}})
}