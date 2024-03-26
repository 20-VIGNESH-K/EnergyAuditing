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
	SievingMachine               int64 `json:"sievingmachine" bson:"sievingmachine"`
	IndirectWrapingMachine       int64 `json:"indirectwrapingmachine" bson:"indirectwrapingmachine"`
	DirectWrapingMachine         int64 `json:"directwrapingmachine" bson:"directwrapingmachine"`
	SampleSizingMachine          int64 `json:"samplesizingmachine" bson:"samplesizingmachine"`
	SampleWrapingMachine         int64 `json:"samplewrapingmachine" bson:"samplewrapingmachine"`
	SampleDrawingMachine         int64 `json:"sampledrawingmachine,omitempty" bson:"sampledrawingmachine,omitempty"`
	AirjetWeavingMachine         int64 `json:"airjetweavingmachine,omitempty" bson:"airjetweavingmachine,omitempty"`
	RapierWeavingMachine         int64 `json:"rapierweavingmachine,omitempty" bson:"rapierweavingmachine,omitempty"`
	Compressor                   int64 `json:"compressor,omitempty" bson:"compressor,omitempty"`
	SievingMachineEnergy         int64 `json:"sievingmachine_energy" bson:"sievingmachine_energy"`
	IndirectWrapingMachineEnergy int64 `json:"indirectwrapingmachine_energy" bson:"indirectwrapingmachine_energy"`
	DirectWrapingMachineEnergy   int64 `json:"directwrapingmachine_energy" bson:"directwrapingmachine_energy"`
	SampleSizingMachineEnergy    int64 `json:"samplesizingmachine_energy" bson:"samplesizingmachine_energy"`
	SampleWrapingMachineEnergy   int64 `json:"samplewrapingmachine_energy" bson:"samplewrapingmachine_energy"`
	SampleDrawingMachineEnergy   int64 `json:"sampledrawingmachine_energy,omitempty" bson:"sampledrawingmachine_energy,omitempty"`
	AirjetWeavingMachineEnergy   int64 `json:"airjetweavingmachine_energy,omitempty" bson:"airjetweavingmachine_energy,omitempty"`
	RapierWeavingMachineEnergy   int64 `json:"rapierweavingmachine_energy,omitempty" bson:"rapierweavingmachine_energy,omitempty"`
	CompressorEnergy             int64 `json:"compressor_energy,omitempty" bson:"compressor_energy,omitempty"`
	Month                        int64 `json:"month" bson:"month"`
	CreatedTime                  time.Time `json:"createdtime" bson:"createdtime"`
}

type Result struct {
	SievingMachineResult         int64 `json:"sievingmachine_result" bson:"sievingmachine_result"`
	IndirectWrapingMachineResult int64 `json:"indirectwrapingmachine_result" bson:"indirectwrapingmachine_result"`
	DirectWrapingMachineResult   int64 `json:"directwrapingmachine_result" bson:"directwrapingmachine_result"`
	SampleSizingMachineResult    int64 `json:"samplesizingmachine_result" bson:"samplesizingmachine_result"`
	SampleWrapingMachineResult   int64 `json:"samplewrapingmachine_result" bson:"samplewrapingmachine_result"`
	SampleDrawingMachineResult   int64 `json:"sampledrawingmachine_result,omitempty" bson:"sampledrawingmachine_result,omitempty"`
	AirjetWeavingMachineResult   int64 `json:"airjetweavingmachine_result,omitempty" bson:"airjetweavingmachine_result,omitempty"`
	RapierWeavingMachineResult   int64 `json:"rapierweavingmachine_result,omitempty" bson:"rapierweavingmachine_result,omitempty"`
	CompressorResult             int64 `json:"compressor_result,omitempty" bson:"compressor_result,omitempty"`
	OneMonthResult               int64 `json:"onemonth" bson:"onemonth"`
	MonthResult                  int64 `json:"monthresult" bson:"monthresult"`
	CreatedTime                  time.Time `json:"createdtime" bson:"createdtime"`
}

type GetUser struct {
	Email    string `json:"email" bson:"email"`
}
