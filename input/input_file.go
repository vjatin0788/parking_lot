package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	errs "github.com/parking_lot/errs"
	parking "github.com/parking_lot/parkingLot"
	"github.com/parking_lot/vehicle"
)

//only allowed commands are executed.
var allowedCommands = map[string]bool{
	CREATE_PARKING_LOT:  true,
	PARK:                true,
	LEAVE:               true,
	STATUS:              true,
	REG_NUM_WITH_COLOR:  true,
	SLOT_NUM_WITH_COLOR: true,
	SLOT_NUM_REG_NUM:    true,
	EXIT:                true,
}

func ReadFile(fileName string) (lines []string, err error) {
	file, err := os.Open(getFile(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		lines = append(lines, strings.TrimSpace(command))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
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

//process file commands
func processCommands(words []string) (err error) {
	parkingLot := parking.MakeParkingLot()

	if len(words) > 0 {
		word := words[0]

		if _, ok := allowedCommands[word]; ok {
			switch word {
			case CREATE_PARKING_LOT:
				if !(len(words) > 1) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

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

			case PARK:
				if !(len(words) > 2) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

				regNo := words[1]
				color := words[2]

				veh := vehicle.InitVehicle().
					AddColor(color).
					AddRegistrationNumber(regNo)

				err = parkingLot.ParkVehicle(*veh)
				if err != nil {
					return
				}

			case LEAVE:
				if !(len(words) > 1) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

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

			case STATUS:
				err = parkingLot.ParkingLotStatus()
				if err != nil {
					return
				}
			case REG_NUM_WITH_COLOR:

				if !(len(words) > 1) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

				color := words[1]

				_, err = parkingLot.GetRegistrationNumWithColor(color)
				if err != nil {
					return
				}
			case SLOT_NUM_WITH_COLOR:

				if !(len(words) > 1) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

				color := words[1]

				_, err = parkingLot.GetSlotNumsForCarWithColor(color)
				if err != nil {
					return
				}

			case SLOT_NUM_REG_NUM:

				if !(len(words) > 1) {
					err = errors.New(errs.ERR_INVALID_ARGUMENT)
					return
				}

				regNo := words[1]

				_, err = parkingLot.GetSlotWithRegisterationNum(regNo)
				if err != nil {
					return
				}
			case EXIT:
				os.Exit(0)
			}
		} else {
			err = errors.New(errs.ERR_COMMAND_NOT_ALLOWED)
		}
	}

	return
}
