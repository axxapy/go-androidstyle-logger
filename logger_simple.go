package l

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

type simpleLogger struct {
	logger Logger
	tag    string
}

func (l *simpleLogger) D(msg ...interface{}) {
	l.logger.D(l.tag, msg...)
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
