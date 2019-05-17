package main

import (
	"fmt"
	"parking_lot/parking_lot/parkingLot"
	"parking_lot/parking_lot/vehicle"
)

func main() {
	fmt.Println("Welcome")

	parking := parkingLot.MakeParkingLot()

	fmt.Println("create_parking_lot", 6)
	fmt.Println("park KA-01-HH-1234 White")
	fmt.Println("park KA-01-HH-9999 White")
	fmt.Println("park KA-01-BB-0001 Black")
	fmt.Println("park KA-01-HH-7777 Red")
	fmt.Println("park KA-01-HH-2701 Blue")
	fmt.Println("park KA-01-HH-3141 Black")
	fmt.Println("park KA-01-P-333 White")
	fmt.Println("park DL-12-AA-9999 White")
	fmt.Println("leave 4")
	fmt.Println("status")
	fmt.Println("park KA-01-P-333 White")
	fmt.Println("park DL-12-AA-9999 White")
	fmt.Println("registration_numbers_for_cars_with_colour White")
	fmt.Println("slot_numbers_for_cars_with_colour White")
	fmt.Println("slot_number_for_registration_number KA-01-HH-3141")
	fmt.Println("slot_number_for_registration_number MH-04-AY-1111")
	fmt.Println("--------------------------->>>>>>>>>>Result<<<<<<<<<<------------------------")

	parking.InitParkingLot(6, true)
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-HH-1234", "White"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-HH-9999", "White"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-BB-0001", "Black"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-HH-7777", "Red"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-HH-2701", "Blue"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-HH-3141", "Black"})
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-P-333", "White"})
	fmt.Println(parking.ParkVehicle(vehicle.Vehicle{"DL-12-AA-9999", "White"}))
	parking.LeaveVehicle(4)
	parking.ParkingLotStatus()
	parking.ParkVehicle(vehicle.Vehicle{"KA-01-P-333", "White"})
	fmt.Println(parking.ParkVehicle(vehicle.Vehicle{"DL-12-AA-9999", "White"}))
	fmt.Println(parking.GetRegistrationNumWithColor("White"))
	fmt.Println(parking.GetSlotNumsForCarWithColor("White"))
	fmt.Println(parking.GetSlotWithRegisterationNum("KA-01-HH-3141"))
	fmt.Println(parking.GetSlotWithRegisterationNum("MH-04-AY-1111"))
	fmt.Println("--------------------------->>>>>>>>>>End<<<<<<<<<<------------------------")
}
