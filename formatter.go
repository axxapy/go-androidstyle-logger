package l

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Formatter func(l Logger, tag string, level LogLevel, msg ...interface{}) []byte

func DefaultFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	return []byte("[" + GetLevelName(level) + "] [" + tag + "] " + fmt.Sprintln(msg...) + "\n")
}

func JsonFormatter(l Logger, tag string, level LogLevel, msg ...interface{}) []byte {
	b, err := json.Marshal(map[string]string{
		"tag":       tag,
		"level":     strconv.Itoa(int(level)),
		"levelName": GetLevelName(level),
		"msg":       fmt.Sprintln(msg...),
	})
	if err != nil {
		panic(err)
	}
	return append(b, '\n')
}
