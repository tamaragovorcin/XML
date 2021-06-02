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
type RoleModel struct {
	C *mongo.Collection
}

func (m *RoleModel) GetAll() ([]models.Role, error) {
	ctx := context.TODO()
	uu := []models.Role{}

	roleCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = roleCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

func (m *RoleModel) FindByID(id string) (*models.Role, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var role = models.Role{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &role, nil
}

func (m *RoleModel) Insert(role models.Role) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), role)
}

func (m *RoleModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
