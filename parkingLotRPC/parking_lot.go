package parkingLotRPC

import (
	parking "github.com/parking_lot/parkingLot"
	"github.com/parking_lot/vehicle"
)

//RPC is used for Inter process communication.
//Interactive shell commands will intract with this app through RPC.
//Seprate functions are implemented for seprate commands.
type ParkingLotRPC struct {
	Slot              int64
	Color             string
	RegistratioNumber string
}

type Serve ParkingLotRPC

func (p *Serve) InitParking(args *ParkingLotRPC, reply *ParkingLotRPC) (err error) {
	parkingLot := parking.MakeParkingLot()
	err = parkingLot.InitParkingLot(args.Slot, true)
	reply = nil
	return
}

func (p *Serve) Parking(args *ParkingLotRPC, reply *ParkingLotRPC) (err error) {
	parkingLot := parking.MakeParkingLot()

	veh := vehicle.InitVehicle().
		AddColor(args.Color).
		AddRegistrationNumber(args.RegistratioNumber)

	err = parkingLot.ParkVehicle(*veh)
	reply = nil
	return
}

func (p *Serve) Leave(args *ParkingLotRPC, reply *ParkingLotRPC) (err error) {
	parkingLot := parking.MakeParkingLot()

	err = parkingLot.LeaveVehicle(args.Slot)
	reply = nil
	return
}

func (p *Serve) GetRegNumWithColor(args *ParkingLotRPC, reply *[]string) (err error) {
	parkingLot := parking.MakeParkingLot()

	var res []string
	res, err = parkingLot.GetRegistrationNumWithColor(args.Color)
	reply = &res
	return
}

func (p *Serve) GetSlotFromReg(args *ParkingLotRPC, reply *int64) (err error) {
	parkingLot := parking.MakeParkingLot()

	var res int64
	res, err = parkingLot.GetSlotWithRegisterationNum(args.RegistratioNumber)
	reply = &res
	return
}

func (p *Serve) GetSlotWithColor(args *ParkingLotRPC, reply *[]int) (err error) {
	parkingLot := parking.MakeParkingLot()

	var res []int
	res, err = parkingLot.GetSlotNumsForCarWithColor(args.Color)
	reply = &res
	return
}

func (p *Serve) Status(args *ParkingLotRPC, reply *ParkingLotRPC) (err error) {
	parkingLot := parking.MakeParkingLot()

	err = parkingLot.ParkingLotStatus()
	reply = nil
	return
}
