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
	logFile := filepath.Join(getwd, "logs", loglevel+".log")
	if logFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE, 0644); err != nil {
		panic("日志文件创建失败")
	} else {
		time := time.Now().Format("2006-01-02 15:04:05")
		logString := time + " " + loglevel + ":" + " " + message + "\n"
		logFile.WriteString(logString)
	}
}
