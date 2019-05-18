package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/parking_lot/parkingLotRPC"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		regno := os.Args[1]
		color := os.Args[2]

		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		args := &parkingLotRPC.ParkingLotRPC{
			Color:             color,
			RegistratioNumber: regno,
		}

		var reply parkingLotRPC.ParkingLotRPC

		err = client.Call("Serve.Parking", args, &reply)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Invalid value")
	}
}
