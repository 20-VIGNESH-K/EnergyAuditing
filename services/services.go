package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/20-VIGNESH-K/EnergyAuditing/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RegisteredUserCollection *mongo.Collection
var WeavingCollection *mongo.Collection
var ResultCollection *mongo.Collection
var AuditCollection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	RegisteredUserCollection = client.Database("Energy_Auditing").Collection("RegisteredUsers")
	WeavingCollection = client.Database("Energy_Auditing").Collection("Weaving")
	ResultCollection = client.Database("Energy_Auditing").Collection("Result")
	AuditCollection = client.Database("Energy_Auditing").Collection("FinalWeavingAudit")
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

func Weaving(w models.Weaving) ([]models.EnergyAuditResult, error) {
	//v := reflect.ValueOf(&w).Elem()
	// for i := 0; i < v.NumField(); i++ {
	// 	// Check if the field is a string
	// 	if v.Field(i).Kind() == reflect.String {
	// 		// If a string value is found, return an error
	// 		return nil,errors.New("string value is not allowed")
	// 	}

	// 	// Check if the value is a pointer and dereference it if needed
	// 	if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
	// 		// If it's a nil pointer, set it to zero value
	// 		v.Field(i).Set(reflect.Zero(v.Field(i).Type().Elem()))
	// 	}
	// }
	w.WeavingID = GetRandomString(22)
	w.CreatedTime = time.Now()
	_, err := WeavingCollection.InsertOne(context.Background(), w)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	results, err1 := MultiplyEnergy(w)
	if err1 != nil {
		return nil, fmt.Errorf("failed to calculate %v", err)
	}

	return results, nil
}

