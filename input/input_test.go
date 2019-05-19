package input

import (
	"fmt"
	"io"
	"os"
	"testing"
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
