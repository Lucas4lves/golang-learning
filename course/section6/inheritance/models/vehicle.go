package models

type Vehicle struct {
	VehicleId int
	ModelName string
	FabYear   int
}

type Bycicle struct {
	Vehicle
	HasGearShifting bool
}
