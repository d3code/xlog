package xlog

import (
	"bufio"
	"os"
)

var (
	configuration *config

	consoleOut    = bufio.NewWriter(os.Stdout)
	writerOutFile *bufio.Writer
	writerErrFile *bufio.Writer

	defaultConsoleConfig = consoleConfig{
		Enabled: true,
		Color:   true,
		Level:   LevelDebug,
		Caller:  CallerShort,
	}
)

func init() {
	configuration = &config{
		Console: defaultConsoleConfig,
		File: fileConfig{
			Enabled: false,
		},
	}
}

func EnableConsoleDefaults() {
	configuration.Console = defaultConsoleConfig
}

func EnableConsole(level Level, caller Caller, prefix string, color bool) {
	configuration.Console.Enabled = true
	configuration.Console.Level = level
	configuration.Console.Color = color
	configuration.Console.Caller = caller
	configuration.Console.Prefix = prefix
}

func DisableConsole() {
	configuration.Console.Enabled = false
}

func EnableFile(path string, level Level, caller Caller) {
	configuration.File.Enabled = true
	configuration.File.Path = path
	configuration.File.Level = level
	configuration.File.Caller = caller

	file, err := os.OpenFile(configuration.File.Path+"application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		DisableFile()
		Fatal(err.Error())
	}
	writerOutFile = bufio.NewWriter(file)

	fileError, errError := os.OpenFile(configuration.File.Path+"application-err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errError != nil {
		DisableFile()
		Fatal(errError.Error())
	}
	writerErrFile = bufio.NewWriter(fileError)
}

func DisableFile() {
	configuration.File.Enabled = false
}

type consoleConfig struct {
	Enabled bool
	Color   bool
	Level   Level
	Caller  Caller
	Prefix  string
}

type fileConfig struct {
	Enabled bool
	Level   Level
	Caller  Caller
	Path    string
}

type config struct {
	Console consoleConfig
	File    fileConfig
}
