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
var WeavingAuditCollection *mongo.Collection
var WeavingResultCollection *mongo.Collection

var TextileCollection *mongo.Collection
var TextileAuditCollection *mongo.Collection
var TextileResultCollection *mongo.Collection

var ITCollection *mongo.Collection
var ITAuditCollection *mongo.Collection
var ITResultCollection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://vignesh:Vignesh123@vigneshk.pvmu4rp.mongodb.net/?retryWrites=true&w=majority&appName=VigneshK"))
	if err != nil {
		panic(err)
	}
	RegisteredUserCollection = client.Database("Energy_Auditing").Collection("RegisteredUsers")
	WeavingCollection = client.Database("Energy_Auditing").Collection("Weaving")
	WeavingResultCollection = client.Database("Energy_Auditing").Collection("WeavingResult")
	WeavingAuditCollection = client.Database("Energy_Auditing").Collection("FinalWeavingAudit")
	TextileCollection = client.Database("Energy_Auditing").Collection("Textile")
	TextileResultCollection = client.Database("Energy_Auditing").Collection("TextileResult")
	TextileAuditCollection = client.Database("Energy_Auditing").Collection("FinalTextileAudit")
	ITCollection = client.Database("Energy_Auditing").Collection("IT")
	ITResultCollection = client.Database("Energy_Auditing").Collection("ITResult")
	ITAuditCollection = client.Database("Energy_Auditing").Collection("FinalITAudit")

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
	if user.Password == "" {
		return errors.New("please provide a password")
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

	w.WeavingID = GetRandomString(22)
	w.CreatedTime = time.Now()
	_, err := WeavingCollection.InsertOne(context.Background(), w)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	results, err1 := MultiplyWeavingEnergy(w)
	if err1 != nil {
		return nil, fmt.Errorf("failed to calculate %v", err)
	}

	return results, nil
}

