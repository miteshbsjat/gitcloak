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

	"github.com/miteshbsjat/gitcloak/pkg/utils"
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
var DECRYPTED_FILE_EXT = ".dcry"

func EncryptedFilePattern(normalFilePattern string) string {
	if !strings.HasSuffix(normalFilePattern, ENCRYPTED_FILE_EXT) {
		return normalFilePattern + ENCRYPTED_FILE_EXT
	}
	return normalFilePattern
}

func DecryptedFilePattern(normalFilePattern string) string {
	if !strings.HasSuffix(normalFilePattern, DECRYPTED_FILE_EXT) {
		return normalFilePattern + DECRYPTED_FILE_EXT
	}
	return normalFilePattern
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

// Remove the given prefix path from given filepath
func RemovePathPrefix(filepath, prefix string) string {
	if strings.HasPrefix(filepath, prefix) {
		return strings.TrimPrefix(filepath, prefix)
	}
	return filepath
}

// Create Shell Script from given list of lines
func CreateShellScript(scriptPath string, lines []string) error {
	scriptContent := strings.Join(lines, utils.LineBreak())

	// You can modify the directory path where you want to create the script.

	// Write the script content to the script file.
	err := os.WriteFile(scriptPath, []byte(scriptContent), 0755)
	if err != nil {
		return err
	}

	return nil
}

func IsPresent(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
