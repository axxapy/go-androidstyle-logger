package l

import (
	"fmt"
	"log"
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
	logger = New()
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

func _print_error(tag string, level string, msg ...interface{}) {
	log.Print("[" + level + "] [" + tag + "] " + fmt.Sprintln(msg...))
}

func _print_info(tag string, level string, msg ...interface{}) {
	log.Output(2, "["+level+"] ["+tag+"] "+fmt.Sprintln(msg...))
}

func _format(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
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
