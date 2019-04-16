package l

import "os"

type defaultLogger struct {
	log_level     LogLevel
	log_level_tag map[string]LogLevel
}

func New() Logger {
	return &defaultLogger{
		log_level:     LOG_LEVEL_DEFAULT,
		log_level_tag: make(map[string]LogLevel),
	}
}

func (l *defaultLogger) D(tag string, msg ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		_print_info(tag, "D", msg)
	}
}

func (l *defaultLogger) Df(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(DEBUG, tag) {
		_print_info(tag, "D", _format(msg, args...))
	}
}

func (l *defaultLogger) V(tag string, msg ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		_print_info(tag, "V", msg)
	}
}

func (l *defaultLogger) Vf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(VERBOSE, tag) {
		_print_info(tag, "V", _format(msg, args...))
	}
}

func (l *defaultLogger) E(tag string, msg ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		_print_error(tag, "E", msg)
	}
}

func (l *defaultLogger) Ef(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(ERROR, tag) {
		_print_error(tag, "E", _format(msg, args...))
	}
}

func (l *defaultLogger) W(tag string, msg ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		_print_info(tag, "W", msg)
	}
}

func (l *defaultLogger) Wf(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(WARNING, tag) {
		_print_info(tag, "W", _format(msg, args...))
	}
}

func (l *defaultLogger) I(tag string, msg ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		_print_info(tag, "I", msg)
	}
}

func (l *defaultLogger) If(tag string, msg string, args ...interface{}) {
	if l.IsLoggable(INFO, tag) {
		_print_info(tag, "I", _format(msg, args...))
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

func (l *defaultLogger) IsLoggable(level LogLevel, tags ...string) bool {
	for _, tag := range tags {
		if tag_level, exists := l.log_level_tag[tag]; exists && tag_level&level != 0 {
			return true
		}
	}
	return l.log_level&level != 0
}
