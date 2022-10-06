package l

import "io"

type Logger interface {
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

	GetTag() string
	SetTag(tag string)

	Fatal(err error)

	SetLogLevel(level LogLevel, tags ...string) Logger
	ResetLogLevel(tags ...string) Logger

	SetFormatter(f Formatter) Logger
	SetWriter(w io.Writer) Logger

	IsLoggable(level LogLevel, tag string) bool
	Check(level LogLevel) Logger

	WithTag(tag string) Logger
}

type taggedLogger struct {
	baseLogger baseLoggerInterface
	tag        string
}

func (l *taggedLogger) D(msg ...interface{}) {
	l.baseLogger.D(l.tag, msg...)
}

func (l *taggedLogger) Df(msg string, args ...interface{}) {
	l.baseLogger.Df(l.tag, msg, args...)
}

func (l *taggedLogger) V(msg ...interface{}) {
	l.baseLogger.V(l.tag, msg...)
}

func (l *taggedLogger) Vf(msg string, args ...interface{}) {
	l.baseLogger.Vf(l.tag, msg, args...)
}

func (l *taggedLogger) E(msg ...interface{}) {
	l.baseLogger.E(l.tag, msg...)
}

func (l *taggedLogger) Ef(msg string, args ...interface{}) {
	l.baseLogger.Ef(l.tag, msg, args...)
}

func (l *taggedLogger) W(msg ...interface{}) {
	l.baseLogger.W(l.tag, msg...)
}

func (l *taggedLogger) Wf(msg string, args ...interface{}) {
	l.baseLogger.Wf(l.tag, msg, args...)
}

func (l *taggedLogger) I(msg ...interface{}) {
	l.baseLogger.I(l.tag, msg...)
}

func (l *taggedLogger) If(msg string, args ...interface{}) {
	l.baseLogger.If(l.tag, msg, args...)
}

func (l *taggedLogger) Fatal(err error) {
	l.baseLogger.Fatal(l.tag, err)
}

func (l *taggedLogger) Check(level LogLevel) Logger {
	if l.baseLogger.IsLoggable(level, l.tag) || l.baseLogger.IsLoggable(level, "") {
		return l
	}
	return nil
}

func (l *taggedLogger) SetTag(tag string) {
	l.tag = tag
}

func (l *taggedLogger) GetTag() string {
	return l.tag
}

func (l *taggedLogger) WithTag(tag string) Logger {
	if l.tag != "" {
		tag = l.tag + ">" + tag
	}
	return &taggedLogger{
		baseLogger: l.baseLogger,
		tag:        tag,
	}
}

func (l *taggedLogger) SetLogLevel(level LogLevel, tags ...string) Logger {
	l.baseLogger.SetLogLevel(level, tags...)
	return l
}

func (l *taggedLogger) ResetLogLevel(tags ...string) Logger {
	l.baseLogger.ResetLogLevel(tags...)
	return l
}

func (l *taggedLogger) SetFormatter(f Formatter) Logger {
	l.baseLogger.SetFormatter(f)
	return l
}

func (l *taggedLogger) SetWriter(w io.Writer) Logger {
	l.baseLogger.SetWriter(w)
	return l
}

func (l *taggedLogger) IsLoggable(level LogLevel, tag string) bool {
	return l.baseLogger.IsLoggable(level, tag)
}
