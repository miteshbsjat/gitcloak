package fs

import (
	"os"
	"strings"
	"testing"
)

func TestAddLineToFile(t *testing.T) {
	testFile := "/tmp/test123M.txt"

	data := "Some random line"
	err := AddLineToFile(testFile, data)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Adding same line again
	err = AddLineToFile(testFile, data)
	if err != nil {
		t.Errorf(err.Error())
	}

	str, err := FileGetString(testFile)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimRight(str, "\n") != data {
		t.Errorf("%s != %s\n", str, data)
	}
}

func TestAppendLineToFile(t *testing.T) {
	testFile := "/tmp/test123MA.txt"
	os.Remove(testFile)
	data := "Some random line"
	err := AppendLineToFile(testFile, data)
	if err != nil {
		t.Errorf(err.Error())
	}

	str, err := FileGetString(testFile)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimRight(str, "\n") != data {
		t.Errorf("%s != %s\n", str, data)
	}
}
