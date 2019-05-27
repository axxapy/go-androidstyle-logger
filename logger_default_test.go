package l

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	w = new(InMemoryWriter)
	l = New().SetWriter(w).(*defaultLogger)
)

func TestLevelFuncs(t *testing.T) {
	w := new(InMemoryWriter)
	l := New().SetWriter(w)

	funcs := map[LogLevel]func(tag string, msg ...interface{}){
		DEBUG:   l.D,
		ERROR:   l.E,
		INFO:    l.I,
		VERBOSE: l.V,
		WARNING: l.W,
	}

	for level, f := range funcs {
		l.SetLogLevel(ALL).ResetLogLevel("MY_TAG")
		w.Reset()

		f("MY_TAG", "some", "message", 123)

		expected := string(DefaultFormatter(l, "MY_TAG", level, "some", "message", 123))
		lst := string(w.Last())
		assert.Equal(t, expected, lst)

		w.Reset()
		l.SetLogLevel(ALL ^ level)
		f("MY_TAG", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestLevelFuncs_f(t *testing.T) {
	w := new(InMemoryWriter)
	l := New().SetWriter(w)

	funcs := map[LogLevel]func(tag string, msg string, args ...interface{}){
		DEBUG:   l.Df,
		ERROR:   l.Ef,
		INFO:    l.If,
		VERBOSE: l.Vf,
		WARNING: l.Wf,
	}

	for level, f := range funcs {
		l.SetLogLevel(ALL).ResetLogLevel("MY_TAG")
		w.Reset()

		f("MY_TAG", "%s - %s - %d", "some", "message", 123)

		expected := string(DefaultFormatter(l, "MY_TAG", level, "some - message - 123"))
		assert.Equal(t, expected, string(w.Last()))

		w.Reset()
		l.SetLogLevel(ALL ^ level)
		f("MY_TAG", "%s - %s - %d", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestDefaultLogger_SetFormatter(t *testing.T) {
	jsFuncPointer := reflect.ValueOf(JsonFormatter).Pointer()
	assert.NotEqual(t, reflect.ValueOf(l.formatter).Pointer(), jsFuncPointer)

	l.SetFormatter(JsonFormatter)
	assert.Equal(t, reflect.ValueOf(l.formatter).Pointer(), jsFuncPointer)
}

func TestDefaultLogger_SetLogLevel(t *testing.T) {
	l.SetLogLevel(ALL)
	assert.Equal(t, ALL, l.log_level)

	l.SetLogLevel(WARNING)
	assert.Equal(t, WARNING, l.log_level)

	l.SetLogLevel(0, "TAG")
	assert.Equal(t, WARNING, l.log_level)
	assert.Equal(t, LogLevel(0), l.log_level_tag["TAG"])
}

func TestDefaultLogger_ResetLogLevel(t *testing.T) {
	l.SetLogLevel(ALL)
	assert.Equal(t, ALL, l.log_level)

	l.SetLogLevel(INFO, "TAG")
	assert.Equal(t, ALL, l.log_level)
	assert.Equal(t, INFO, l.log_level_tag["TAG"])

	l.ResetLogLevel()
	assert.Equal(t, LOG_LEVEL_DEFAULT, l.log_level)
	assert.Equal(t, INFO, l.log_level_tag["TAG"])

	l.ResetLogLevel("TAG")
	assert.Equal(t, LogLevel(0), l.log_level_tag["TAG"])
}

func TestDefaultLogger_IsLoggable(t *testing.T) {
	l.SetLogLevel(ALL)
	assert.True(t, l.IsLoggable(INFO, ""))
	assert.True(t, l.IsLoggable(INFO, "TAG"))

	l.SetLogLevel(ALL ^ INFO)
	assert.False(t, l.IsLoggable(INFO, ""))
	assert.False(t, l.IsLoggable(INFO, "TAG"))

	l.SetLogLevel(ALL, "TAG", "TAG1")
	assert.False(t, l.IsLoggable(INFO, ""))
	assert.True(t, l.IsLoggable(INFO, "TAG"))
	assert.True(t, l.IsLoggable(INFO, "TAG1"))
	assert.False(t, l.IsLoggable(INFO, "TAG2"))
}

func TestDefaultLogger_WithTag(t *testing.T) {
	ll := l.WithTag("XXX")
	assert.Equal(t, "XXX", ll.(*simpleLogger).tag)
}

func TestDefaultLogger_Check(t *testing.T) {
	l := New().SetLogLevel(ALL)
	assert.NotNil(t, l.Check(INFO))

	l.SetLogLevel(ALL ^ INFO)
	assert.Nil(t, l.Check(INFO))
}

func TestDefaultLogger_CheckWithTag(t *testing.T) {
	l := New().SetLogLevel(ALL)
	assert.NotNil(t, l.CheckWithTag(INFO, "MY_TAG"))

	l.SetLogLevel(ALL ^ INFO, "MY_TAG")
	assert.Nil(t, l.CheckWithTag(INFO, "MY_TAG"))
	assert.NotNil(t, l.CheckWithTag(WARNING, "MY_TAG"))
	assert.NotNil(t, l.CheckWithTag(INFO, "OTHER_TAG"))
}
