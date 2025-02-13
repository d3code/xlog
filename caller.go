package xlog

type Caller int

const (
	CallerShort Caller = iota
	CallerLong
	CallerNone
)
