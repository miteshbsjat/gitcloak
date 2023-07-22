package gitcloak

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addLineToFile(filePath, lineToAdd string) error {
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

func appendLineToFile(filePath, line string) error {
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
