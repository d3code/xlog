package xlog

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelSuccess
	LevelWarn
	LevelError
	LevelFatal
)

var levelNames = map[Level]string{
	LevelTrace:   "TRACE",
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelSuccess: "SUCCESS",
	LevelWarn:    "WARN",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

func (l Level) name() string {
	return levelNames[l]
}
