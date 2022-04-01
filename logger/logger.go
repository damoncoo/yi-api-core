package logger

import (
	"log"
	"os"
)

// DebugLog DebugLog
type DebugLog bool

// Debug Debug
var Debug DebugLog

var dbgLog = log.New(os.Stdout, "[DEBUG] ", log.Ltime)

// Printf Printf
func (d DebugLog) Printf(format string, args ...interface{}) {
	if d {
		dbgLog.Printf(format, args...)
	}
}

// Println Println
func (d DebugLog) Println(args ...interface{}) {
	if d {
		dbgLog.Println(args...)
	}
}
