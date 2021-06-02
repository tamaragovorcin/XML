package mongodb

import (
	"context"
	"errors"
	"gomod/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DisposableImageModel struct {
	C *mongo.Collection
}

func (m *DisposableImageModel) GetAll() ([]models.DisposableImage, error) {
	ctx := context.TODO()
	mm := []models.DisposableImage{}

	disposableImageCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = disposableImageCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *DisposableImageModel) FindByID(id string) (*models.DisposableImage, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var disposableImage = models.DisposableImage{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&disposableImage)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &disposableImage, nil
}

func (m *DisposableImageModel) Insert(disposableImage models.DisposableImage) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), disposableImage)
}

func (m *DisposableImageModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
