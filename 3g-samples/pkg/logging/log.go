package logging

import (
	"3g-samples/pkg/file"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var (
	DefaultPrefix      = ""
	logger             *log.Logger
	DefaultCallerDepth = 2
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	logPrefix          = ""
)

type Level int

const (
	Levels = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func InitLogger() {
	filePath := GetLogFilePath()
	fileName := GetLogFileName()
	f, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	logger = log.New(f, DefaultPrefix, log.LstdFlags)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
}
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}
