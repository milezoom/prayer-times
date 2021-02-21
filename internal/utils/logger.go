package utils

import (
	"os"
	"path/filepath"
	"time"
)

//LogWriter give file writer to write log file
func LogWriter(prefix string) *os.File {
	ex, _ := os.Executable()
	logPath := filepath.Dir(ex) + "/logs"

	os.Mkdir(logPath, 0744)

	currentDate := time.Now().Local().Format("2006-01-02")

	if prefix == "" {
		prefix = "debug"
	}

	filename := logPath + "/" + prefix + "-" + currentDate + ".log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return file
}
