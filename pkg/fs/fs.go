package fs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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