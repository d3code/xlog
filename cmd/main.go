package main

import "github.com/d3code/xlog"

func main() {

	xlog.EnableConsole(xlog.LevelTrace, xlog.CallerLong, "", true)

	xlog.Tracef("This is a trace message")
	xlog.Debugf("This is a debug message")
	xlog.Infof("This is an info message")
	xlog.Warnf("This is a warn message")
	xlog.Errorf("This is an error message")
	xlog.Fatalf("This is an error message")
}
