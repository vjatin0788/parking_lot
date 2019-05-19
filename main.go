package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"

	"github.com/parking_lot/input"
	"github.com/parking_lot/parkingLot"
	"github.com/parking_lot/parkingLotRPC"
)

func main() {

	//make parking lot

	parkingLot.MakeParkingLot()

	serve := new(parkingLotRPC.Serve)

	rpc.Register(serve)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":9999")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)

	arg := os.Args
	if len(arg) > 1 {
		fileName := os.Args[1]

		//check for text file
		if strings.Contains(fileName, "txt") {
			readFile, err := input.ReadFile(fileName)
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			input.ProcessFile(readFile)
		}
	} else if len(arg) == 1 {
		//interactive console
		input.CreateIntractiveShell()
	}

}
