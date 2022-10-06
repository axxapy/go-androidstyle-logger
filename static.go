package l

import (
	"io"
	"strconv"
)

type LogLevel uint8

const (
	DEBUG LogLevel = 1 << iota
	ERROR
	INFO
	VERBOSE
	WARNING
	WTF
)

const (
	ALL = DEBUG ^ ERROR ^ INFO ^ VERBOSE ^ WARNING ^ WTF

	LOG_LEVEL_DEFAULT = WARNING ^ ERROR ^ INFO
)

var (
	logger     = New()
	levelNames = map[LogLevel]string{
		DEBUG:   "D",
		ERROR:   "E",
		INFO:    "I",
		VERBOSE: "V",
		WARNING: "W",
		WTF:     "WTF",
	}
)

type Flag uint8

const (
	FLAG_FILE_FULL_PATH Flag = 1 << iota
	FLAG_FILE_ONLY_NAME
	FLAG_FILE_WITH_PACKAGE
	FLAG_NO_FILENAME
)

func New() Logger {
	logger := &taggedLogger{
		baseLogger: newBaseLogger(),
	}
	logger.baseLogger.callerDeep++
	return logger
}

func Default() Logger {
	return logger
}

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

func WithTag(tag string) Logger {
	return logger.WithTag(tag)
}

func Check(level LogLevel) Logger {
	if logger.IsLoggable(level, "") {
		return logger
	}
	return nil
}

func CheckWithTag(level LogLevel, tag string) Logger {
	if logger.IsLoggable(level, tag) {
		return logger.WithTag(tag)
	}
	return nil
}

func D(tag string, msg ...interface{}) {
	logger.WithTag(tag).D(msg...)
}

func Df(tag string, msg string, args ...interface{}) {
	logger.WithTag(tag).Df(msg, args...)
}

func V(tag string, msg ...interface{}) {
	logger.WithTag(tag).V(msg...)
}

func Vf(tag string, msg string, args ...interface{}) {
	logger.WithTag(tag).Vf(msg, args...)
}

func E(tag string, msg ...interface{}) {
	logger.WithTag(tag).E(msg...)
}

func Ef(tag string, msg string, args ...interface{}) {
	logger.WithTag(tag).Ef(msg, args...)
}

func W(tag string, msg ...interface{}) {
	logger.WithTag(tag).W(msg...)
}

func Wf(tag string, msg string, args ...interface{}) {
	logger.WithTag(tag).Wf(msg, args...)
}

func I(tag string, msg ...interface{}) {
	logger.WithTag(tag).I(msg...)
}

func If(tag string, msg string, args ...interface{}) {
	logger.WithTag(tag).If(msg, args...)
}

func Fatal(tag string, err error) {
	logger.WithTag(tag).Fatal(err)
}
