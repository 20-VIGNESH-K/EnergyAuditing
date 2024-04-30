package models

import "time"

type CreateUser struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Address  string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	Country  string `json:"country" bson:"country"`
	Zipcode  string `json:"zipcode" bson:"zipcode"`
}

type LogIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Weaving struct {
	SievingMachine               int       `json:"sievingmachine" bson:"sievingmachine"`
	IndirectWrapingMachine       int       `json:"indirectwrapingmachine" bson:"indirectwrapingmachine"`
	DirectWrapingMachine         int       `json:"directwrapingmachine" bson:"directwrapingmachine"`
	SampleSizingMachine          int       `json:"samplesizingmachine" bson:"samplesizingmachine"`
	SampleWrapingMachine         int       `json:"samplewrapingmachine" bson:"samplewrapingmachine"`
	SampleDrawingMachine         int       `json:"sampledrawingmachine,omitempty" bson:"sampledrawingmachine,omitempty"`
	AirjetWeavingMachine         int       `json:"airjetweavingmachine,omitempty" bson:"airjetweavingmachine,omitempty"`
	RapierWeavingMachine         int       `json:"rapierweavingmachine,omitempty" bson:"rapierweavingmachine,omitempty"`
	Compressor                   int       `json:"compressor,omitempty" bson:"compressor,omitempty"`
	SievingMachineEnergy         int       `json:"sievingmachine_energy" bson:"sievingmachine_energy"`
	IndirectWrapingMachineEnergy int       `json:"indirectwrapingmachine_energy" bson:"indirectwrapingmachine_energy"`
	DirectWrapingMachineEnergy   int       `json:"directwrapingmachine_energy" bson:"directwrapingmachine_energy"`
	SampleSizingMachineEnergy    int       `json:"samplesizingmachine_energy" bson:"samplesizingmachine_energy"`
	SampleWrapingMachineEnergy   int       `json:"samplewrapingmachine_energy" bson:"samplewrapingmachine_energy"`
	SampleDrawingMachineEnergy   int       `json:"sampledrawingmachine_energy,omitempty" bson:"sampledrawingmachine_energy,omitempty"`
	AirjetWeavingMachineEnergy   int       `json:"airjetweavingmachine_energy,omitempty" bson:"airjetweavingmachine_energy,omitempty"`
	RapierWeavingMachineEnergy   int       `json:"rapierweavingmachine_energy,omitempty" bson:"rapierweavingmachine_energy,omitempty"`
	CompressorEnergy             int       `json:"compressor_energy,omitempty" bson:"compressor_energy,omitempty"`
	Month                        int       `json:"month" bson:"month"`
	CreatedTime                  time.Time `json:"createdtime" bson:"createdtime"`
	WeavingID                    string    `json:"weavingid" bson:"weavingid"`
}
type WeavingResult struct {
	SievingMachineResult         int       `json:"sievingmachine_result" bson:"sievingmachine_result"`
	IndirectWrapingMachineResult int       `json:"indirectwrapingmachine_result" bson:"indirectwrapingmachine_result"`
	DirectWrapingMachineResult   int       `json:"directwrapingmachine_result" bson:"directwrapingmachine_result"`
	SampleSizingMachineResult    int       `json:"samplesizingmachine_result" bson:"samplesizingmachine_result"`
	SampleWrapingMachineResult   int       `json:"samplewrapingmachine_result" bson:"samplewrapingmachine_result"`
	SampleDrawingMachineResult   int       `json:"sampledrawingmachine_result,omitempty" bson:"sampledrawingmachine_result,omitempty"`
	AirjetWeavingMachineResult   int       `json:"airjetweavingmachine_result,omitempty" bson:"airjetweavingmachine_result,omitempty"`
	RapierWeavingMachineResult   int       `json:"rapierweavingmachine_result,omitempty" bson:"rapierweavingmachine_result,omitempty"`
	CompressorResult             int       `json:"compressor_result,omitempty" bson:"compressor_result,omitempty"`
	MonthResult                  int       `json:"monthresult" bson:"monthresult"`
	CreatedTime                  time.Time `json:"createdtime" bson:"createdtime"`
	WeavingID                    string    `json:"weavingid" bson:"weavingid"`
}

type Textile struct {
	CreatedTime              time.Time `json:"createdtime" bson:"createdtime"`
	TextileID                string    `json:"textileid" bson:"textileid"`
	SpinningMachine          int       `json:"spinningmachine" bson:"spinningmachine"`
	SpinningMachineEnergy    int       `json:"spinningmachineenergy" bson:"spinningmachineenergy"`
	WeavingLoom              int       `json:"weavingloom" bson:"weavingloom"`
	WeavingLoomEnergy        int       `json:"weavingloomenergy" bson:"weavingloomenergy"`
	KnittingMachine          int       `json:"knittingmachine" bson:"knittingmachine"`
	KnittingMachineEnergy    int       `json:"knittingmachineenergy" bson:"knittingmachineenergy"`
	DyeingMachine            int       `json:"dyeingmachine" bson:"dyeingmachine"`
	DyeingMachineEnergy      int       `json:"dyeingmachineenergy" bson:"dyeingmachineenergy"`
	PrintingMachine          int       `json:"printingmachine" bson:"printingmachine"`
	PrintingMachineEnergy    int       `json:"printingmachineenergy" bson:"printingmachineenergy"`
	FinishingMachine         int       `json:"finishingmachine" bson:"finishingmachine"`
	FinishingMachineEnergy   int       `json:"finishingmachineenergy" bson:"finishingmachineenergy"`
	CuttingMachine           int       `json:"cuttingmachine" bson:"cuttingmachine"`
	CuttingMachineEnergy     int       `json:"cuttingmachineenergy" bson:"cuttingmachineenergy"`
	SewingMachine            int       `json:"sewingmachine" bson:"sewingmachine"`
	SewingMachineEnergy      int       `json:"sewingmachineenergy" bson:"sewingmachineenergy"`
	EmbroideryMachine        int       `json:"embroiderymachine" bson:"embroiderymachine"`
	EmbroideryMachineEnergy  int       `json:"embroiderymachineenergy" bson:"embroiderymachineenergy"`
	SteamersAndPresses       int       `json:"steamersandpresses" bson:"steamersandpresses"`
	SteamersAndPressesEnergy int       `json:"steamersandpressesenergy" bson:"steamersandpressesenergy"`
	Month                    int       `json:"month" bson:"month"`
}

