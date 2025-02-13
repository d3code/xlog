package xlog

import (
	"bufio"
	"bytes"
	"testing"
	"time"
)

func TestLogLevels(t *testing.T) {
	tests := []struct {
		level   Level
		message string
	}{
		{LevelTrace, "This is a trace message"},
		{LevelDebug, "This is a debug message"},
		{LevelInfo, "This is an info message"},
		{LevelWarn, "This is a warn message"},
		{LevelError, "This is an error message"},
		{LevelFatal, "This is a fatal message"},
	}

	EnableConsole(LevelTrace, CallerShort, true)

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			var buf bytes.Buffer
			consoleOut = bufio.NewWriter(&buf)

			Log(tt.level, tt.message)

			if !bytes.Contains(buf.Bytes(), []byte(tt.message)) {
				t.Errorf("expected log message to contain %q, got %q", tt.message, buf.String())
			}
		})
	}
}

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	consoleOut = bufio.NewWriter(&buf)

	Log(LevelInfo, "Test message")

	if !bytes.Contains(buf.Bytes(), []byte("Test message")) {
		t.Errorf("expected log message to contain %q, got %q", "Test message", buf.String())
	}
}

func TestConsoleWriter(t *testing.T) {
	var buf bytes.Buffer
	consoleOut = bufio.NewWriter(&buf)

	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")

	msg := logItem{
		Level:     LevelInfo,
		Message:   "Test console writer",
		Timestamp: now,
		Caller:    "test.go",
		Line:      42,
	}

	consoleWriter(msg)

	b := []byte(formattedTime + " \x1b[34mINFO \x1b[0m (42    ) test.go               Test console writer\n")
	if !bytes.Equal(buf.Bytes(), b) {
		t.Errorf("expected log message to be %q, got %q", string(b), buf.String())
	}
}
