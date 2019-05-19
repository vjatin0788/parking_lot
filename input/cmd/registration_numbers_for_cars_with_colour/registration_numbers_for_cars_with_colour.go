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
		color := os.Args[1]

		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		args := &parkingLotRPC.ParkingLotRPC{
			Color: color,
		}

		var reply []string

		err = client.Call("Serve.GetRegNumWithColor", args, &reply)
		if err != nil {
			fmt.Println(err)
		}
	}
}
