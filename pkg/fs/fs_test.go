package fs

import (
	"os"
	"strings"
	"testing"

	"github.com/miteshbsjat/goshell"
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
	goshell.RunCommand("rm -f " + testFile)
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
	goshell.RunCommand("rm -f " + testFile)
}

func TestGetFilePathId(t *testing.T) {
	basePath := "/home/user/files/"
	filePath := "/home/user/files/documents/report.txt"
	uniqueId := GetFilePathId(filePath, basePath)
	basePath = "/home/user/"
	filePath = "/home/user/documents/report.txt"
	uniqueId2 := GetFilePathId(filePath, basePath)
	if uniqueId != uniqueId2 {
		t.Errorf("Unique ID: %v != %v", uniqueId, uniqueId2)
	}
}
