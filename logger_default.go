package l

import (
	"fmt"
	"io"
	"os"
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

	SetLogLevel(level LogLevel, tags ...string) Logger
	ResetLogLevel(tags ...string) Logger
	SetFormatter(f Formatter) Logger
	SetWriter(w io.Writer) Logger

	IsLoggable(level LogLevel, tag string) bool

	WithTag(tag string) SimpleLogger

	Check(level LogLevel) Logger
	CheckWithTag(level LogLevel, tag string) SimpleLogger
}

type defaultLogger struct {
	log_level     LogLevel
	log_level_tag map[string]LogLevel
	formatter     Formatter
	writer        io.Writer
}

func New() Logger {
	return &defaultLogger{
		log_level:     LOG_LEVEL_DEFAULT,
		log_level_tag: make(map[string]LogLevel),
		formatter:     DefaultFormatter,
		writer:        os.Stderr,
	}
}

func (l *defaultLogger) format(tag string, level LogLevel, msg ...interface{}) []byte {
	return l.formatter(l, tag, level, msg...)
}

func (l *defaultLogger) print(tag string, level LogLevel, msg ...interface{}) {
	l.writer.Write(l.format(tag, level, msg...))
}

func (l *defaultLogger) D(tag string, msg ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l.print(tag, DEBUG, msg...)
	}
}

func (l *defaultLogger) Df(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l.print(tag, DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) V(tag string, msg ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l.print(tag, VERBOSE, msg...)
	}
}

func (l *defaultLogger) Vf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l.print(tag, VERBOSE, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) E(tag string, msg ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l.print(tag, ERROR, msg...)
	}
}

func (l *defaultLogger) Ef(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l.print(tag, ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) W(tag string, msg ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l.print(tag, WARNING, msg...)
	}
}

func (l *defaultLogger) Wf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l.print(tag, WARNING, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) I(tag string, msg ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l.print(tag, INFO, msg...)
	}
}

func (l *defaultLogger) If(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l.print(tag, INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) Fatal(tag string, err error) {
	if err != nil {
		E(tag, err)
		os.Exit(1)
	}
}

func (l *defaultLogger) WithTag(tag string) SimpleLogger {
	return &simpleLogger{
		logger: l,
		tag:    tag,
	}
}

func (l *defaultLogger) SetLogLevel(level LogLevel, tags ...string) Logger {
	if len(tags) < 1 {
		l.log_level = level
	} else {
		for _, tag := range tags {
			l.log_level_tag[tag] = level
		}
	}
	return l
}

func (l *defaultLogger) ResetLogLevel(tags ...string) Logger {
	if len(tags) < 1 {
		l.log_level = LOG_LEVEL_DEFAULT
	} else {
		for _, tag := range tags {
			delete(l.log_level_tag, tag)
		}
	}
	return l
}

func (l *defaultLogger) SetFormatter(f Formatter) Logger {
	l.formatter = f
	return l
}

func (l *defaultLogger) SetWriter(w io.Writer) Logger {
	l.writer = w
	return l
}

func (l *defaultLogger) IsLoggable(level LogLevel, tag string) bool {
	if tag != "" {
		if tagLevel, exists := l.log_level_tag[tag]; exists {
			return tagLevel&level != 0
		}
	}

	return l.log_level&level != 0
}

func (l *defaultLogger) Check(level LogLevel) Logger {
	if l.IsLoggable(level, "") {
		return l
	}
	return nil
}

func (l *defaultLogger) CheckWithTag(level LogLevel, tag string) SimpleLogger {
	if l.IsLoggable(level, tag) {
		return l.WithTag(tag)
	}
	return nil
}
