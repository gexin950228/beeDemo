package utils

import (
	"os"
	"path/filepath"
	"time"
)

func LogToFile(loglevel, message string) {
	getwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	logFileName := filepath.Join(getwd, "logs", loglevel+".log")
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	} else {
		time := time.Now().Format("2006-01-02 15:04:05")
		logString := time + " " + loglevel + ":" + " " + message + "\n"
		logFile.WriteString(logString)
	}
}
