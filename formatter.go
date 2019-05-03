package l

import (
	"encoding/json"
	"fmt"
)

type Formatter func(l Logger, tag string, level LogLevel, msg ...interface{}) []byte

type jsonLogEntry struct {
	Level     int    `json:"level"`
	LevelName string `json:"levelName"`
	Tag       string `json:"tag"`
	Msg       string `json:"msg"`
}

func DefaultFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	return []byte("[" + GetLevelName(level) + "] [" + tag + "] " + fmt.Sprintln(msg...))
}

func JsonFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	message := fmt.Sprintln(msg...)
	if len(msg) > 0 {
		message = message[:len(message)-1]
	}
	b, err := json.Marshal(jsonLogEntry{
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