func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func MultiplyEnergy(m models.Weaving) ([]models.EnergyAuditResult, error) {
	log.Println("**************multiplyenergy called ***************************")

	var r models.Result
	r.AirjetWeavingMachineResult = m.AirjetWeavingMachine * m.AirjetWeavingMachineEnergy * m.Month
	r.CompressorResult = m.Compressor * m.CompressorEnergy * m.Month
	r.DirectWrapingMachineResult = m.DirectWrapingMachine * m.DirectWrapingMachineEnergy * m.Month
	r.IndirectWrapingMachineResult = m.IndirectWrapingMachine * m.IndirectWrapingMachineEnergy * m.Month
	r.RapierWeavingMachineResult = m.RapierWeavingMachine * m.RapierWeavingMachineEnergy * m.Month
	r.SampleDrawingMachineResult = m.SampleDrawingMachine * m.SampleDrawingMachineEnergy * m.Month
	r.SampleSizingMachineResult = m.SampleSizingMachine * m.SampleSizingMachineEnergy * m.Month
	r.SievingMachineResult = m.SievingMachine * m.SievingMachineEnergy * m.Month
	r.SampleWrapingMachineResult = m.SampleWrapingMachine * m.SampleWrapingMachineEnergy * m.Month

	r.MonthResult = r.AirjetWeavingMachineResult + r.CompressorResult + r.DirectWrapingMachineResult + r.IndirectWrapingMachineResult + r.RapierWeavingMachineResult + r.SampleDrawingMachineResult +
		r.SampleSizingMachineResult + r.SievingMachineResult + r.SampleWrapingMachineResult

	r.CreatedTime = time.Now()
	r.WeavingID = m.WeavingID
	_, err := ResultCollection.InsertOne(context.Background(), r)
	log.Println("results:", r)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)

	}

	var results []models.EnergyAuditResult
	var bestSievingMachine = 10
	var bestIndirectWrapingMachine = 10
	var bestDirectWrapingMachine = 10
	var bestSampleSizingMachine = 10
	var bestSampleWrapingMachine = 10
	var bestSampleDrawingMachine = 10
	var bestAirjetWeavingMachine = 10
	var bestRapierWeavingMachine = 10
	var bestCompressor = 10
	bestSievingMachine = bestSievingMachine * m.SievingMachine * m.Month
	bestIndirectWrapingMachine = bestIndirectWrapingMachine * m.IndirectWrapingMachine * m.Month
	bestDirectWrapingMachine = bestDirectWrapingMachine * m.DirectWrapingMachine * m.Month
	bestSampleSizingMachine = bestSampleSizingMachine * m.SampleSizingMachine * m.Month
	bestSampleWrapingMachine = bestSampleWrapingMachine * m.SampleWrapingMachine * m.Month
	bestSampleDrawingMachine = bestSampleDrawingMachine * m.SampleDrawingMachine * m.Month
	bestAirjetWeavingMachine = bestAirjetWeavingMachine * m.AirjetWeavingMachine * m.Month
	bestRapierWeavingMachine = bestRapierWeavingMachine * m.RapierWeavingMachine * m.Month
	bestCompressor = bestCompressor * m.Compressor * m.Month

	appendResult := func(machine string, bestmachine string, machineenergyusege int, bestMachineEnergyUsage int, energySaved int, isBest bool) {
		results = append(results, models.EnergyAuditResult{
			Machine:                machine,
			BestMachine:            bestmachine,
			MachineEnergyUsage:     machineenergyusege,
			BestMachineEnergyUsage: bestMachineEnergyUsage,
			Energy_Saved:           energySaved,
			IsBest:                 isBest,
		})
	}

	if r.SievingMachineResult-bestSievingMachine <= 0 {
		appendResult("Sieving Machine", "Best Sieving Machine",r.SievingMachineResult, bestSievingMachine,  0, true)
	} else {
		appendResult("Sieving Machine", "Best Sieving Machine", r.SievingMachineResult,bestSievingMachine, r.SievingMachineResult-bestSievingMachine, false)
	}
	if r.IndirectWrapingMachineResult-bestIndirectWrapingMachine <= 0 {
		appendResult("Indirect Wraping Machine", "Best Indirect Wraping Machine",r.IndirectWrapingMachineResult,bestIndirectWrapingMachine, 0, true)
	} else {
		appendResult("Indirect Wraping Machine", "Best Indirect Wraping Machine",r.IndirectWrapingMachineResult,bestIndirectWrapingMachine, r.IndirectWrapingMachineResult-bestIndirectWrapingMachine, false)
	}
	if r.DirectWrapingMachineResult-bestDirectWrapingMachine <= 0 {
		appendResult("Direct Wraping Machine", "Best Direct Wraping Machine",r.DirectWrapingMachineResult,bestDirectWrapingMachine,  0, true)
	} else {
		appendResult("Direct Wraping Machine", "Best Direct Wraping Machine", r.DirectWrapingMachineResult,bestDirectWrapingMachine,r.DirectWrapingMachineResult-bestDirectWrapingMachine, false)
	}
	if r.SampleSizingMachineResult-bestSampleSizingMachine <= 0 {
		appendResult("Sample Sizing Machine", "Best Sample Sizing Machine",r.SampleSizingMachineResult,bestSampleSizingMachine , 0, true)
	} else {
		appendResult("Sample Sizing Machine", "Best Sample Sizing Machine",r.SampleSizingMachineResult,bestSampleSizingMachine , r.SampleSizingMachineResult-bestSampleSizingMachine, false)
	}
	if r.SampleWrapingMachineResult-bestSampleWrapingMachine <= 0 {
		appendResult("Sample Wraping Machine", "Best Sample Wraping Machine",r.SampleWrapingMachineResult,bestSampleWrapingMachine, 0, true)
	} else {
		appendResult("Sample Wraping Machine", "Best Sample Wraping Machine",r.SampleWrapingMachineResult,bestSampleWrapingMachine, r.SampleWrapingMachineResult-bestSampleWrapingMachine, false)
	}
	if r.SampleDrawingMachineResult-bestSampleDrawingMachine <= 0 {
		appendResult("Sample Drawing Machine", "Best Sample Drawing Machine",r.SampleDrawingMachineResult,bestSampleDrawingMachine, 0, true)
	} else {
		appendResult("Sample Drawing Machine", "Best Sample Drawing Machine",r.SampleDrawingMachineResult,bestSampleDrawingMachine, r.SampleDrawingMachineResult-bestSampleDrawingMachine, false)
	}
	if r.AirjetWeavingMachineResult-bestAirjetWeavingMachine <= 0 {
		appendResult("Airjet Weaving Machine", "Best Airjet Weaving Machine",r.AirjetWeavingMachineResult,bestAirjetWeavingMachine, 0, true)
	} else {
		appendResult("Airjet Weaving Machine", "Best Airjet Weaving Machine",r.AirjetWeavingMachineResult,bestAirjetWeavingMachine, r.AirjetWeavingMachineResult-bestAirjetWeavingMachine, false)
	}
	if r.RapierWeavingMachineResult-bestRapierWeavingMachine <= 0 {
		appendResult("Rapier Weaving Machine", "Best Rapier Weaving Machine",r.RapierWeavingMachineResult,bestRapierWeavingMachine, 0, true)
	} else {
		appendResult("Rapier Weaving Machine", "Best Rapier Weaving Machine", r.RapierWeavingMachineResult,bestRapierWeavingMachine,r.RapierWeavingMachineResult-bestRapierWeavingMachine, false)
	}
	if r.CompressorResult-bestCompressor <= 0 {
		appendResult("Compressor", "Best Compressor", r.CompressorResult,bestCompressor , 0, true)
	} else {
		appendResult("Compressor", "Best Compressor", r.CompressorResult,bestCompressor ,r.CompressorResult-bestCompressor, false)
	}
	var resultstore models.EnergyAuditStore
	resultstore.CreatedTime = time.Now()
	resultstore.WeavingID = r.WeavingID
	resultstore.AuditResult = results
	_, err = AuditCollection.InsertOne(context.Background(), resultstore)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func GetUser(user models.GetUser) (*models.CreateUser, string, error) {
	var customer models.CreateUser
	filter := bson.M{"email": user.Email}
	err := RegisteredUserCollection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return nil, "No Records Found", err
	}
	return &customer, "Success", nil
}
