package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gomod/pkg/models"
)

type MessageModel struct {
	C *mongo.Collection
}

func (m *MessageModel) GetAll() ([]models.Message, error) {
	ctx := context.TODO()
	mm := []models.Message{}

	messageCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = messageCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *MessageModel) FindByID(id string) (*models.Message, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var message = models.Message{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &message, nil
}

func (m *MessageModel) Insert(message models.Message) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), message)
}

func (m *MessageModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
