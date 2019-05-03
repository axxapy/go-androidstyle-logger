package l

import (
	"io"
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

func GetLevelName(level LogLevel) string {
	if name, exists := levelNames[level]; exists {
		return name
	}
	return strconv.Itoa(int(level))
}

func SetLogLevel(level LogLevel, tags ...string) {
	logger.SetLogLevel(level, tags...)
}

func ResetLogLevel(tags ...string) {
	logger.ResetLogLevel(tags...)
}

func SetFormatter(f Formatter) {
	logger.SetFormatter(f)
}

func SetWriter(w io.Writer) {
	logger.SetWriter(w)
}

func WithTag(tag string) SimpleLogger {
	return logger.WithTag(tag)
}

func Check(level LogLevel, tag string) Logger {
	if logger.IsLoggable(level, tag) {
		return logger
	}
	return nil
}

func CheckLevel(level LogLevel) Logger {
	if logger.IsLoggable(level, "") {
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

func D(tag string, msg ...interface{}) {
	logger.D(tag, msg...)
}

func Df(tag string, msg string, args ...interface{}) {
	logger.Df(tag, msg, args...)
}

func V(tag string, msg ...interface{}) {
	logger.V(tag, msg...)
}

func Vf(tag string, msg string, args ...interface{}) {
	logger.Vf(tag, msg, args...)
}

func E(tag string, msg ...interface{}) {
	logger.E(tag, msg...)
}

func Ef(tag string, msg string, args ...interface{}) {
	logger.Ef(tag, msg, args...)
}

func W(tag string, msg ...interface{}) {
	logger.W(tag, msg...)
}

func Wf(tag string, msg string, args ...interface{}) {
	logger.Wf(tag, msg, args...)
}

func I(tag string, msg ...interface{}) {
	logger.I(tag, msg...)
}

func If(tag string, msg string, args ...interface{}) {
	logger.If(tag, msg, args...)
}

func Fatal(tag string, err error) {
	logger.Fatal(tag, err)
}
