package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/20-VIGNESH-K/EnergyAuditing/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RegisteredUserCollection *mongo.Collection
var WeavingCollection *mongo.Collection
var ResultCollection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	RegisteredUserCollection = client.Database("Energy_Auditing").Collection("RegisteredUsers")
	WeavingCollection = client.Database("Energy_Auditing").Collection("Weaving")
	ResultCollection = client.Database("Energy_Auditing").Collection("Result")
}

func CreateUser(user models.CreateUser) error {

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

func Weaving(w models.Weaving) error {
	v := reflect.ValueOf(&w).Elem()
	for i := 0; i < v.NumField(); i++ {
		// Check if the field is a string
		if v.Field(i).Kind() == reflect.String {
			// If a string value is found, return an error
			return errors.New("string value is not allowed")
		}

		// Check if the value is a pointer and dereference it if needed
		if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
			// If it's a nil pointer, set it to zero value
			v.Field(i).Set(reflect.Zero(v.Field(i).Type().Elem()))
		}
	}
	w.CreatedTime = time.Now()
	_, err := WeavingCollection.InsertOne(context.Background(), w)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	err1 := MultiplyEnergy(w)
	if err1 != nil {
		return fmt.Errorf("failed to calculate %v", err)
	}
	return nil
}

func MultiplyEnergy(m models.Weaving) error {
	log.Println("**************multiplyenergy called ***************************")

	var r models.Result
	fmt.Println("--------------", m, r)
	r.AirjetWeavingMachineResult = m.AirjetWeavingMachine * m.AirjetWeavingMachineEnergy
	r.CompressorResult = m.Compressor * m.CompressorEnergy
	r.DirectWrapingMachineResult = m.DirectWrapingMachine * m.DirectWrapingMachineEnergy
	r.IndirectWrapingMachineResult = m.IndirectWrapingMachine * m.IndirectWrapingMachineEnergy
	r.RapierWeavingMachineResult = m.RapierWeavingMachine * m.RapierWeavingMachineEnergy
	r.SampleDrawingMachineResult = m.SampleDrawingMachine * m.SampleDrawingMachineEnergy
	r.SampleSizingMachineResult = m.SampleSizingMachine * m.SampleSizingMachineEnergy
	r.SievingMachineResult = m.SievingMachine * m.SievingMachineEnergy
	r.SampleWrapingMachineResult = m.SampleWrapingMachine * m.SampleWrapingMachineEnergy

	r.OneMonthResult = r.AirjetWeavingMachineResult + r.CompressorResult + r.DirectWrapingMachineResult + r.IndirectWrapingMachineResult + r.RapierWeavingMachineResult + r.SampleDrawingMachineResult +
		r.SampleSizingMachineResult + r.SievingMachineResult + r.SampleWrapingMachineResult

	r.MonthResult = m.Month * r.OneMonthResult
	r.CreatedTime = time.Now()
	_, err := ResultCollection.InsertOne(context.Background(), r)
	log.Println("results:", r)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)

	}

	return nil
}

func GetUser(user models.GetUser)(*models.CreateUser,string,error){
	var customer models.CreateUser
	filter := bson.M{"email":user.Email}
	err := RegisteredUserCollection.FindOne(context.Background(),filter).Decode(&customer)
	if err != nil{
		return nil,"No Records Found",err
	}
	return &customer,"Success",nil
}
