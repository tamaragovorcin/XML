package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"users/pkg/models"
)

// UserModel represent a mgo database session with a user model data.
type AgentModel struct {
	C *mongo.Collection
}

func (m *AgentModel) GetAll() ([]models.Agent, error) {
	ctx := context.TODO()
	uu := []models.Agent{}

	agentCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = agentCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

func (m *AgentModel) FindByID(id string) (*models.Agent, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var agent = models.Agent{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&agent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &agent, nil
}

func (m *AgentModel) Insert(agent models.Agent) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), agent)
}

func (m *AgentModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
