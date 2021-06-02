package mongodb

import (
	"campaigns/pkg/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PartnershipModel struct {
	C *mongo.Collection
}

func (m *PartnershipModel) GetAll() ([]models.Partnership, error) {
	ctx := context.TODO()
	mm := []models.Partnership{}

	partnershipCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = partnershipCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *PartnershipModel) FindByID(id string) (*models.Partnership, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var partnership = models.Partnership{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&partnership)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &partnership, nil
}

func (m *PartnershipModel) Insert(partnership models.Partnership) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), partnership)
}

func (m *PartnershipModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
