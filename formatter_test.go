package l

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter(t *testing.T) {
	origNow := now
	defer func() { now = origNow }()

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02 15:04:05.999", "2019-05-27 02:23:14.34")
		return t
	}

	tag := "TAG"

	for level, levelName := range levelNames {
		result := DefaultFormatter(tag, level, "some", "message", 123)
		expected := fmt.Sprintf("[%s] [%s] [%s] some message 123\n", now().Format("2006-01-02 15:04:05.000"), levelName, tag)
		assert.Equal(t, expected, string(result))
	}
}

func TestJsonFormatter(t *testing.T) {
	origNow := now
	defer func() { now = origNow }()

	frozenTime := time.Now()
	now = func() time.Time {
		return frozenTime
	}

	tag := "TAG"
	ts := now().UnixNano()

	for level, levelName := range levelNames {
		line := JsonFormatter(tag, level, "some", "message", 123)
		expect := fmt.Sprintf(`{"timestamp":%d,"level":%d,"levelName":"%s","tag":"%s","msg":"some message 123"}`+"\n", ts, level, levelName, tag)
		assert.Equal(t, expect, string(line))
	}
}
