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
	line := fmt.Sprintf(":%d", msg.Line)

	// Color timestamp and caller
	if consoleConfiguration.Color {
		timestamp = color.String(timestamp, "grey")
		caller = color.String(caller, "grey")
		line = color.String(line, "grey")
	}

	// Format message
	message := formatMessage(msg.Message, msg.Level, consoleConfiguration.Color)

	// Write to console
	output := formatOutput(consoleConfiguration.Caller, timestamp, level, line, caller, message)
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
	index := secondToLastIndex(caller, "/")
	caller = caller[index+1:]

	return caller
}

func secondToLastIndex(s, sep string) int {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex == -1 {
		return -1
	}
	secondToLastIndex := strings.LastIndex(s[:lastIndex], sep)
	return secondToLastIndex
}

func formatMessage(message string, level Level, useColor bool) string {
	if useColor {
		return colorizeMessage(message, level)
	}
	return message
}

func formatOutput(callerConfig Caller, timestamp, level string, line string, caller, message string) string {
	if callerConfig == CallerNone {
		return fmt.Sprintf("%s  %s  %s\n", timestamp, level, message)
	}
	// Calculate the length of the caller string and line number with stepped function
	callerLength := len(caller) + len(line)
	steppedLength := ((callerLength/6)+1)*6 + 2

	// Combine caller and line number into a set length
	callerLine := fmt.Sprintf("%s%s", caller, line)
	callerLine = fmt.Sprintf("%-*s", steppedLength, callerLine)

	return fmt.Sprintf("%s  %s  %s  %s\n", timestamp, level, callerLine, message)
}

func colorizeLevel(level string, logLevel Level) string {
	colorMap := map[Level]string{
		LevelTrace:   "grey",
		LevelDebug:   "grey",
		LevelInfo:    "blue",
		LevelSuccess: "green",
		LevelWarn:    "yellow",
		LevelError:   "red",
		LevelFatal:   "red",
	}
	if messageColor, ok := colorMap[logLevel]; ok {
		return color.String(level, messageColor)
	}
	return level
}

func colorizeMessage(message string, logLevel Level) string {
	colorMap := map[Level]string{
		LevelTrace:   "grey",
		LevelDebug:   "grey",
		LevelSuccess: "green",
		LevelWarn:    "yellow",
		LevelError:   "red",
		LevelFatal:   "red",
	}
	if messageColor, ok := colorMap[logLevel]; ok {
		return color.String(message, messageColor)
	}
	return message
}
