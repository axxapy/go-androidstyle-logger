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

	SetLogLevel(level LogLevel, tags ...string)
	ResetLogLevel(tags ...string)
	SetFormatter(f Formatter)
	SetWriter(w io.Writer)

	IsLoggable(level LogLevel, tags ...string) bool

	WithTag(tag string) SimpleLogger
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
	}
}

func (l *defaultLogger) format(tag string, level LogLevel, msg ...interface{}) []byte {
	return l.formatter(l, tag, level, msg...)
}

func (l *defaultLogger) _print_error(tag string, level LogLevel, msg ...interface{}) {
	l.writer.Write(l.format(tag, level, msg...))
}

func (l *defaultLogger) _print_info(tag string, level LogLevel, msg ...interface{}) {
	l.writer.Write(l.format(tag, level, msg...))
}

func (l *defaultLogger) D(tag string, msg ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l._print_info(tag, DEBUG, msg)
	}
}

func (l *defaultLogger) Df(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		l._print_info(tag, DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) V(tag string, msg ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l._print_info(tag, VERBOSE, msg)
	}
}

func (l *defaultLogger) Vf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		l._print_info(tag, VERBOSE, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) E(tag string, msg ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l._print_error(tag, ERROR, msg)
	}
}

func (l *defaultLogger) Ef(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		l._print_error(tag, ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) W(tag string, msg ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l._print_info(tag, WARNING, msg)
	}
}

func (l *defaultLogger) Wf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		l._print_info(tag, WARNING, fmt.Sprintf(msg, args...))
	}
}

func (l *defaultLogger) I(tag string, msg ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l._print_info(tag, INFO, msg)
	}
}

func (l *defaultLogger) If(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		l._print_info(tag, INFO, fmt.Sprintf(msg, args...))
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

func (l *defaultLogger) SetLogLevel(level LogLevel, tags ...string) {
	if len(tags) < 1 {
		l.log_level = level
	} else {
		for _, tag := range tags {
			l.log_level_tag[tag] = level
		}
	}
}

func (l *defaultLogger) ResetLogLevel(tags ...string) {
	if len(tags) < 1 {
		l.log_level = LOG_LEVEL_DEFAULT
	} else {
		for _, tag := range tags {
			delete(l.log_level_tag, tag)
		}
	}
}

func (l *defaultLogger) SetFormatter(f Formatter) {
	l.formatter = f
}

func (l *defaultLogger) SetWriter(w io.Writer) {
	l.writer = w
}

func (l *defaultLogger) IsLoggable(level LogLevel, tags ...string) bool {
	if len(tags) < 1 {
		return l.log_level&level != 0
	}

	for _, tag := range tags {
		if tag_level, exists := l.log_level_tag[tag]; exists && tag_level&level != 0 {
			return true
		}
	}
	return false
}
