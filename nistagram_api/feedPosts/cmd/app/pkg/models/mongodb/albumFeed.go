package mongodb

import (
	"context"
	"errors"

	"feedPosts/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MovieModel represent a mgo database session with a movie data model.
type AlbumFeedModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the movies table.
func (m *AlbumFeedModel) All() ([]models.AlbumFeed, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.AlbumFeed{}

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
func (m *AlbumFeedModel) FindByID(id primitive.ObjectID) (*models.AlbumFeed, error) {
	var feedAlbum = models.AlbumFeed{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&feedAlbum)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &feedAlbum, nil
}

// Insert will be used to insert a new movie registry
func (m *AlbumFeedModel) Insert(movie models.AlbumFeed) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), movie)
}

// Delete will be used to delete a movie registry
func (m *AlbumFeedModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *AlbumFeedModel) Update(album models.AlbumFeed) (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(context.TODO(),bson.M{"_id":album.Id},bson.D{{"$set",bson.M{"likes":album.Likes,"dislikes":album.Dislikes,"comments":album.Comments,"post.user":album.Post.User,
		"post.dateTime":album.Post.DateTime,"post.tagged":album.Post.Tagged,"post.location":album.Post.Location,
		"post.description":album.Post.Description,"post.blocked":album.Post.Blocked,"post.hashtags":album.Post.Hashtags}}})
}
