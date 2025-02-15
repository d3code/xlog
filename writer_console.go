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

	// Check if message level is enabled
	consoleConfiguration := configuration.Console
	if msg.Level < consoleConfiguration.Level {
		return
	}

	// Get writer
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(consoleOut)

	// Format timestamp, level, caller, line
	timestamp := msg.Timestamp.Format("2006-01-02 15:04:05")
	level := formatLevel(msg.Level, consoleConfiguration.Color)
	caller := formatCaller(msg.Caller, consoleConfiguration.Caller)

	// Format message
	message := formatMessage(msg.Message, msg.Level, consoleConfiguration.Color)

	// Write to console
	output := formatOutput(consoleConfiguration.Caller, timestamp, level, msg.Line, caller, message)
	if consoleConfiguration.Color {
		if msg.Level == LevelTrace {
			output = color.String(output, "grey")
		}
		if msg.Level == LevelFatal {
			output = color.String(output, "red")
		}
	}

	_, _ = consoleOut.WriteString(output)
}

func formatLevel(level Level, useColor bool) string {
	lvl := level.name()
	if len(lvl) < 8 {
		lvl = fmt.Sprintf("%-8s", lvl)
	}
	if useColor {
		return colorizeLevel(lvl, level)
	}
	return lvl
}

func formatCaller(caller string, callerConfig Caller) string {
	if callerConfig == CallerShort {
		caller = caller[strings.LastIndex(caller, "/")+1:]
	}

	return caller
}

func formatMessage(message string, level Level, useColor bool) string {
	if useColor {
		return colorizeMessage(message, level)
	}
	return message
}

func formatOutput(callerConfig Caller, timestamp, level string, line int, caller, message string) string {
	if callerConfig == CallerNone {
		return fmt.Sprintf("%s  %s  %s\n", timestamp, level, message)
	}
	// Combine caller and line number into a set length
	callerLine := fmt.Sprintf("%s (%d)", caller, line)
	if len(callerLine) > 24 {
		callerLine = callerLine[len(callerLine)-24:]
	} else {
		callerLine = fmt.Sprintf("%-24s", callerLine)
	}
	return fmt.Sprintf("%s  %s  %s  %s\n", timestamp, level, callerLine, message)
}

func colorizeLevel(level string, logLevel Level) string {
	colorMap := map[Level]string{
		LevelDebug:   "grey",
		LevelInfo:    "blue",
		LevelSuccess: "green",
		LevelWarn:    "yellow",
		LevelError:   "red",
	}
	if messageColor, ok := colorMap[logLevel]; ok {
		return color.String(level, messageColor)
	}
	return level
}

func colorizeMessage(message string, logLevel Level) string {
	colorMap := map[Level]string{
		LevelDebug:   "grey",
		LevelSuccess: "green",
		LevelWarn:    "yellow",
		LevelError:   "red",
	}
	if messageColor, ok := colorMap[logLevel]; ok {
		return color.String(message, messageColor)
	}
	return message
}
