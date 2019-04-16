package l

import (
	"fmt"
	"log"
)

const (
	DEBUG   = 1
	ERROR   = 2
	INFO    = 4
	VERBOSE = 8
	WARNING = 16
	WTF     = 32
	ALL     = DEBUG ^ ERROR ^ INFO ^ VERBOSE ^ WARNING ^ WTF

	LOG_LEVEL_DEFAULT = WARNING ^ ERROR ^ INFO
)

var (
	log_level     = LOG_LEVEL_DEFAULT
	log_level_tag = map[string]int{}
	//logger    = log.New()
	logger = &defaultLogger{}
)

func SetLogLevel(level int, tags ...string) {
	if len(tags) < 1 {
		log_level = level
	} else {
		for _, tag := range tags {
			log_level_tag[tag] = level
		}
	}
}

func ResetLogLevel(tags ...string) {
	if len(tags) < 1 {
		log_level = LOG_LEVEL_DEFAULT
	} else {
		for _, tag := range tags {
			delete(log_level_tag, tag)
		}
	}
}

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

	WithTag(tag string) SimpleLogger
}

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

func _print_error(tag string, level string, msg ...interface{}) {
	log.Print("["+level+"] ["+tag+"] " + fmt.Sprintln(msg...))
}

func _print_info(tag string, level string, msg ...interface{}) {
	log.Output(2, "["+level+"] ["+tag+"] " + fmt.Sprintln(msg...))
}

func _format(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...);
}

func is_printable(tag string, level int) bool {
	print_level := log_level
	if tag_level, exists := log_level_tag[tag]; exists {
		print_level = tag_level
	}
	return print_level & level != 0
}

func Check(level int) Logger {
	if is_printable("", level) {
		return logger
	}
	return nil
}

func CheckTag(tag string, level int) SimpleLogger {
	if is_printable(tag, level) {
		return logger.WithTag(tag)
	}
	return nil
}