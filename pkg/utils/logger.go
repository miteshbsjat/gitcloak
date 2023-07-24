package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

var logger *log.Logger
var once sync.Once

// InitLogger initializes the central singleton logger with the specified log file path.
// If the log file path is empty, the logger will write to os.Stderr by default.
func InitLogger(logFilePath string) {
	once.Do(func() {
		var output io.Writer
		if logFilePath != "" {
			file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				output = file
			} else {
				log.Println("Failed to open log file:", err)
				output = os.Stderr
			}
		} else {
			output = os.Stderr
		}

		logger = log.New(output, "", log.Ldate|log.Ltime)
	})
}

// GetLogger returns the central singleton logger instance.
func GetLogger() *log.Logger {
	if logger == nil {
		InitLogger("")
	}
	return logger
}

func trimLeftToMax(str string, max int) string {
	if len(str) > max {
		// Trim from the left if the string length exceeds the maximum
		return "..." + str[len(str)-max+3:]
	}
	return str
}

// logWithCallerInfo adds file line number and function name to the log message.
func logWithCallerInfo(depth int, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	funcName := trimLeftToMax(getCurrentFuncName(depth), 15)
	file = trimLeftToMax(file, 20)

	msg := fmt.Sprintf(format, args...)
	msgWithCallerInfo := fmt.Sprintf("%s:%d %s - %s", file, line, funcName, msg)

	GetLogger().Println(msgWithCallerInfo)
}

func getCurrentFuncName(depth int) string {
	pc, _, _, _ := runtime.Caller(depth + 1)
	return runtime.FuncForPC(pc).Name()
}

// Info logs an info level message with file line number and function name.
func Info(format string, args ...interface{}) {
	logWithCallerInfo(2, "[INFO]  "+format, args...)
}

// Error logs an error level message with file line number and function name.
func Error(format string, args ...interface{}) {
	logWithCallerInfo(2, "[ERROR] "+format, args...)
}

func Error3(format string, args ...interface{}) {
	logWithCallerInfo(3, "[ERROR] "+format, args...)
}

func Warn(format string, args ...interface{}) {
	logWithCallerInfo(2, "[WARN]  "+format, args...)
}
