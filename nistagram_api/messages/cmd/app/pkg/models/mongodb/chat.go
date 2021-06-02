package mongodb

import (
	"context"
	"errors"
	"gomod/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatModel struct {
	C *mongo.Collection
}

func (m *ChatModel) GetAll() ([]models.Chat, error) {
	ctx := context.TODO()
	mm := []models.Chat{}

	chatCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = chatCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *ChatModel) FindByID(id string) (*models.Chat, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var chat = models.Chat{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &chat, nil
}

func (m *ChatModel) Insert(chat models.Chat) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), chat)
}

func (m *ChatModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
