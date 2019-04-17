package l

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter(t *testing.T) {
	l := &defaultLogger{}
	tag := "TAG"

	for level, levelName := range levelNames {
		line := DefaultFormatter(l, tag, level, "some", "message", 123)
		expect := fmt.Sprintf("[%s] [%s] some message 123\n", levelName, tag)
		assert.Equal(t, expect, string(line))
	}
}

func TestJsonFormatter(t *testing.T) {
	l := &defaultLogger{}
	tag := "TAG"

	for level, levelName := range levelNames {
		line := JsonFormatter(l, tag, level, "some", "message", 123)
		expect := fmt.Sprintf(`{"level":"%d","levelName":"%s","msg":"some message 123","tag":"%s"}`+"\n", level, levelName, tag)
		assert.Equal(t, expect, string(line))
	}
}
