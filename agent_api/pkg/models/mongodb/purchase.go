
package mongodb

import (
"AgentApp/pkg/models"
"context"
"errors"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
)

// UserModel represent a mgo database session with a user model data.
type PurchaseModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *PurchaseModel) GetAll() ([]models.Purchase, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.Purchase{}

	// Find all users
	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

// FindByID will be used to find a new user registry by id
func (m *PurchaseModel) FindByID(id string) (*models.Purchase, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find user by id
	var product = models.Purchase{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&product)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &product, nil
}

// Insert will be used to insert a new user
func (m *PurchaseModel) Insert(purchase models.Purchase) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), purchase)
}

// Delete will be used to delete a user
func (m *PurchaseModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
