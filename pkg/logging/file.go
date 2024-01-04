package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// LogSavePath
	LogSavePath = "runtime/logs/"
	// LogSaveName
	LogSaveName = "log"
	// LogFileExt a
	LogFileExt = "log"
	// TimeFormat
	TimeFormat = "20060102"
)

// get log file path
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// get log file full path
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// open log file; if not exist, create dir then open
func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			mkDir()
		}
		if os.IsPermission(err) {
			log.Fatal("permission not allowed, the err:%v", err)
		}
	}
	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("fail to openfile: %v", err)
	}
	return handle
}

// mkdir
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+LogSavePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
