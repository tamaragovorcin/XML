package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"users/pkg/models"
)
type VerificationModel struct {
	C *mongo.Collection
}

func (m *VerificationModel) GetAll() ([]models.Verification, error) {
	ctx := context.TODO()
	mm := []models.Verification{}

	verificationCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = verificationCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *VerificationModel) FindByID(id string) (*models.Verification, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var verification = models.Verification{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&verification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &verification, nil
}

func (m *VerificationModel) Insert(verification models.Verification) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), verification)
}

func (m *VerificationModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
