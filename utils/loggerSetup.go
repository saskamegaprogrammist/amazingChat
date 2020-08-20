package utils

import (
	"github.com/google/logger"
	"os"
)

const verbose = true

var loggerFile *os.File

func LoggerSetup() {
	var err error
	loggerFile, err = os.OpenFile(LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Faigled to open log file: %v", err)
	}
	logger.Init("LoggerExample", verbose, true, loggerFile)
}

func LoggerClose() {
	loggerFile.Close()
	logger.Close()
}