package main

import (
	"fmt"
	"inheritance/models"
)

func main() {
	car := models.Vehicle{
		VehicleId: 1,
		ModelName: "Monza",
		FabYear:   1986,
	}

	bike := models.Bycicle{
		Vehicle: models.Vehicle{
			VehicleId: 2,
			ModelName: "Monark",
			FabYear:   1977,
		},
		HasGearShifting: true,
	}

	fmt.Println(car)
	fmt.Println(bike.VehicleId)
}
