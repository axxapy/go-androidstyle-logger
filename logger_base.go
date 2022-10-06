package l

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type baseLoggerInterface interface {
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

	SetLogLevel(level LogLevel, tags ...string) baseLoggerInterface
	ResetLogLevel(tags ...string) baseLoggerInterface
	SetFormatter(f Formatter) baseLoggerInterface
	SetWriter(w io.Writer) baseLoggerInterface

	IsLoggable(level LogLevel, tag string) bool

	WithTag(tag string) Logger

	Check(level LogLevel) baseLoggerInterface
	CheckWithTag(level LogLevel, tag string) Logger

	Flags() Flag
	SetFlags(flags Flag) baseLoggerInterface
}

type baseLogger struct {
	logLevel       LogLevel
	logLevelPerTag map[string]LogLevel
	formatter      Formatter
	writer         io.Writer
	callerDeep     int
	uint8          int
	flags          Flag
}

func newBaseLogger() *baseLogger {
	return &baseLogger{
		logLevel:       LOG_LEVEL_DEFAULT,
		logLevelPerTag: make(map[string]LogLevel),
		formatter:      DefaultFormatter,
		writer:         os.Stderr,
		callerDeep:     3,
		flags:          FLAG_FILE_ONLY_NAME,
	}
}

func (l *baseLogger) format(tag string, level LogLevel, file string, line int, msg ...interface{}) []byte {
	return l.formatter(tag, level, file, line, msg...)
}

func (l *baseLogger) caller() (string, int) {
	if l.flags&FLAG_NO_FILENAME != 0 {
		return "", 0
	}

	pc, file, line, ok := runtime.Caller(l.callerDeep)
	if !ok {
		return "", 0
	}

	if l.flags&FLAG_FILE_FULL_PATH != 0 {
		return file, line
	}

	file = filepath.Base(file)

	if l.flags&FLAG_FILE_ONLY_NAME != 0 {
		return file, line
	}

	if l.flags&FLAG_FILE_WITH_PACKAGE != 0 {
		parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		pl := len(parts)
		packageName := ""
		if parts[pl-2][0] == '(' {
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}
		return packageName + "/" + file, line
	}

	return "", 0
}

func (l *baseLogger) print(tag string, level LogLevel, msg ...interface{}) {
	file, line := l.caller()
	l.writer.Write(l.format(tag, level, file, line, msg...))
}

func (l *baseLogger) D(tag string, msg ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l.print(tag, DEBUG, msg...)
	}
}

func (l *baseLogger) Df(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l.print(tag, DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *baseLogger) V(tag string, msg ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l.print(tag, VERBOSE, msg...)
	}
}

func (l *baseLogger) Vf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l.print(tag, VERBOSE, fmt.Sprintf(msg, args...))
	}
}

func (l *baseLogger) E(tag string, msg ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l.print(tag, ERROR, msg...)
	}
}

func (l *baseLogger) Ef(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l.print(tag, ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *baseLogger) W(tag string, msg ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l.print(tag, WARNING, msg...)
	}
}

func (l *baseLogger) Wf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l.print(tag, WARNING, fmt.Sprintf(msg, args...))
	}
}

func (l *baseLogger) I(tag string, msg ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l.print(tag, INFO, msg...)
	}
}

func (l *baseLogger) If(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l.print(tag, INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *baseLogger) Fatal(tag string, err error) {
	if err != nil {
		E(tag, err)
		os.Exit(1)
	}
}

func (l *baseLogger) WithTag(tag string) Logger {
	return &taggedLogger{
		baseLogger: l,
		tag:        tag,
	}
}

func (l *baseLogger) SetLogLevel(level LogLevel, tags ...string) baseLoggerInterface {
	if len(tags) < 1 {
		l.logLevel = level
	} else {
		for _, tag := range tags {
			l.logLevelPerTag[tag] = level
		}
	}
	return l
}

func (l *baseLogger) ResetLogLevel(tags ...string) baseLoggerInterface {
	if len(tags) < 1 {
		l.logLevel = LOG_LEVEL_DEFAULT
	} else {
		for _, tag := range tags {
			delete(l.logLevelPerTag, tag)
		}
	}
	return l
}

func (l *baseLogger) SetFormatter(f Formatter) baseLoggerInterface {
	l.formatter = f
	return l
}

func (l *baseLogger) SetWriter(w io.Writer) baseLoggerInterface {
	l.writer = w
	return l
}

func (l *baseLogger) IsLoggable(level LogLevel, tag string) bool {
	if tag != "" {
		if tagLevel, exists := l.logLevelPerTag[tag]; exists {
			return tagLevel&level != 0
		}
	}

	return l.logLevel&level != 0
}

func (l *baseLogger) Check(level LogLevel) baseLoggerInterface {
	if l.IsLoggable(level, "") {
		return l
	}
	return nil
}

func (l *baseLogger) CheckWithTag(level LogLevel, tag string) Logger {
	if l.IsLoggable(level, tag) {
		return l.WithTag(tag)
	}
	return nil
}

func (l *baseLogger) Flags() Flag {
	return l.flags
}

func (l *baseLogger) SetFlags(flags Flag) baseLoggerInterface {
	l.flags = flags
	return l
}
