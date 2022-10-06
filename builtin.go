package l

import (
	"errors"
	"log"
)

func ReplaceGoDefaultLogger(logger Logger, defaultLevel LogLevel) error {
	logger = logger.WithTag("builtin")
	var logFunc builtinWriter
	switch defaultLevel {
	case DEBUG:
		logFunc = logger.D
		break
	case ERROR:
		logFunc = logger.E
		break
	case INFO:
		logFunc = logger.I
		break
	case VERBOSE:
		logFunc = logger.V
		break
	case WARNING:
		logFunc = logger.W
		break
	default:
		return errors.New("only following log levels are allowed: DEBUG, ERROR, INFO, VERBOSE, WARNING")
	}

	log.Default().SetOutput(builtinWriter(logFunc))
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	if l, ok := logger.(*taggedLogger); ok {
		l.baseLogger.callerDeep += 3
	}

	return nil
}

type builtinWriter func(msg ...interface{})

func (w builtinWriter) Write(p []byte) (n int, err error) {
	if p[len(p)-1] == '\n' {
		w(string(p[:len(p)-1]))
	} else {
		w(string(p))
	}
	return len(p), nil
}
