package fs

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	. "github.com/miteshbsjat/gitcloak/pkg/utils"
)

func AddLineToFile(filePath, lineToAdd string) error {
	// Open the file in read-write mode, create if it doesn't exist, and append if it does exist
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Check if the line already exists in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == lineToAdd {
			return nil
		}
	}

	// The line does not exist, so add it to the file
	_, err = fmt.Fprintln(file, lineToAdd)
	if err != nil {
		return err
	}

	return nil
}

func FileGetBytes(filename string) ([]byte, error) {
	return os.ReadFile(filename) //#nosec G304
}

func FileGetString(filename string, timeout ...time.Duration) (string, error) {
	bytes, err := FileGetBytes(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func AppendLineToFile(filePath, line string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	if err != nil {
		return err
	}

	return nil
}

// Generate unique seed int64 from given filepath
func GetFilePathId(filePath, basePath string) int64 {
	// Remove the common base path from the file path
	relativePath := strings.TrimPrefix(filePath, basePath)

	// Calculate a hash of the relative path using FNV-1a
	hash := fnv.New64a()
	hash.Write([]byte(relativePath))
	return int64(hash.Sum64())
}

func RegexFromPattern(regexPattern string) (*regexp.Regexp, error) {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		Warn("Invalid regex pattern: %v", err)
		return nil, err
	}
	return regex, nil
}

var ENCRYPTED_FILE_EXT = ".ecry"

func EncryptedFilePattern(normalFilePattern string) string {
	return normalFilePattern + ENCRYPTED_FILE_EXT
}

// removes .ecry from encrypted file name given
func DecryptedFileName(encryptedFileName string) string {
	if strings.HasSuffix(encryptedFileName, ENCRYPTED_FILE_EXT) {
		return encryptedFileName[:len(encryptedFileName)-len(ENCRYPTED_FILE_EXT)]
	}
	return encryptedFileName
}

// func findMatchingFiles(rootDir string, regex *regexp.Regexp, fileChannel chan<- string, errorChannel chan<- error, wg *sync.WaitGroup) {
func FindMatchingFiles(rootDir string, regex *regexp.Regexp, fileChannel chan<- string, errorChannel chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Warn("Error: %v", err)
			return err
		}

		if !info.IsDir() && regex.MatchString(path) {
			fileChannel <- path
		}
		return nil
	})

	if err != nil {
		errorChannel <- err
	}
}

func ProcessFiles(fileChannel <-chan string, errorChannel chan<- error, done chan<- bool) {
	for filename := range fileChannel {
		fmt.Println("Filename:", filename)
	}

	done <- true
}
