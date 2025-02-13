package xlog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// Trace logs a message with level Trace
func Trace(msg string) {
	log(LevelTrace, msg)
}

// Tracef logs a formatted message with level Trace
func Tracef(format string, args ...interface{}) {
	log(LevelTrace, fmt.Sprintf(format, args...))
}

// Debug logs a message with level Debug
func Debug(msg string) {
	log(LevelDebug, msg)
}

// Debugf logs a formatted message with level Debug
func Debugf(format string, args ...interface{}) {
	log(LevelDebug, fmt.Sprintf(format, args...))
}

// Info logs a message with level Info
func Info(msg string) {
	log(LevelInfo, msg)
}

// Infof logs a formatted message with level Info
func Infof(format string, args ...interface{}) {
	log(LevelInfo, fmt.Sprintf(format, args...))
}

// Warn logs a message with level Warn
func Warn(msg string) {
	log(LevelWarn, msg)
}

// Warnf logs a formatted message with level Warn
func Warnf(format string, args ...interface{}) {
	log(LevelWarn, fmt.Sprintf(format, args...))
}

// Error logs a message with level Error
func Error(msg string) {
	log(LevelError, msg)
}

// Errorf logs a formatted message with level Error
func Errorf(format string, args ...interface{}) {
	log(LevelError, fmt.Sprintf(format, args...))
}

// Fatal logs a message with level Fatal and exits the program
func Fatal(msg string) {
	log(LevelFatal, msg)
}

// Fatalf logs a formatted message with level Fatal and exits the program
func Fatalf(format string, args ...interface{}) {
	log(LevelFatal, fmt.Sprintf(format, args...))
}

func Log(level Level, message string) {
	log(level, message)
}

// Log logs a message with the specified level
// The message is written to the console and/or file depending on the configuration
func log(level Level, message string) {
	msg := logItem{
		Level:     level,
		Message:   strings.TrimSpace(message),
		Timestamp: time.Now(),
	}

	_, caller, no, ok := runtime.Caller(2)
	if ok {
		msg.Line = no
		msg.Caller = caller
	}

	if configuration.Console.Enabled {
		consoleWriter(msg)
	}
	if configuration.File.Enabled {
		fileWriter(msg)
	}

	if level == LevelFatal {
		os.Exit(1)
	}
}

type logItem struct {
	Level     Level
	Timestamp time.Time
	Caller    string
	Line      int
	Message   string
}
