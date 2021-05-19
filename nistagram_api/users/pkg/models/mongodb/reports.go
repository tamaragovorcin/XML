package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"users/pkg/models"
)
type ReportModel struct {
	C *mongo.Collection
}

func (m *ReportModel) GetAll() ([]models.Report, error) {
	ctx := context.TODO()
	mm := []models.Report{}

	reportCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = reportCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *ReportModel) FindByID(id string) (*models.Report, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report = models.Report{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&report)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &report, nil
}

func (m *ReportModel) Insert(report models.Report) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), report)
}

func (m *ReportModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
