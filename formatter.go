package l

import (
	"encoding/json"
	"fmt"
	"time"
)

type Formatter func(l Logger, tag string, level LogLevel, msg ...interface{}) []byte

type jsonLogEntry struct {
	Timestamp int64  `json:"timestamp"`
	Level     int    `json:"level"`
	LevelName string `json:"levelName"`
	Tag       string `json:"tag"`
	Msg       string `json:"msg"`
}

var Now = func() time.Time {
	return time.Now()
}

func DefaultFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	return []byte("[" + Now().Format("2006-01-02 15:04:05.000") + "] [" + GetLevelName(level) + "] [" + tag + "] " + fmt.Sprintln(msg...))
}

func JsonFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	message := fmt.Sprintln(msg...)
	if len(msg) > 0 {
		message = message[:len(message)-1]
	}
	b, err := json.Marshal(jsonLogEntry{
		Timestamp: Now().UnixNano(),
		Tag:       tag,
		Level:     int(level),
		LevelName: GetLevelName(level),
		Msg:       message,
	})
	if err != nil {
		panic(err)
	}
	return append(b, '\n')
}