type TextileResult struct {
	TextileID                string    `json:"textileid" bson:"textileid"`
	CreatedTime              time.Time `json:"createdtime" bson:"createdtime"`
	SpinningMachineResult    int       `json:"spinningmachineresult" bson:"spinningmachineresult"`
	WeavingLoomResult        int       `json:"weavingloomresult" bson:"weavingloomresult"`
	KnittingMachineResult    int       `json:"knittingmachineresult" bson:"knittingmachineresult"`
	DyeingMachineResult      int       `json:"dyeingmachineresult" bson:"dyeingmachineresult"`
	PrintingMachineResult    int       `json:"printingmachineresult" bson:"printingmachineresult"`
	FinishingMachineResult   int       `json:"finishingmachineresult" bson:"finishingmachineresult"`
	CuttingMachineResult     int       `json:"cuttingmachineresult" bson:"cuttingmachineresult"`
	SewingMachineResult      int       `json:"sewingmachineresult" bson:"sewingmachineresult"`
	EmbroideryMachineResult  int       `json:"embroiderymachineresult" bson:"embroiderymachineresult"`
	SteamersAndPressesResult int       `json:"steamersandpressesresult" bson:"steamersandpressesresult"`
	MonthResult              int       `json:"monthresult" bson:"monthresult"`
}

type IT struct {
	CreatedTime              time.Time `json:"createdtime" bson:"createdtime"`
	ITID                     string    `json:"itid" bson:"itid"`
	DataStorageSystems       int       `json:"datastoragesystems" bson:"datastoragesystems"`
	DesktopComputers         int       `json:"desktopcomputers" bson:"desktopcomputers"`
	Laptops                  int       `json:"laptops" bson:"laptops"`
	Monitors                 int       `json:"monitors" bson:"monitors"`
	Printers                 int       `json:"printers" bson:"printers"`
	CoolingSystems           int       `json:"coolingsystems" bson:"coolingsystems"`
	UPS                      int       `json:"ups" bson:"ups"`
	DataStorageSystemsEnergy int       `json:"datastoragesystemsenergy" bson:"datastoragesystemsenergy"`
	DesktopComputersEnergy   int       `json:"desktopcomputersenergy" bson:"desktopcomputersenergy"`
	LaptopsEnergy            int       `json:"laptopsenergy" bson:"laptopsenergy"`
	MonitorsEnergy           int       `json:"monitorsenergy" bson:"monitorsenergy"`
	PrintersEnergy           int       `json:"printerenergy" bson:"printerenergy"`
	CoolingSystemsEnergy     int       `json:"coolingsystemsenergy" bson:"coolingsystemsenergy"`
	UPSEnergy                int       `json:"upsenergy" bson:"upsenergy"`
	Month                    int       `json:"month" bson:"month"`
}

type ITResult struct {
	CreatedTime              time.Time `json:"createdtime" bson:"createdtime"`
	ITID                     string    `json:"itid" bson:"itid"`
	DataStorageSystemsResult int       `json:"datastoragesystemsresult" bson:"datastoragesystemsresult"`
	DesktopComputersResult   int       `json:"desktopcomputersresult" bson:"desktopcomputersresult"`
	LaptopsResult            int       `json:"laptopsresult" bson:"laptopsresult"`
	MonitorsResult           int       `json:"monitorsresult" bson:"monitorsresult"`
	PrintersResult           int       `json:"printersresult" bson:"printersresult"`
	CoolingSystemsResult     int       `json:"coolingsystemsresult" bson:"coolingsystemsresult"`
	UPSResult                int       `json:"upsresult" bson:"upsresult"`
	MonthResult              int       `json:"monthresult" bson:"monthresult"`
}

type GetUser struct {
	Email string `json:"email" bson:"email"`
}

type EnergyAuditResult struct {
	Machine                string `json:"machine" bson:"machine"`
	BestMachine            string `json:"bestmachine" bson:"bestmachine"`
	MachineEnergyUsage     int    `json:"machineenergyusage" bson:"machineenergyusage"`
	BestMachineEnergyUsage int    `json:"bestmachineenergyusage" bson:"bestmachineenergyusage"`
	Energy_Saved           int    `json:"energysaved" bson:"energysaved"`
	IsBest                 bool   `json:"isbest" bson:"isbest"`
}

type EnergyAuditStore struct {
	ID          string              `json:"id" bson:"id"`
	CreatedTime time.Time           `json:"createdtime" bson:"createdtime"`
	AuditResult []EnergyAuditResult `json:"energyauditresult" bson:"energyauditresult"`
}
