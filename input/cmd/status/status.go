package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/parking_lot/parkingLotRPC"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		client, err := rpc.DialHTTP("tcp", "localhost:9999")
		if err != nil {
			fmt.Println(err)
			return
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
