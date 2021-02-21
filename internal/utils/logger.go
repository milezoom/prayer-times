package utils

import (
	"os"
	"path/filepath"
	"time"
)

//LogWriter give file writer to write log file
func LogWriter() *os.File {
	ex, _ := os.Executable()
	logPath := filepath.Dir(ex) + "/logs"

	os.Mkdir(logPath, 0744)

	currentDate := time.Now().Local().Format("2006-01-02")

	file, err := os.OpenFile(logPath+"/"+currentDate+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return file
}
