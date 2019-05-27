package l

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter(t *testing.T) {
	origNow := Now
	defer func() {Now = origNow}()

	now, _ := time.Parse("2006-01-02 15:04:05.999", "2019-05-27 02:23:14.34")
	Now = func() time.Time {
		return now
	}

	l := &defaultLogger{}
	tag := "TAG"

	for level, levelName := range levelNames {
		result := DefaultFormatter(l, tag, level, "some", "message", 123)
		expected := fmt.Sprintf("[%s] [%s] [%s] some message 123\n", Now().Format("2006-01-02 15:04:05.000"), levelName, tag)
		assert.Equal(t, expected, string(result))
	}
}

func TestJsonFormatter(t *testing.T) {
	origNow := Now
	defer func() {Now = origNow}()

	now := time.Now()
	Now = func() time.Time {
		return now
	}

	l := &defaultLogger{}
	tag := "TAG"
	ts := Now().UnixNano()

	for level, levelName := range levelNames {
		line := JsonFormatter(l, tag, level, "some", "message", 123)
		expect := fmt.Sprintf(`{"timestamp":%d,"level":%d,"levelName":"%s","tag":"%s","msg":"some message 123"}`+"\n", ts, level, levelName, tag)
		assert.Equal(t, expect, string(line))
	}
}
