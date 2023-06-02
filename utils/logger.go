package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var staclogger *log.Logger
var execlogger *log.Logger
var stacOnce sync.Once
var execOnce sync.Once
var logFiles []*os.File

func GetStacLogger() *log.Logger {
	stacOnce.Do(func() {
		staclogger = createStacLogger(Config.StacLog)
	})
	return staclogger
}

func GetExecLogger() *log.Logger {
	execOnce.Do(func() {
		execlogger = createExecLogger(Config.ExecLog)
	})
	return execlogger
}

func createStacLogger(fname string) *log.Logger {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	logFiles = append(logFiles, file)
	return log.New(file, "stac: ", log.Lshortfile|log.Ldate|log.Ltime)
}

func createExecLogger(fname string) *log.Logger {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	logFiles = append(logFiles, file)
	return log.New(file, "exec: ", log.Lshortfile|log.Ldate|log.Ltime)
}

func CloseAllLogFiles() {
	for _, f := range logFiles {
		f.Sync()
		f.Close()
	}
}
