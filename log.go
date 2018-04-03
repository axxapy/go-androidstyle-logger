package l

import (
	"log"
	"fmt"
	"os"
)

const (
	DEBUG   = 1
	ERROR   = 2
	INFO    = 4
	VERBOSE = 8
	WARNING = 16
	WTF     = 32
	ALL     = DEBUG ^ ERROR ^ INFO ^ VERBOSE ^ WARNING ^ WTF
)

var (
	log_level     = WARNING ^ ERROR ^ INFO
	log_level_tag = map[string]int{}
	//logger    = log.New()
)

func SetLogLevel(level int) {
	log_level = level
}

func SetTagLogLevel(tag string, level int) {
	log_level_tag[tag] = level
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

func D(tag string, msg ...interface{}) {
	if is_printable(tag, DEBUG) {
		_print_info(tag, "D", msg)
	}
}

func Df(tag string, msg string, args ...interface{}) {
	if is_printable(tag, DEBUG) {
		_print_info(tag, "D", _format(msg, args...))
	}
}

func V(tag string, msg ...interface{}) {
	if is_printable(tag, VERBOSE) {
		_print_info(tag, "V", msg)
	}
}

func Vf(tag string, msg string, args ...interface{}) {
	if is_printable(tag, VERBOSE) {
		_print_info(tag, "V", _format(msg, args...))
	}
}

func E(tag string, msg ...interface{}) {
	if is_printable(tag, ERROR) {
		_print_error(tag, "E", msg)
	}
}

func Ef(tag string, msg string, args ...interface{}) {
	if is_printable(tag, ERROR) {
		_print_error(tag, "E", _format(msg, args...))
	}
}

func W(tag string, msg ...interface{}) {
	if is_printable(tag, WARNING) {
		_print_info(tag, "W", msg)
	}
}

func Wf(tag string, msg string, args ...interface{}) {
	if is_printable(tag, WARNING) {
		_print_info(tag, "W", _format(msg, args...))
	}
}

func I(tag string, msg ...interface{}) {
	if is_printable(tag, INFO) {
		_print_info(tag, "I", msg)
	}
}

func If(tag string, msg string, args ...interface{}) {
	if is_printable(tag, INFO) {
		_print_info(tag, "I", _format(msg, args...))
	}
}

func Fatal(tag string, err error) {
	if err != nil {
		E(tag, err)
		os.Exit(1)
	}
}

func is_printable(tag string, level int) bool {
	print_level := log_level
	if tag_level, exists := log_level_tag[tag]; exists {
		print_level = tag_level
	}
	return print_level & level != 0
}
