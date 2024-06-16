package xlog

import (
    "os"
    "time"
)

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

    time.Sleep(1 * time.Second)
    os.Exit(1)
}
