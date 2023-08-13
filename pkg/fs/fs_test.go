package fs

import (
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
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

func TestTraversalProcessing(t *testing.T) {
	rootDir := gitcloak.GetGitCloakBase() + "/.."
	// regexPattern := `.*_test.go$`
	regexPattern := `.*/pkg/enc.*/.*_test.go`

	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		t.Errorf("Invalid regex pattern: %v", err)
		return
	}

	fileChannel := make(chan string, 10)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go FindMatchingFiles(rootDir, regex, fileChannel, &wg)
	go ProcessFiles(fileChannel, done)

	wg.Wait()
	close(fileChannel)

	<-done
	// t.Errorf("All filenames displayed.")
}
