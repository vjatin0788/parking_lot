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
	if len(args) == 2 {
		reg := os.Args[1]

		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		args := &parkingLotRPC.ParkingLotRPC{
			RegistratioNumber: reg,
		}

		var reply int64

		err = client.Call("Serve.GetSlotFromReg", args, &reply)
		if err != nil {
			fmt.Println(err)
		}
	}
}
