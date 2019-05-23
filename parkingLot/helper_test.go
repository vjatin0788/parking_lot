package parkingLot

import (
	"testing"

	errs "github.com/parking_lot/errs"
	"github.com/parking_lot/vehicle"
)

func TestInitSlotHeap(t *testing.T) {
	p := &ParkingLot{NumberOfSlots: 5}
	p.initSlotHeap()

	if p.Slots == nil {
		t.Errorf("Test Failed, Empty slot")
	}

	if p.Slots != nil && p.Slots.Capacity != p.NumberOfSlots {
		t.Errorf("Test Failed, expected:%d, found:%d", p.NumberOfSlots, p.Slots.Capacity)
	}
}

func TestAddVehicleToSlot(t *testing.T) {
	p := &ParkingLot{NumberOfSlots: 5}
	p.initSlotHeap()

	err := p.addVehicleToSlot(1, nil)
	if err == nil {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_INVALID_VEHICLE, err)
	}

	if err != nil && err.Error() != errs.ERR_INVALID_VEHICLE {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_INVALID_VEHICLE, err)
	}

	err = p.addVehicleToSlot(1, vehicle.InitVehicle().AddColor("Blue").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed %v", err)
	}

	err = p.addVehicleToSlot(1, vehicle.InitVehicle().AddColor("Blue").AddRegistrationNumber("PB-03-AZ-1239"))
	if err == nil {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_INVALID_VEHICLE, err)
	}

	if err != nil && err.Error() != errs.ERR_SLOT_OCCUPIED {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_SLOT_OCCUPIED, err)
	}

}

func TestAddRegToSlot(t *testing.T) {
	p := &ParkingLot{NumberOfSlots: 5}
	p.initSlotHeap()

	err := p.addRegistrationSlot(1, nil)
	if err == nil {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_INVALID_VEHICLE, err)
	}

	if err != nil && err.Error() != errs.ERR_INVALID_VEHICLE {
		t.Errorf("Test Failed, expected:%s, found:%v", errs.ERR_INVALID_VEHICLE, err)
	}
}

func TestGetSlotFromHeap(t *testing.T) {
	p := &ParkingLot{NumberOfSlots: 1}
	p.initSlotHeap()

	p.getSlotFromHeap()

	s, err := p.getSlotFromHeap()
	if err != nil && s != -1 {
		t.Errorf("Test Failed %v", err)
	}
}
