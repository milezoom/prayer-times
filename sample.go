package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, callerFile, _, _ := runtime.Caller(0)
	logPath := filepath.Dir(callerFile) + "logs"

	fmt.Print(logPath)
}
