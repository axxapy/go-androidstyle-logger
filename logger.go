package l

import "os"

type defaultLogger struct{}

func (l *defaultLogger) D(tag string, msg ...interface{}) {
	if is_printable(tag, DEBUG) {
		_print_info(tag, "D", msg)
	}
}

func (l *defaultLogger) Df(tag string, msg string, args ...interface{}) {
	if is_printable(tag, DEBUG) {
		_print_info(tag, "D", _format(msg, args...))
	}
}

func (l *defaultLogger) V(tag string, msg ...interface{}) {
	if is_printable(tag, VERBOSE) {
		_print_info(tag, "V", msg)
	}
}

func (l *defaultLogger) Vf(tag string, msg string, args ...interface{}) {
	if is_printable(tag, VERBOSE) {
		_print_info(tag, "V", _format(msg, args...))
	}
}

func (l *defaultLogger) E(tag string, msg ...interface{}) {
	if is_printable(tag, ERROR) {
		_print_error(tag, "E", msg)
	}
}

func (l *defaultLogger) Ef(tag string, msg string, args ...interface{}) {
	if is_printable(tag, ERROR) {
		_print_error(tag, "E", _format(msg, args...))
	}
}

func (l *defaultLogger) W(tag string, msg ...interface{}) {
	if is_printable(tag, WARNING) {
		_print_info(tag, "W", msg)
	}
}

func (l *defaultLogger) Wf(tag string, msg string, args ...interface{}) {
	if is_printable(tag, WARNING) {
		_print_info(tag, "W", _format(msg, args...))
	}
}

func (l *defaultLogger) I(tag string, msg ...interface{}) {
	if is_printable(tag, INFO) {
		_print_info(tag, "I", msg)
	}
}

func (l *defaultLogger) If(tag string, msg string, args ...interface{}) {
	if is_printable(tag, INFO) {
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
		tag: tag,
	}
}

type simpleLogger struct {
	logger Logger
	tag    string
}

func (l *simpleLogger) D(msg ...interface{}) {
	l.logger.D(l.tag, msg)
}

func (l *simpleLogger) Df(msg string, args ...interface{}) {
	l.logger.Df(l.tag, msg, args...)
}

func (l *simpleLogger) V(msg ...interface{}) {
	l.logger.V(l.tag, msg...)
}

func (l *simpleLogger) Vf(msg string, args ...interface{}) {
	l.logger.Vf(l.tag, msg, args...)
}

func (l *simpleLogger) E(msg ...interface{}) {
	l.logger.E(l.tag, msg...)
}

func (l *simpleLogger) Ef(msg string, args ...interface{}) {
	l.logger.Ef(l.tag, msg, args...)
}

func (l *simpleLogger) W(msg ...interface{}) {
	l.logger.W(l.tag, msg...)
}

func (l *simpleLogger) Wf(msg string, args ...interface{}) {
	l.logger.Wf(l.tag, msg, args...)
}

func (l *simpleLogger) I(msg ...interface{}) {
	l.logger.I(l.tag, msg...)
}

func (l *simpleLogger) If(msg string, args ...interface{}) {
	l.logger.If(l.tag, msg, args...)
}

func (l *simpleLogger) Fatal(err error) {
	l.logger.Fatal(l.tag, err)
}
