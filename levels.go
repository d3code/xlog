package xlog

import (
    "os"
)

func Trace(msg ...any) {
    Log(LevelTrace, msg...)
}
func Debug(msg ...any) {
    Log(LevelDebug, msg...)
}

func Info(msg ...any) {
    Log(LevelInfo, msg...)
}

func Warn(msg ...any) {
    Log(LevelWarn, msg...)
}

func Error(msg ...any) {
    Log(LevelError, msg...)
}
func Fatal(msg ...any) {
    Log(LevelFatal, msg...)
    os.Exit(1)
}
