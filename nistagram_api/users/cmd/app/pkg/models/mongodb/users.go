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
type UserModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *UserModel) GetAll() ([]models.User, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.User{}

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
func (m *UserModel) FindByID(id primitive.ObjectID) (*models.User, error) {

	/*p, err := primitive.ObjectIDFromHex(id)
	fmt.Println("LUNAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA %s" , p)
	if err != nil {
		return nil, err
	}
*/
	// Find user by id
	var user = models.User{}
	err := m.C.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

func (m *UserModel)  FindByUsername(username string) (*models.User, error) {
	/*p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fmt.Println("dsgsrgrsgd  %s", p)
	// Find user by id*/
	var user = models.User{}
	err := m.C.FindOne(context.TODO(), bson.M{"profileInformation.username": username}).Decode(&user)

	if err != nil {
		// Checks if the user was not found

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

// Insert will be used to insert a new user
func (m *UserModel) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

// Delete will be used to delete a user
func (m *UserModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

func (m *UserModel) Update(user models.User) (*mongo.UpdateResult, error) {

	return m.C.UpdateOne(context.TODO(),bson.M{"_id":user.Id},bson.D{{"$set",bson.M{"biography":user.Biography,"profileInformation.name":user.ProfileInformation.Name,
		"profileInformation.lastName":user.ProfileInformation.LastName, "profileInformation.username":user.ProfileInformation.Username,"profileInformation.email":user.ProfileInformation.Email,
		"profileInformation.phoneNumber":user.ProfileInformation.PhoneNumber,"profileInformation.dateOfBirth":user.ProfileInformation.DateOfBirth,
		"webSite":user.Website,"private":user.Private,"profileInformation.gender":user.ProfileInformation.Gender}}})
}