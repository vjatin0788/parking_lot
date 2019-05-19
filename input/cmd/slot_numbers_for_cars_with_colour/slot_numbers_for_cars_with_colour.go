package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/parking_lot/parkingLotRPC"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		color := os.Args[1]

		client, err := rpc.DialHTTP("tcp", "localhost:9999")
		if err != nil {
			fmt.Println(err)
			return
		}

		args := &parkingLotRPC.ParkingLotRPC{
			Color: color,
		}

		var reply []int

		err = client.Call("Serve.GetSlotWithColor", args, &reply)
		if err != nil {
			fmt.Println(err)
		}
	}
}