func Textile(w models.Textile) ([]models.EnergyAuditResult, error) {

	w.TextileID = GetRandomString(22)
	w.CreatedTime = time.Now()
	_, err := TextileCollection.InsertOne(context.Background(), w)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	results, err1 := MultiplyTextileEnergy(w)
	if err1 != nil {
		return nil, fmt.Errorf("failed to calculate %v", err)
	}

	return results, nil
}
func IT(w models.IT) ([]models.EnergyAuditResult, error) {

	w.ITID = GetRandomString(22)
	w.CreatedTime = time.Now()
	_, err := ITCollection.InsertOne(context.Background(), w)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	results, err1 := MultiplyITEnergy(w)
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

func MultiplyWeavingEnergy(m models.Weaving) ([]models.EnergyAuditResult, error) {
	log.Println("**************multiplyenergy called ***************************")

	var r models.WeavingResult
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
	_, err := WeavingResultCollection.InsertOne(context.Background(), r)
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
		appendResult("Sieving Machine", "Best Sieving Machine", r.SievingMachineResult, bestSievingMachine, 0, true)
	} else {
		appendResult("Sieving Machine", "Best Sieving Machine", r.SievingMachineResult, bestSievingMachine, r.SievingMachineResult-bestSievingMachine, false)
	}
	if r.IndirectWrapingMachineResult-bestIndirectWrapingMachine <= 0 {
		appendResult("Indirect Wraping Machine", "Best Indirect Wraping Machine", r.IndirectWrapingMachineResult, bestIndirectWrapingMachine, 0, true)
	} else {
		appendResult("Indirect Wraping Machine", "Best Indirect Wraping Machine", r.IndirectWrapingMachineResult, bestIndirectWrapingMachine, r.IndirectWrapingMachineResult-bestIndirectWrapingMachine, false)
	}
	if r.DirectWrapingMachineResult-bestDirectWrapingMachine <= 0 {
		appendResult("Direct Wraping Machine", "Best Direct Wraping Machine", r.DirectWrapingMachineResult, bestDirectWrapingMachine, 0, true)
	} else {
		appendResult("Direct Wraping Machine", "Best Direct Wraping Machine", r.DirectWrapingMachineResult, bestDirectWrapingMachine, r.DirectWrapingMachineResult-bestDirectWrapingMachine, false)
	}
	if r.SampleSizingMachineResult-bestSampleSizingMachine <= 0 {
		appendResult("Sample Sizing Machine", "Best Sample Sizing Machine", r.SampleSizingMachineResult, bestSampleSizingMachine, 0, true)
	} else {
		appendResult("Sample Sizing Machine", "Best Sample Sizing Machine", r.SampleSizingMachineResult, bestSampleSizingMachine, r.SampleSizingMachineResult-bestSampleSizingMachine, false)
	}
	if r.SampleWrapingMachineResult-bestSampleWrapingMachine <= 0 {
		appendResult("Sample Wraping Machine", "Best Sample Wraping Machine", r.SampleWrapingMachineResult, bestSampleWrapingMachine, 0, true)
	} else {
		appendResult("Sample Wraping Machine", "Best Sample Wraping Machine", r.SampleWrapingMachineResult, bestSampleWrapingMachine, r.SampleWrapingMachineResult-bestSampleWrapingMachine, false)
	}
	if r.SampleDrawingMachineResult-bestSampleDrawingMachine <= 0 {
		appendResult("Sample Drawing Machine", "Best Sample Drawing Machine", r.SampleDrawingMachineResult, bestSampleDrawingMachine, 0, true)
	} else {
		appendResult("Sample Drawing Machine", "Best Sample Drawing Machine", r.SampleDrawingMachineResult, bestSampleDrawingMachine, r.SampleDrawingMachineResult-bestSampleDrawingMachine, false)
	}
	if r.AirjetWeavingMachineResult-bestAirjetWeavingMachine <= 0 {
		appendResult("Airjet Weaving Machine", "Best Airjet Weaving Machine", r.AirjetWeavingMachineResult, bestAirjetWeavingMachine, 0, true)
	} else {
		appendResult("Airjet Weaving Machine", "Best Airjet Weaving Machine", r.AirjetWeavingMachineResult, bestAirjetWeavingMachine, r.AirjetWeavingMachineResult-bestAirjetWeavingMachine, false)
	}
	if r.RapierWeavingMachineResult-bestRapierWeavingMachine <= 0 {
		appendResult("Rapier Weaving Machine", "Best Rapier Weaving Machine", r.RapierWeavingMachineResult, bestRapierWeavingMachine, 0, true)
	} else {
		appendResult("Rapier Weaving Machine", "Best Rapier Weaving Machine", r.RapierWeavingMachineResult, bestRapierWeavingMachine, r.RapierWeavingMachineResult-bestRapierWeavingMachine, false)
	}
	if r.CompressorResult-bestCompressor <= 0 {
		appendResult("Compressor", "Best Compressor", r.CompressorResult, bestCompressor, 0, true)
	} else {
		appendResult("Compressor", "Best Compressor", r.CompressorResult, bestCompressor, r.CompressorResult-bestCompressor, false)
	}
	var resultstore models.EnergyAuditStore
	resultstore.CreatedTime = time.Now()
	resultstore.ID = r.WeavingID
	resultstore.AuditResult = results
	_, err = WeavingAuditCollection.InsertOne(context.Background(), resultstore)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func MultiplyTextileEnergy(m models.Textile) ([]models.EnergyAuditResult, error) {
	log.Println("**************multiplyenergy called ***************************")

	var r = models.TextileResult{
		TextileID:                m.TextileID,
		CreatedTime:              time.Now(),
		SpinningMachineResult:    m.SpinningMachine * m.SpinningMachineEnergy * m.Month,
		WeavingLoomResult:        m.WeavingLoom * m.WeavingLoomEnergy * m.Month,
		KnittingMachineResult:    m.KnittingMachine * m.KnittingMachineEnergy * m.Month,
		DyeingMachineResult:      m.DyeingMachine * m.DyeingMachineEnergy * m.Month,
		PrintingMachineResult:    m.PrintingMachine * m.PrintingMachineEnergy * m.Month,
		FinishingMachineResult:   m.FinishingMachine * m.FinishingMachineEnergy * m.Month,
		CuttingMachineResult:     m.CuttingMachine * m.CuttingMachineEnergy * m.Month,
		SewingMachineResult:      m.SewingMachine * m.SewingMachineEnergy * m.Month,
		EmbroideryMachineResult:  m.EmbroideryMachine * m.EmbroideryMachineEnergy * m.Month,
		SteamersAndPressesResult: m.SteamersAndPresses * m.SteamersAndPressesEnergy * m.Month,
	}

	r.MonthResult = r.SpinningMachineResult + r.WeavingLoomResult + r.KnittingMachineResult + r.DyeingMachineResult + r.PrintingMachineResult + r.FinishingMachineResult +
		r.CuttingMachineResult + r.SewingMachineResult + r.EmbroideryMachineResult + r.SteamersAndPressesResult

	_, err := TextileResultCollection.InsertOne(context.Background(), r)
	log.Println("results:", r)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)

	}
	var results []models.EnergyAuditResult
	bestSpinningMachine := 10
	bestWeavingLoom := 10
	bestKnittingMachine := 10
	bestDyeingMachine := 10
	bestPrintingMachine := 10
	bestFinishingMachine := 10
	bestCuttingMachine := 10
	bestSewingMachine := 10
	bestEmbroideryMachine := 10
	bestSteamersAndPresses := 10

	bestSpinningMachine = bestSpinningMachine * m.SpinningMachine * m.Month
	bestWeavingLoom = bestWeavingLoom * m.WeavingLoom * m.Month
	bestKnittingMachine = bestKnittingMachine * m.KnittingMachine * m.Month
	bestDyeingMachine = bestDyeingMachine * m.DyeingMachine * m.Month
	bestPrintingMachine = bestPrintingMachine * m.PrintingMachine * m.Month
	bestFinishingMachine = bestFinishingMachine * m.FinishingMachine * m.Month
	bestCuttingMachine = bestCuttingMachine * m.CuttingMachine * m.Month
	bestSewingMachine = bestSewingMachine * m.SewingMachine * m.Month
	bestEmbroideryMachine = bestEmbroideryMachine * m.EmbroideryMachine * m.Month
	bestSteamersAndPresses = bestSteamersAndPresses * m.SteamersAndPresses * m.Month

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

	if r.SpinningMachineResult-bestSpinningMachine <= 0 {
		appendResult("Spinning Machine", "Best Spinning Machine", r.SpinningMachineResult, bestSpinningMachine, 0, true)
	} else {
		appendResult("Spinning Machine", "Best Spinning Machine", r.SpinningMachineResult, bestSpinningMachine, r.SpinningMachineResult-bestSpinningMachine, false)
	}
	if r.WeavingLoomResult-bestWeavingLoom <= 0 {
		appendResult("WeavingLoom", "Best WeavingLoom", r.WeavingLoomResult, bestWeavingLoom, 0, true)
	} else {
		appendResult("WeavingLoom", "Best WeavingLoom", r.WeavingLoomResult, bestWeavingLoom, r.WeavingLoomResult-bestWeavingLoom, false)
	}
	if r.KnittingMachineResult-bestKnittingMachine <= 0 {
		appendResult("Knitting Machine", "Best Knitting Machine", r.KnittingMachineResult, bestKnittingMachine, 0, true)
	} else {
		appendResult("Knitting Machine", "Best Knitting Machine", r.KnittingMachineResult, bestKnittingMachine, r.KnittingMachineResult-bestKnittingMachine, false)
	}
	if r.DyeingMachineResult-bestDyeingMachine <= 0 {
		appendResult("Dyeing Machine", "Best Dyeing Machine", r.DyeingMachineResult, bestDyeingMachine, 0, true)
	} else {
		appendResult("Dyeing Machine", "Best Dyeing Machine", r.DyeingMachineResult, bestDyeingMachine, r.DyeingMachineResult-bestDyeingMachine, false)
	}
	if r.PrintingMachineResult-bestPrintingMachine <= 0 {
		appendResult("Printing Machine", "Best Printing Machine", r.PrintingMachineResult, bestPrintingMachine, 0, true)
	} else {
		appendResult("Printing Machine", "Best Printing Machine", r.PrintingMachineResult, bestPrintingMachine, r.PrintingMachineResult-bestPrintingMachine, false)
	}
	if r.CuttingMachineResult-bestCuttingMachine <= 0 {
		appendResult("Cutting Machine", "Best Cutting Machine", r.CuttingMachineResult, bestCuttingMachine, 0, true)
	} else {
		appendResult("Cutting Machine", "Best Cutting Machine", r.CuttingMachineResult, bestCuttingMachine, r.CuttingMachineResult-bestCuttingMachine, false)
	}
	if r.FinishingMachineResult-bestFinishingMachine <= 0 {
		appendResult("Finishing Machine", "Best Finishing Machine", r.FinishingMachineResult, bestFinishingMachine, 0, true)
	} else {
		appendResult("Finishing Machine", "Best Finishing Machine", r.FinishingMachineResult, bestFinishingMachine, r.FinishingMachineResult-bestFinishingMachine, false)
	}
	if r.EmbroideryMachineResult-bestEmbroideryMachine <= 0 {
		appendResult("Embroidery Machine", "Best Embroidery Machine", r.EmbroideryMachineResult, bestEmbroideryMachine, 0, true)
	} else {
		appendResult("Embroidery Machine", "Best Embroidery Machine", r.EmbroideryMachineResult, bestEmbroideryMachine, r.EmbroideryMachineResult-bestEmbroideryMachine, false)
	}
	if r.SewingMachineResult-bestSewingMachine <= 0 {
		appendResult("Sewing Machine", "Best Sewing Machine", r.SewingMachineResult, bestSewingMachine, 0, true)
	} else {
		appendResult("Sewing Machine", "Best Sewing Machine", r.SewingMachineResult, bestSewingMachine, r.SewingMachineResult-bestSewingMachine, false)
	}
	if r.SteamersAndPressesResult-bestSteamersAndPresses <= 0 {
		appendResult("Steamers And Presses", "Best Steamers And Presses", r.SteamersAndPressesResult, bestSteamersAndPresses, 0, true)
	} else {
		appendResult("Steamers And Presses", "Best Steamers And Presses", r.SteamersAndPressesResult, bestSteamersAndPresses, r.SteamersAndPressesResult-bestSteamersAndPresses, false)
	}
	var resultstore models.EnergyAuditStore
	resultstore.CreatedTime = time.Now()
	resultstore.ID = r.TextileID
	resultstore.AuditResult = results
	_, err = TextileAuditCollection.InsertOne(context.Background(), resultstore)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func MultiplyITEnergy(m models.IT) ([]models.EnergyAuditResult, error) {
	log.Println("**************multiplyenergy called ***************************")

	var r = models.ITResult{
		CreatedTime:              time.Now(),
		ITID:                     m.ITID,
		DataStorageSystemsResult: m.DataStorageSystems * m.DataStorageSystemsEnergy * m.Month,
		DesktopComputersResult:   m.DesktopComputers * m.DesktopComputersEnergy * m.Month,
		LaptopsResult:            m.Laptops * m.LaptopsEnergy * m.Month,
		MonitorsResult:           m.Monitors * m.MonitorsEnergy * m.Month,
		PrintersResult:           m.Printers * m.PrintersEnergy * m.Month,
		CoolingSystemsResult:     m.CoolingSystems * m.CoolingSystemsEnergy * m.Month,
		UPSResult:                m.UPS * m.UPSEnergy * m.Month,
	}

	r.MonthResult = r.DataStorageSystemsResult + r.DesktopComputersResult + r.LaptopsResult + r.MonitorsResult + r.PrintersResult + r.CoolingSystemsResult + r.UPSResult

	_, err := ITResultCollection.InsertOne(context.Background(), r)
	log.Println("results:", r)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)

	}

	var results []models.EnergyAuditResult
	bestDataStorageSystems := 10
	bestDesktopComputers := 10
	bestLaptops := 10
	bestMonitors := 10
	bestPrinters := 10
	bestCoolingSystems := 10
	bestUPS := 10

	bestDataStorageSystems = bestDataStorageSystems * m.DataStorageSystems * m.Month
	bestDesktopComputers = bestDesktopComputers * m.DesktopComputers * m.Month
	bestLaptops = bestLaptops * m.Laptops * m.Month
	bestMonitors = bestMonitors * m.Monitors * m.Month
	bestPrinters = bestPrinters * m.Printers * m.Month
	bestCoolingSystems = bestCoolingSystems * m.CoolingSystems * m.Month
	bestUPS = bestUPS * m.UPS * m.Month

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

	if r.DataStorageSystemsResult-bestDataStorageSystems <= 0 {
		appendResult("Data Storage Systems", "Best Data Storage Systems", r.DataStorageSystemsResult, bestDataStorageSystems, 0, true)
	} else {
		appendResult("Data Storage Systems", "Best Data Storage Systems", r.DataStorageSystemsResult, bestDataStorageSystems, r.DataStorageSystemsResult-bestDataStorageSystems, false)
	}
	if r.DesktopComputersResult-bestDesktopComputers <= 0 {
		appendResult("Desktop Computers", "Best Desktop Computers", r.DesktopComputersResult, bestDesktopComputers, 0, true)
	} else {
		appendResult("Desktop Computers", "Best Desktop Computers", r.DesktopComputersResult, bestDesktopComputers, r.DesktopComputersResult-bestDesktopComputers, false)
	}
	if r.LaptopsResult-bestLaptops <= 0 {
		appendResult("Laptops", "Best Laptops", r.LaptopsResult, bestLaptops, 0, true)
	} else {
		appendResult("Laptops", "Best Laptops", r.LaptopsResult, bestLaptops, r.LaptopsResult-bestLaptops, false)
	}
	if r.MonitorsResult-bestMonitors <= 0 {
		appendResult("Monitors", "Best Monitors", r.MonitorsResult, bestMonitors, 0, true)
	} else {
		appendResult("Monitors", "Best Monitors", r.MonitorsResult, bestMonitors, r.MonitorsResult-bestMonitors, false)
	}
	if r.PrintersResult-bestPrinters <= 0 {
		appendResult("Printer", "Best Printer", r.PrintersResult, bestPrinters, 0, true)
	} else {
		appendResult("Printer", "Best Printer", r.PrintersResult, bestPrinters, r.PrintersResult-bestPrinters, false)
	}
	if r.CoolingSystemsResult-bestCoolingSystems <= 0 {
		appendResult("Cooling Systems", "Best Cooling Systems", r.CoolingSystemsResult, bestCoolingSystems, 0, true)
	} else {
		appendResult("Cooling Systems", "Best Cooling Systems", r.CoolingSystemsResult, bestCoolingSystems, r.CoolingSystemsResult-bestCoolingSystems, false)
	}
	if r.UPSResult-bestUPS <= 0 {
		appendResult("UPS", "Best UPS", r.UPSResult, bestUPS, 0, true)
	} else {
		appendResult("UPS", "Best UPS", r.UPSResult, bestUPS, r.UPSResult-bestUPS, false)
	}

	var resultstore models.EnergyAuditStore
	resultstore.CreatedTime = time.Now()
	resultstore.ID = r.ITID
	resultstore.AuditResult = results
	_, err = ITAuditCollection.InsertOne(context.Background(), resultstore)
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
