package utils

import (
	"fmt"
	"os"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Return False if Error
func CheckIfError2(err error, format string, args ...interface{}) bool {
	if err == nil {
		return false
	}
	Error3(format, args)
	Error3("\x1b[31;1m%s\x1b[0m\n", err)
	// os.Exit(0)
	return true
}

// Info should be used to describe the example commands that are about to run.
func Infop(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warningp(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
