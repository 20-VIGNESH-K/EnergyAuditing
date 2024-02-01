package services

import (
	"context"
	"errors"

	"github.com/20-VIGNESH-K/EnergyAuditing/models"
	"github.com/20-VIGNESH-K/EnergyAuditing/validation"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RegisteredUserCollection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	RegisteredUserCollection = client.Database("Energy_Auditing").Collection("RegisteredUsers")
}

func CreateUser(user models.CreateUser) error {

	name := validation.IsValidName(user.Name)
	if !name {
		return errors.New("invalid Name")
	}
	email := validation.IsValidEmail(user.Email)
	if !email {
		return errors.New("invalid Email")
	}
	phoneNumber := validation.IsValidPhoneNumber(user.Phone)
	if !phoneNumber {
		return errors.New("invalid Phone Number")
	}

	existingUser, err := GetUserByEmail(user.Email)
	if err == nil && existingUser.Email == user.Email {
		return errors.New("user already exists")
	}

	_, err = RegisteredUserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
func LogIn(user models.LogIn) error {
	var user1 models.LogIn
	err := RegisteredUserCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&user1)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found
			return errors.New("user does not exist")
		}
		// Other database error
		return err
	}

	if user.Password != user1.Password {
		return errors.New("password does not match")
	}

	return nil
}

func GetUserByEmail(email string) (models.CreateUser, error) {
	var user models.CreateUser
	err := RegisteredUserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}
