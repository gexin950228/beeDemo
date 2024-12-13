package utils

import (
	"fmt"
	"os"
	"time"
)

func LogToFile(loglevel, message string) {
	logFile := "D:\\beeDemo\\logs\\" + loglevel + ".log"
	fmt.Println(logFile)
	if logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, 0644); err != nil {
		panic("日志文件创建失败")
	} else {
		time := time.Now().Format("2006-01-02 15:04:05")
		logString := time + " " + loglevel + ":" + " " + message + "\n"
		logFile.WriteString(logString)
	}
}
