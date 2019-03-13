package logger

import (
	"fmt"
	"os"
)

// Fatal error logger (red)
func Fatal(x interface{}) {
	prefixLog("FATAL", 31, x)
	os.Exit(1)
}

// Error logger (red)
func Error(x interface{}) {
	prefixLog("ERROR", 31, x)
}

// Warn logger (yellow)
func Warn(x interface{}) {
	prefixLog("WARN", 33, x)
}

// Success logger (green)
func Success(x interface{}) {
	prefixLog("SUCCESS", 32, x)
}

// Log something
func Log(x interface{}) {
	prefixLog("LOG", 0, x)
}

func prefixLog(prefix string, color int, x interface{}) {
	fmt.Printf("\x1b[%dm[%s]\x1b[0m %s\n", color, prefix, x)
}
