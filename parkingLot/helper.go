package parkingLot

import (
	"errors"
	"strings"

	errs "github.com/parking_lot/errs"
	"github.com/parking_lot/util"
	"github.com/parking_lot/vehicle"
)

func (p *ParkingLot) EnablePrint() *ParkingLot {
	p.PrintEnabled = true
	return p
}

func (p *ParkingLot) DisablePrint() *ParkingLot {
	p.PrintEnabled = false
	return p
}

func (p *ParkingLot) addNumberOfSlots(number int64) *ParkingLot {
	p.NumberOfSlots = number
	return p
}

func (p *ParkingLot) initSlotHeap() *ParkingLot {

	if p.Slots == nil {
		var (
			slot []int64
			i    int64
		)

		slot = make([]int64, p.NumberOfSlots)
		for i = 0; i < p.NumberOfSlots; i++ {
			slot[i] = i + 1
		}

		hp := util.InitHeap(p.NumberOfSlots, 2)
		hp.BuildHeapMin(slot, p.NumberOfSlots)

		p.Slots = hp
	}

	return p
}

func (p *ParkingLot) addVehicleToSlot(slot int64, veh *vehicle.Vehicle) (err error) {
	if p.VehicleSlot == nil {
		p.VehicleSlot = make(map[int64]vehicle.Vehicle)
	}

	//check empty veh
	if veh == nil {
		err = errors.New(errs.ERR_INVALID_VEHICLE)
		return
	}

	//check if slot is already occupoed
	if _, ok := p.VehicleSlot[slot]; ok {
		err = errors.New(errs.ERR_SLOT_OCCUPIED)
		return
	}

	p.VehicleSlot[slot] = *veh
	return
}

func (p *ParkingLot) addRegistrationSlot(slot int64, veh *vehicle.Vehicle) (err error) {
	if p.RegistrationSlot == nil {
		p.RegistrationSlot = make(map[string]int64)
	}

	//check empty veh
	if veh == nil {
		err = errors.New(errs.ERR_INVALID_VEHICLE)
		return
	}

	//check if slot is already occupoed
	if _, ok := p.RegistrationSlot[strings.ToUpper(veh.RegisterationNumber)]; !ok {
		p.RegistrationSlot[strings.ToUpper(veh.RegisterationNumber)] = slot
	}
	return
}

func (p *ParkingLot) addColorRegistrationSlot(slot int64, veh *vehicle.Vehicle) (err error) {
	if p.ColorRegistrationSlot == nil {
		p.ColorRegistrationSlot = make(map[string]map[string]int64)
	}

	//check empty veh
	if veh == nil {
		err = errors.New(errs.ERR_INVALID_VEHICLE)
		return
	}

	//check if slot is already occupoed
	if _, ok := p.ColorRegistrationSlot[veh.Color]; !ok {
		p.ColorRegistrationSlot[veh.Color] = make(map[string]int64)
	}

	regSlot := p.ColorRegistrationSlot[veh.Color]
	if _, ok := regSlot[veh.RegisterationNumber]; !ok {
		regSlot[veh.RegisterationNumber] = slot
	}

	return
}

func (p *ParkingLot) getSlotFromHeap() (s int64, err error) {
	if p.Slots == nil {
		s = -1
		err = errors.New(errs.ERR_SLOT_NOT_INIT)
		return
	}

	s = p.Slots.DeleteMin()
	if s <= 0 {
		err = errors.New(errs.ERR_PARKING_LOT_FULL)
		return
	}
	p.SlotsAvailable = p.Slots.Count

	return
}

func (p *ParkingLot) addSlotToHeap(slot int64) (err error) {
	if p.Slots == nil {
		err = errors.New(errs.ERR_SLOT_NOT_INIT)
		return
	}

	p.Slots.InsertMin(slot)
	p.SlotsAvailable = p.Slots.Count
	return
}

func (p *ParkingLot) isFull() (b bool) {
	if p.SlotsAvailable == p.NumberOfSlots {
		b = true
		return
	}
	return
}

func (p *ParkingLot) getVehicleFromSlot(slot int64) (veh vehicle.Vehicle, err error) {
	if p.VehicleSlot == nil {
		err = errors.New(errs.ERR_VEHICLE_SLOT)
		return
	}

	if vehicle, ok := p.VehicleSlot[slot]; ok {
		veh = vehicle
	}
	return
}

func (p *ParkingLot) getSlotFromRegistration(veh vehicle.Vehicle) (slot int64, err error) {
	if p.RegistrationSlot == nil {
		err = errors.New(errs.ERR_REGISTER_SLOT)
		return
	}

	slot = -1
	if s, ok := p.RegistrationSlot[veh.RegisterationNumber]; ok {
		slot = s
	}
	return
}

func (p ParkingLot) getRegistrationNumberFromColor(veh vehicle.Vehicle) (regs []string, err error) {
	if p.ColorRegistrationSlot == nil {
		err = errors.New(errs.ERR_COLOR_REGISTER_SLOT)
		return
	}

	regs = make([]string, 0)
	if veh.Color != "" {
		if regslot, ok := p.ColorRegistrationSlot[veh.Color]; ok {
			for key := range regslot {
				regs = append(regs, key)
			}
		}
	}

	return
}

func (p ParkingLot) getSlotsFromColor(veh vehicle.Vehicle) (slot []int64, err error) {
	if p.ColorRegistrationSlot == nil {
		err = errors.New(errs.ERR_COLOR_REGISTER_SLOT)
		return
	}

	slot = make([]int64, 0)
	if veh.Color != "" {
		if regslot, ok := p.ColorRegistrationSlot[veh.Color]; ok {
			for _, slots := range regslot {
				slot = append(slot, slots)
			}
		}
	}

	return
}

func (p *ParkingLot) deleteVehicleFromSlot(slot int64) (err error) {
	if p.VehicleSlot == nil {
		err = errors.New(errs.ERR_VEHICLE_SLOT)
		return
	}

	delete(p.VehicleSlot, slot)
	return
}

func (p *ParkingLot) deleteSlotFromRegistration(veh vehicle.Vehicle) (err error) {
	if p.RegistrationSlot == nil {
		err = errors.New(errs.ERR_REGISTER_SLOT)
		return
	}

	delete(p.RegistrationSlot, veh.RegisterationNumber)
	return
}

func (p ParkingLot) deleteRegistrationNumberFromColor(veh vehicle.Vehicle) (err error) {
	if p.ColorRegistrationSlot == nil {
		err = errors.New(errs.ERR_COLOR_REGISTER_SLOT)
		return
	}

	if veh.Color != "" {
		if regslot, ok := p.ColorRegistrationSlot[veh.Color]; ok {
			delete(regslot, veh.RegisterationNumber)
			p.ColorRegistrationSlot[veh.Color] = regslot
		}
	}

	return
}
