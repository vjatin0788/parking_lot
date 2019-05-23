package parkingLot

import (
	"fmt"
	"testing"

	errs "github.com/parking_lot/errs"
	"github.com/parking_lot/vehicle"
)

func TestParkingLotInitInvalidSlotValue(t *testing.T) {
	resp := errs.ERR_INVALID_SLOT_VALUE
	err := MakeParkingLot().InitParkingLot(0, true)
	if err != nil && err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}
}

func TestParkingLotInit(t *testing.T) {

	err := MakeParkingLot().InitParkingLot(0, true)
	if err == nil {
		t.Errorf("Test Failed %v", err)
	}
}

func TestParkingLotParkNotInit(t *testing.T) {
	p := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	err := p.ParkVehicle(*vehicle.InitVehicle())
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}
}

func TestParkingLotParkInvalidVehicle(t *testing.T) {
	p := &ParkingLot{}

	p.InitParkingLot(1, true)

	resp := errs.ERR_EMPTY_VEHICLE_DATA
	err := p.ParkVehicle(*vehicle.InitVehicle())
	if err == nil {
		t.Errorf("Test Failed, expected:%s, found:%v", resp, err)
	}
}

func TestParkingLotParkVehicle(t *testing.T) {
	p := &ParkingLot{}
	p.InitParkingLot(1, true)

	err := p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	res := fmt.Sprintf("%s , PB-03-AZ-1234 on slot: 1", errs.ERR_CAR_ALREADY_PARKED)
	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil && err.Error() != res {
		t.Errorf("Test Failed, expected:%s, found:%v", res, err)
	}

	res = errs.ERR_PARKING_LOT_FULL
	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1233"))
	if err != nil && err.Error() != res {
		t.Errorf("Test Failed, expected:%s, found:%v", res, err)
	}

}

func TestParkingLotLeaveNotInit(t *testing.T) {
	p := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	err := p.LeaveVehicle(1)
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}
}

func TestParkingLotLeaveInvalidSlot(t *testing.T) {
	p := &ParkingLot{}
	p.InitParkingLot(1, true)

	err := p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	resp := errs.ERR_INVALID_SLOT_VALUE
	err = p.LeaveVehicle(2)
	if err == nil {
		t.Errorf("Test Failed, expected:%s, found:%v", resp, err)
	}
}

func TestParkingLotLeave(t *testing.T) {
	p := &ParkingLot{}
	p.InitParkingLot(2, true)

	err := p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1233"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	err = p.LeaveVehicle(2)
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	res := errs.ERR_SLOT_EMPTY
	err = p.LeaveVehicle(2)
	if err != nil && err.Error() != res {
		t.Errorf("Test Failed, expected:%s, found:%v", res, err)
	}
}

func TestParkingLotStatusNotInit(t *testing.T) {
	p := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	err := p.ParkingLotStatus()
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}
}

func TestParkingLotGetRegNumWithColor(t *testing.T) {

	pk := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	_, err := pk.GetRegistrationNumWithColor("White")
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}

	p := MakeParkingLot()

	p.InitParkingLot(2, true)

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1233"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	res, err := p.GetRegistrationNumWithColor("White")
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	expected := []string{"PB-03-AZ-1233", "PB-03-AZ-1234"}
	for idx := range res {
		if res[idx] != expected[idx] {
			t.Errorf("Test Failed, expected:%s, found:%s", expected[idx], res[idx])
		}
	}
}

func TestParkingLotGetSlotWithColor(t *testing.T) {

	pk := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	_, err := pk.GetSlotNumsForCarWithColor("White")
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}

	p := &ParkingLot{}

	p.InitParkingLot(2, true)

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1233"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	res, err := p.GetSlotNumsForCarWithColor("White")
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	expected := []int{1, 2}
	for idx := range res {
		if res[idx] != expected[idx] {
			t.Errorf("Test Failed, expected:%d, found:%d", expected[idx], res[idx])
		}
	}
}

func TestParkingLotGetSlotWithReg(t *testing.T) {

	pk := &ParkingLot{}
	resp := errs.ERR_PARKING_LOT_NOT_INIT
	_, err := pk.GetSlotWithRegisterationNum("PB-03-AZ-1234")
	if err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}

	p := &ParkingLot{}

	p.InitParkingLot(2, true)

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1234"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	err = p.ParkVehicle(*vehicle.InitVehicle().AddColor("white").AddRegistrationNumber("PB-03-AZ-1233"))
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	res, err := p.GetSlotWithRegisterationNum("PB-03-AZ-1234")
	if err != nil {
		t.Errorf("Test Failed, %v", err)
	}

	if res != 1 {
		t.Errorf("Test Failed, expected:%d, found:%d", 1, res)
	}

}
