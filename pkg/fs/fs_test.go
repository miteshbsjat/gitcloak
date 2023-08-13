package fs

import (
	"os"
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

	regex, err := RegexFromPattern(regexPattern)
	if err != nil {
		t.Errorf("Invalid regex pattern: %v", err)
		return
	}

	encRegexPattern := EncryptedFilePattern(regexPattern)
	if encRegexPattern != regexPattern+ENCRYPTED_FILE_EXT {
		t.Errorf("regexPattern %v != %v", regexPattern, encRegexPattern)
	}

	decFile := "test.txt"
	encFile := EncryptedFilePattern(decFile)
	normalFile := DecryptedFileName(encFile)
	if normalFile != decFile {
		t.Errorf("decFile %v != %v", decFile, normalFile)
	}

	fileChannel := make(chan string, 10)
	errorChannel := make(chan error)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go FindMatchingFiles(rootDir, regex, fileChannel, errorChannel, &wg)
	go ProcessFiles(fileChannel, errorChannel, done)

	wg.Wait()
	close(fileChannel)

	<-done

	// Non-blocking getting message from channel
	select {
	case err := <-errorChannel:
		t.Errorf("received error %v", err)
	default:
		t.Log("No error")
	}
}
