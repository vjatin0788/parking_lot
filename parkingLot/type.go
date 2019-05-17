package parkingLot

import (
	"parking_lot/parking_lot/util"
	"parking_lot/parking_lot/vehicle"
)

type ParkingLot struct {
	NumberOfSlots           int64
	Slots                   *util.HeapNode
	IsParkingLotInitialized bool
	SlotsAvailable          int64
	VehicleSlot             map[int64]vehicle.Vehicle
	RegistrationSlot        map[string]int64
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
