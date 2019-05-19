package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	parking "github.com/parking_lot/parkingLot"
	"github.com/parking_lot/vehicle"
)

const (
	ERR_COMMAND_NOT_ALLOWED = "Command not allowed"
)

var allowedCommands = map[string]bool{
	"create_parking_lot": true,
	"park":               true,
	"leave":              true,
	"status":             true,
	"registration_numbers_for_cars_with_colour": true,
	"slot_numbers_for_cars_with_colour":         true,
	"slot_number_for_registration_number":       true,
}

func ReadFile(fileName string) (lines []string, err error) {
	file, err := os.Open(getFile(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	for err != io.EOF {
		// Includes the delimiter
		var l string
		l, err = br.ReadString('\n')

		if err != nil && err != io.EOF {
			return
		}

		// Trimming space to remove the delimiter at the end
		lines = append(lines, strings.TrimSpace(l))
	}

	return
}

//Get file from root
func getFile(fileName string) string {
	wd, _ := os.Getwd()

	if !strings.HasSuffix(wd, "file") {
		wd += ""
	}

	return wd + "/" + fileName
}

//Process input read file
func ProcessFile(readFile []string) {
	for _, value := range readFile {
		inputLine := strings.Split(value, "\n")
		if len(inputLine) > 0 {
			ilString := inputLine[0]
			words := strings.Fields(ilString)
			err := processCommands(words)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
}

func processCommands(words []string) (err error) {
	parkingLot := parking.MakeParkingLot()

	if len(words) > 1 {
		word := words[0]

		if _, ok := allowedCommands[word]; ok {
			switch word {
			case "create_parking_lot":
				slotStr := words[1]

				var slot int64
				slot, err = strconv.ParseInt(slotStr, 10, 64)
				if err != nil {
					return
				}

				err = parkingLot.InitParkingLot(slot, true)
				if err != nil {
					return
				}

			case "park":
				if len(words) > 2 {
					regNo := words[1]
					color := words[2]

					veh := vehicle.InitVehicle().
						AddColor(color).
						AddRegistrationNumber(regNo)

					err = parkingLot.ParkVehicle(*veh)
					if err != nil {
						return
					}
				}

			case "leave":
				slotStr := words[1]

				var slot int64
				slot, err = strconv.ParseInt(slotStr, 10, 64)
				if err != nil {
					return
				}

				err = parkingLot.LeaveVehicle(slot)
				if err != nil {
					return
				}

			case "status":
				err = parkingLot.ParkingLotStatus()
				if err != nil {
					return
				}
			case "registration_numbers_for_cars_with_colour":
				color := words[1]

				_, err = parkingLot.GetRegistrationNumWithColor(color)
				if err != nil {
					return
				}
			case "slot_numbers_for_cars_with_colour":
				color := words[1]

				_, err = parkingLot.GetSlotNumsForCarWithColor(color)
				if err != nil {
					return
				}

			case "slot_number_for_registration_number":
				regNo := words[1]

				_, err = parkingLot.GetSlotWithRegisterationNum(regNo)
				if err != nil {
					return
				}
			}
		} else {
			err = errors.New(ERR_COMMAND_NOT_ALLOWED)
		}
	}

	return
}
