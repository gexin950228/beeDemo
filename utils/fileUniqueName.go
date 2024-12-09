package utils

import (
	"strings"
	"time"
)

func UniqueName(filename string) string {
	name, format := strings.Split(filename, ".")[0], strings.Split(filename, ".")[1]
	now := time.Now()
	timeString := now.Format("20060102150405")
	uniqueName := name + "-" + timeString + "." + format
	return uniqueName
}
