package l

import (
	"fmt"
	"strconv"
)

type LogLevel uint

const (
	DEBUG   = LogLevel(1)
	ERROR   = LogLevel(2)
	INFO    = LogLevel(4)
	VERBOSE = LogLevel(8)
	WARNING = LogLevel(16)
	WTF     = LogLevel(32)
	ALL     = DEBUG ^ ERROR ^ INFO ^ VERBOSE ^ WARNING ^ WTF

	LOG_LEVEL_DEFAULT = WARNING ^ ERROR ^ INFO
)

var (
	//logger    = log.New()
	logger     = New()
	levelNames = map[LogLevel]string{
		DEBUG:   "D",
		ERROR:   "E",
		INFO:    "I",
		VERBOSE: "V",
		WTF:     "WTF",
	}
)

type Logger interface {
	D(tag string, msg ...interface{})
	Df(tag string, msg string, args ...interface{})
	V(tag string, msg ...interface{})
	Vf(tag string, msg string, args ...interface{})
	E(tag string, msg ...interface{})
	Ef(tag string, msg string, args ...interface{})
	W(tag string, msg ...interface{})
	Wf(tag string, msg string, args ...interface{})
	I(tag string, msg ...interface{})
	If(tag string, msg string, args ...interface{})

	Fatal(tag string, err error)

	SetLogLevel(level LogLevel, tags ...string)
	ResetLogLevel(tags ...string)
	SetFormatter(f Formatter)

	IsLoggable(level LogLevel, tags ...string) bool

	WithTag(tag string) SimpleLogger
}

type SimpleLogger interface {
	D(msg ...interface{})
	Df(msg string, args ...interface{})
	V(msg ...interface{})
	Vf(msg string, args ...interface{})
	E(msg ...interface{})
	Ef(msg string, args ...interface{})
	W(msg ...interface{})
	Wf(msg string, args ...interface{})
	I(msg ...interface{})
	If(msg string, args ...interface{})

	Fatal(err error)
}

func GetLevelName(level LogLevel) string {
	if name, exists := levelNames[level]; exists {
		return name
	}
	return strconv.Itoa(int(level))
}

type Formatter func(tag string, level LogLevel, msg ...interface{}) string

func defaultFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) string {
	return "[" + GetLevelName(level) + "] [" + tag + "] " + fmt.Sprintln(msg...)
}

func Check(level LogLevel, tags ...string) Logger {
	if logger.IsLoggable(level, tags...) {
		return logger
	}
	return nil
}

func CheckLevel(level LogLevel) Logger {
	if logger.IsLoggable(level) {
		return logger
	}
	return nil
}

func CheckTagLevel(tag string, level LogLevel) SimpleLogger {
	if logger.IsLoggable(level, tag) {
		return logger.WithTag(tag)
	}
	return nil
}
