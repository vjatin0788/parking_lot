package input

import (
	"fmt"
	"io"
	"os"
	"testing"

	errs "github.com/parking_lot/errs"
)

func TestReadFile(t *testing.T) {
	data, err := ReadFile("file_input.txt")
	if err != nil && err != io.EOF {
		t.Errorf("Test Failed, %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Test Failed, Empty String:%v", data)
	}
}

func TestGetFile(t *testing.T) {
	wd, _ := os.Getwd()
	res := fmt.Sprintf("%s/%s", wd, "files/file_input.txt")
	str := getFile("files/file_input.txt")
	if str != res {
		t.Errorf("Test Failed, expected:%s , found:%s", res, str)
	}
}

func TestProcessCommandInvalidCommand(t *testing.T) {
	res := errs.ERR_COMMAND_NOT_ALLOWED
	err := processCommands([]string{"qwerty"})
	if err != nil && err.Error() != res {
		t.Errorf("Test Failed, expected:%s , found:%s", res, err.Error())
	}
}

func TestProcessCommandInit(t *testing.T) {
	resp := errs.ERR_INVALID_SLOT_VALUE
	err := processCommands([]string{"create_parking_lot", "6"})
	if err != nil && err.Error() != resp {
		t.Errorf("Test Failed, expected:%s, found:%s", resp, err.Error())
	}

}
