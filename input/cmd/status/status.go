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
	if len(args) == 1 {
		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		args := &parkingLotRPC.ParkingLotRPC{}

		var reply parkingLotRPC.ParkingLotRPC

		err = client.Call("Serve.Status", args, &reply)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Invalid slot value")
	}
}
