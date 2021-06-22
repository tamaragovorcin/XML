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
type CartModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *CartModel) GetAll() ([]models.Cart, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.Cart{}

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
func (m *CartModel) FindByID(id string) (*models.Cart, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find user by id
	var product = models.Cart{}
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
func (m *CartModel) Insert(product models.Cart) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), product)
}

// Delete will be used to delete a user
func (m *CartModel) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {


	return m.C.DeleteOne(context.TODO(), bson.M{"_id": id})
}
