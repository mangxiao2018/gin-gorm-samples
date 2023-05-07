package logging

import (
	"fmt"
	"time"
)

func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s", "log", time.Now().Format("20060102"), "log")
}

func GetLogFilePath() string {
	return fmt.Sprintf("%s%s.%s", "runtime/", "logs/")
}
