package main

import "github.com/d3code/xlog"

func main() {

	xlog.EnableConsole(xlog.LevelTrace, xlog.CallerLong, "[TEST]", true)
	xlog.EnableFile("log.txt", xlog.LevelTrace, xlog.CallerShort)

	xlog.Tracef("This is a trace message")
	xlog.Debugf("This is a debug message")
	xlog.Infof("This is an info message")
	xlog.Warnf("This is a warn message")
	xlog.Successf("This is a success message")
	xlog.Errorf("This is an error message")
	xlog.Fatalf("This is a fatal message")
}
