package l

func SetLogLevel(level LogLevel, tags ...string) {
	logger.SetLogLevel(level, tags...)
}

func ResetLogLevel(tags ...string) {
	logger.ResetLogLevel(tags...)
}

func WithTag(tag string) SimpleLogger {
	return logger.WithTag(tag)
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