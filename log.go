package xlog

import (
    "fmt"
    "runtime"
    "strings"
    "time"
)

type Level int

const (
    LevelDebug Level = iota
    LevelInfo
    LevelWarn
    LevelError
    LevelFatal
)

func (l Level) name() string {
    switch l {
    case LevelDebug:
        return "DEBUG"
    case LevelInfo:
        return "INFO"
    case LevelWarn:
        return "WARN"
    case LevelError:
        return "ERROR"
    case LevelFatal:
        return "FATAL"
    default:
        return ""
    }
}

func GetLevel(level string) Level {
    switch strings.ToUpper(level) {
    case "DEBUG":
        return LevelDebug
    case "INFO":
        return LevelInfo
    case "WARN":
        return LevelWarn
    case "ERROR":
        return LevelError
    case "FATAL":
        return LevelFatal
    default:
        return LevelInfo
    }
}

func Log(level Level, message ...any) {
    var str = make([]string, len(message))
    for i, v := range message {
        str[i] = fmt.Sprintf("%v", v)
    }

    msg := logItem{
        Level:     level,
        Message:   strings.Join(str, " "),
        Timestamp: time.Now(),
    }

    _, caller, no, ok := runtime.Caller(2)
    if ok {
        msg.Line = no
        msg.Caller = caller
    }

    if consoleConfig != nil && consoleConfig.Enabled {
        consoleWriter(msg)
    }

    if len(logChannels) > 0 {
        for _, channel := range logChannels {
            channel <- msg
        }
    }
}

type logItem struct {
    Level     Level
    Timestamp time.Time
    Caller    string
    Line      int
    Message   string
}
