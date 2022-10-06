package l

import (
	"encoding/json"
	"fmt"
	"time"
)

type Formatter func(tag string, level LogLevel, filename string, line int, msg ...interface{}) []byte

type jsonLogEntry struct {
	Timestamp int64  `json:"timestamp"`
	Level     int    `json:"level"`
	LevelName string `json:"levelName"`
	FileName  string `json:"fileName"`
	Line      int    `json:"line"`
	Tag       string `json:"tag"`
	Msg       string `json:"msg"`
}

var now = func() time.Time {
	return time.Now()
}

func DefaultFormatter(tag string, level LogLevel, filename string, line int, msg ...interface{}) []byte {
	tpl := "[%s] [%s] "
	args := []interface{}{now().Format("2006-01-02 15:04:05.000"), GetLevelName(level)}
	if tag != "" {
		tpl += "[%s] "
		args = append(args, tag)
	}
	if filename != "" {
		tpl += "[%s:%d] "
		args = append(args, filename, line)
	}
	return []byte(fmt.Sprintf(tpl, args...) + fmt.Sprintln(msg...))
}

func JsonFormatter(tag string, level LogLevel, filename string, line int, msg ...interface{}) []byte {
	message := fmt.Sprintln(msg...)
	if len(msg) > 0 {
		message = message[:len(message)-1]
	}
	b, err := json.Marshal(jsonLogEntry{
		Timestamp: now().UnixNano(),
		Tag:       tag,
		Level:     int(level),
		LevelName: GetLevelName(level),
		FileName:  filename,
		Line:      line,
		Msg:       message,
	})
	if err != nil {
		panic(err)
	}
	return append(b, '\n')
}
