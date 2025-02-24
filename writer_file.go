package xlog

import (
	"bufio"
	"fmt"
)

func fileWriter(msg logItem) {
	var writer *bufio.Writer
	if msg.Level == LevelError || msg.Level == LevelFatal {
		writer = writerErrFile
	} else {
		writer = writerOutFile
	}

	var timestamp = msg.Timestamp.Format("2006-01-02T15:04:05")
	var line = fmt.Sprintf(":%d", msg.Line)
	var level = msg.Level.name()
	var message = msg.Message

	var caller = formatCaller(msg.Caller, configuration.File.Caller)

	var output string
	if configuration.File.Caller == CallerNone {
		output = fmt.Sprintf("%s [%s] %s\n", timestamp, level, message)
	} else {
		output = fmt.Sprintf("%s [%s] %s%s %s\n", timestamp, level, caller, line, message)
	}

	_, _ = writer.WriteString(output)
	_ = writer.Flush()
}
