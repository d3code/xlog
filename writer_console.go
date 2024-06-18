package xlog

import (
    "bufio"
    "fmt"
    "github.com/d3code/xlog/color"
    "strings"
    "sync"
)

var mutexConsole = sync.Mutex{}

func consoleWriter(msg logItem) {
    mutexConsole.Lock()
    defer mutexConsole.Unlock()

    config := consoleConfig
    if msg.Level < GetLevel(config.Level) {
        return
    }

    var writer *bufio.Writer

    if msg.Level >= LevelError {
        writer = writerErrConsole
    } else {
        writer = writerOutConsole
    }

    var timestamp = msg.Timestamp.Format("2006-01-02 15:04:05")
    var level = msg.Level.name()
    var caller = msg.Caller
    var line = fmt.Sprintf("(%d)", msg.Line)
    var message = msg.Message

    if len(level) < 5 {
        level = fmt.Sprintf("%-5s", level)
    }

    if config.Caller == "short" {
        caller = caller[strings.LastIndex(caller, "/")+1:]
    }

    if len(caller) > 20 {
        caller = caller[len(caller)-20:]
    } else {
        caller = fmt.Sprintf("%-20s", caller)
    }

    if len(line) < 6 {
        line = fmt.Sprintf("%-6s", line)
    }

    if config.Color {
        timestamp = color.String(timestamp, "grey")
        caller = color.String(caller, "grey")
        line = color.String(line, "grey")

        switch msg.Level {
        case LevelTrace:
            level = color.String(level, "grey")
            message = color.String(message, "grey")
        case LevelDebug:
            level = color.String(level, "grey")
            message = color.String(message, "grey")
        case LevelInfo:
            level = color.String(level, "blue")
        case LevelWarn:
            level = color.String(level, "yellow")
            message = color.String(message, "yellow")
        case LevelError:
            level = color.String(level, "red")
            message = color.String(message, "red")
        case LevelFatal:
            level = color.String(level, "red")
            message = color.String(message, "red")
        }
    }

    var output string
    if config.Caller == "short" || config.Caller == "long" {
        output = fmt.Sprintf("%s %s %s %s  %s\n", timestamp, level, line, caller, message)
    } else {
        output = fmt.Sprintf("%s %s %s\n", timestamp, level, message)
    }

    writer.WriteString(output)
    writer.Flush()
}
