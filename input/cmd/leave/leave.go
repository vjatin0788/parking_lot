package main

import (
	"fmt"
	"log"
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

		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		args := &parkingLotRPC.ParkingLotRPC{
			Slot: slot,
		}

		var reply parkingLotRPC.ParkingLotRPC

		err = client.Call("Serve.Leave", args, &reply)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Invalid slot value")
	}
}
