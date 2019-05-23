package parkingLot

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	errs "github.com/parking_lot/errs"
	"github.com/parking_lot/vehicle"
)

var (
	parlingLotInstance *ParkingLot
	once               sync.Once
)

//make singleton instance
func MakeParkingLot() ParkingLotClient {
	once.Do(
		func() {
			parlingLotInstance = &ParkingLot{}
		})
	return parlingLotInstance
}

//Initialize parking lot with default values, printEnable is used to enable print output in functions.
func (p *ParkingLot) InitParkingLot(slots int64, printEnabled bool) (err error) {
	if slots == 0 {
		err = errors.New(errs.ERR_INVALID_SLOT_VALUE)
		return
	}

	if !p.IsParkingLotInitialized {
		p.IsParkingLotInitialized = true

		//it will create the parking slot of size n.
		p.addNumberOfSlots(slots).
			initSlotHeap()

		if printEnabled {
			p.EnablePrint()
		}

		if p.PrintEnabled {
			fmt.Printf("Created a parking lot with %d slots\n", slots)
		}

		return
	}

	err = errors.New(errs.ERR_PARKING_LOT_INIT)
	return
}

//Park vehicle will fetch the nearest empty slot from Min heap.
//It will add the vehicle to vehicle slot map.
//It will add the registration number to registration slot map
//It will add the registration data, color and slot to the color registration slot map.
//Time complexity for adding to map O(1), Fetching from Min heap is O(1)
func (p *ParkingLot) ParkVehicle(veh vehicle.Vehicle) (err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	if veh.Color == "" || veh.RegisterationNumber == "" {
		err = errors.New(errs.ERR_EMPTY_VEHICLE_DATA)
		return
	}

	if p.isFull() {
		err = errors.New(errs.ERR_PARKING_LOT_FULL)
		return
	}

	veh.RegisterationNumber = strings.ToUpper(veh.RegisterationNumber)
	veh.Color = strings.Title(veh.Color)

	if s, ok := p.RegistrationSlot[veh.RegisterationNumber]; !ok {
		var slot int64

		slot, err = p.getSlotFromHeap()
		if err != nil {
			return
		}

		err = p.addVehicleToSlot(slot, &veh)
		if err != nil {
			return
		}

		err = p.addRegistrationSlot(slot, &veh)
		if err != nil {
			return
		}

		err = p.addColorRegistrationSlot(slot, &veh)
		if err != nil {
			return
		}

		if p.PrintEnabled {
			fmt.Printf("Allocated slot number: %d\n", slot)
		}

	} else {
		err = errors.New(fmt.Sprintf("%s , %s on slot: %d", errs.ERR_CAR_ALREADY_PARKED, veh.RegisterationNumber, s))
	}

	return
}

//This function will remove the vehicle from slot and slot will be added to heap.
//This will remove the slot and vehicle data from all the maps.
//Time complexity for heap will be O(Logn).
func (p *ParkingLot) LeaveVehicle(slot int64) (err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	if slot > p.NumberOfSlots {
		err = errors.New(errs.ERR_INVALID_SLOT_VALUE)
		return
	}

	if veh, ok := p.VehicleSlot[slot]; ok {
		err = p.addSlotToHeap(slot)
		if err != nil {
			return
		}

		err = p.deleteVehicleFromSlot(slot)
		if err != nil {
			return
		}

		err = p.deleteSlotFromRegistration(veh)
		if err != nil {
			return
		}

		err = p.deleteRegistrationNumberFromColor(veh)
		if err != nil {
			return
		}

		if p.PrintEnabled {
			fmt.Printf("Slot number %d is free\n", slot)
		}
	} else {
		err = errors.New(errs.ERR_SLOT_EMPTY)
	}

	return
}

//It will print the status of the parking lot.
//time complexity will  be O(n).
func (p *ParkingLot) ParkingLotStatus() (err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	fmt.Println("Slot No.    Registration No    Colour")
	var idx int64

	for idx = 1; idx <= p.NumberOfSlots; idx++ {
		if veh, ok := p.VehicleSlot[idx]; ok {
			fmt.Printf("%-12v%-19v%-6v\n", idx, strings.ToUpper(veh.RegisterationNumber), veh.Color)
		}
	}

	return
}

//It will fetch the registration number from the color registration slot map.
//Time complexity for fetching from map O(1)
//Sorting the string array. O(klogk) where `k` is subset of `n`.
func (p *ParkingLot) GetRegistrationNumWithColor(color string) (res []string, err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	color = strings.Title(color)

	res = make([]string, 0)
	if reg, ok := p.ColorRegistrationSlot[color]; ok {
		for regNo := range reg {
			res = append(res, regNo)
		}
	}

	if len(res) == 0 {
		err = errors.New(errs.ERR_EMPTY_REG_FOR_COLOR)
	}

	sort.Strings(res)

	if p.PrintEnabled {
		for _, reg := range res {
			fmt.Printf("%s ", reg)
		}
		fmt.Println()
	}

	return
}

//It will fetch the slot from the color registration slot map.
//Time complexity for fetching from map O(1)
//Sorting the string array. O(klogk) where `k` is subset of `n`.
func (p *ParkingLot) GetSlotNumsForCarWithColor(color string) (res []int, err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	color = strings.Title(color)

	res = make([]int, 0)
	if reg, ok := p.ColorRegistrationSlot[color]; ok {
		for _, slot := range reg {
			res = append(res, int(slot))
		}
	}

	if len(res) == 0 {
		err = errors.New(errs.ERR_EMPTY_REG_FOR_COLOR)
	}

	sort.Ints(res)

	if p.PrintEnabled {
		for _, sl := range res {
			fmt.Printf("%d ", sl)
		}
		fmt.Println()
	}

	return
}

//It will fetch the registration number from the registration slot map.
//Time complexity for fetching from map O(1)
func (p *ParkingLot) GetSlotWithRegisterationNum(register string) (s int64, err error) {
	if !p.IsParkingLotInitialized {
		err = errors.New(errs.ERR_PARKING_LOT_NOT_INIT)
		return
	}

	register = strings.ToUpper(register)

	if slot, ok := p.RegistrationSlot[register]; ok {
		s = slot
	} else {
		err = errors.New(errs.ERR_EMPTY_REG_SLOT)
	}

	if p.PrintEnabled && err == nil {
		fmt.Println(s)
	}

	return
}
