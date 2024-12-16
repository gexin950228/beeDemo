package utils

import (
	"os"
	"path/filepath"
	"time"
)

func LogToFile(loglevel, message string) {
<<<<<<< HEAD
	getwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	logFile := filepath.Join(getwd, "logs", loglevel+".log")
	if logFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE, 0644); err != nil {
=======
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logFile := filepath.Join(path, "logs", loglevel+".log")
	if logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, 0644); err != nil {
>>>>>>> 8b8e7ec9f9f25040bdeebafc323e59cc42cb1e31
		panic("日志文件创建失败")
	} else {
		time := time.Now().Format("2006-01-02 15:04:05")
		logString := time + " " + loglevel + ":" + " " + message + "\n"
		logFile.WriteString(logString)
	}
}
