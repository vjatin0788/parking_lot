package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/parking_lot/input"
	"github.com/parking_lot/parkingLot"
)

func main() {

	//make parking lot

	parkingLot.MakeParkingLot()

	arg := os.Args
	if len(arg) > 1 {
		fileName := os.Args[1]

		//check for text file
		if strings.Contains(fileName, "txt") {
			readFile, err := input.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
			}
			input.ProcessFile(readFile)
		}
	} else if len(arg) == 1 {
		//interactive console
		input.CreateIntractiveShell().RunShell()
	}

}
