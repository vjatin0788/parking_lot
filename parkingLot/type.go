package parkingLot

import (
	"github.com/parking_lot/util"
	"github.com/parking_lot/vehicle"
)

//To implement the parking lot, Certain set of data structure is use to solve the problem efficiently.
// Min heap is used to store the empty slots and fetch nearest slot.
//Maps are used to fetch the data optimally.
type ParkingLot struct {
	//Total number of slots in parking lot.
	NumberOfSlots           int64
	//Min Heap containing the empty slots.
	Slots                   *util.HeapNode
	//boolean to store whether parking lot is init or not.
	IsParkingLotInitialized bool
	//number of slots left empty.
	SlotsAvailable          int64
	//vehicle slot map. this will help us to fetch the status of the parking lot.
	VehicleSlot             map[int64]vehicle.Vehicle
	//registration slot map for storing the slot number corresponding registration number.
	RegistrationSlot        map[string]int64
	//This will help us to solve the queries like slot_number_for_registration_number and slot_numbers_for_cars_with_colour
	ColorRegistrationSlot   map[string]map[string]int64
	PrintEnabled            bool
}

type ParkingLotClient interface {
	InitParkingLot(slots int64, printEnabled bool) error
	ParkVehicle(veh vehicle.Vehicle) error
	LeaveVehicle(slot int64) error
	ParkingLotStatus() error
	GetRegistrationNumWithColor(color string) ([]string, error)
	GetSlotNumsForCarWithColor(color string) ([]int, error)
	GetSlotWithRegisterationNum(register string) (int64, error)
}
