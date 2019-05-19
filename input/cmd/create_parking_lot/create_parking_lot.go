package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"

	"github.com/parking_lot/parkingLotRPC"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		slotStr := os.Args[1]
		slot, err := strconv.ParseInt(slotStr, 10, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		client, err := rpc.DialHTTP("tcp", "localhost:9999")
		if err != nil {
			fmt.Println(err)
			return
		}

		args := &parkingLotRPC.ParkingLotRPC{
			Slot: slot,
		}

		var reply parkingLotRPC.ParkingLotRPC

		err = client.Call("Serve.InitParking", args, &reply)
		if err != nil {
			fmt.Println(err)
		}
	}
}
