package l

import (
	"github.com/stretchr/testify/assert"
	"go-androidstyle-logger/_mocks"
	"reflect"
	"testing"
)

var (
	w = &mocks.Writer{}
	l = New().SetWriter(w).(*defaultLogger)
)

func testLevel(t *testing.T, level LogLevel, f func (tag string, msg ...interface{})) {
	w.Reset()
	l.SetLogLevel(ALL)
	f("MY_TAG", "some", "message", 123)

	expected := string(DefaultFormatter(l, "MY_TAG", level, "some", "message", 123))
	assert.Equal(t, expected, string(w.Last()))

	w.Reset()
	l.SetLogLevel(ALL^level)
	f("MY_TAG", "some", "message", 123)
	assert.Nil(t, w.Last())
}

func testLevelf(t *testing.T, level LogLevel, f func (tag string, msg string, args ...interface{})) {
	w.Reset()
	l.SetLogLevel(ALL)
	f("MY_TAG", "%s - %s - %d", "some", "message", 123)

	expected := string(DefaultFormatter(l, "MY_TAG", level, "some - message - 123"))
	assert.Equal(t, expected, string(w.Last()))

	w.Reset()
	l.SetLogLevel(ALL^level)
	f("MY_TAG", "%s - %s - %d", "some", "message", 123)
	assert.Nil(t, w.Last())
}

func TestDefaultLogger_D(t *testing.T) {
	testLevel(t, DEBUG, l.D)
}

func TestDefaultLogger_Df(t *testing.T) {
	testLevelf(t, DEBUG, l.Df)
}

func TestDefaultLogger_E(t *testing.T) {
	testLevel(t, ERROR, l.E)
}

func TestDefaultLogger_Ef(t *testing.T) {
	testLevelf(t, ERROR, l.Ef)
}

func TestDefaultLogger_I(t *testing.T) {
	testLevel(t, INFO, l.I)
}

func TestDefaultLogger_If(t *testing.T) {
	testLevelf(t, INFO, l.If)
}

func TestDefaultLogger_V(t *testing.T) {
	testLevel(t, VERBOSE, l.V)
}

func TestDefaultLogger_Vf(t *testing.T) {
	testLevelf(t, VERBOSE, l.Vf)
}

func TestDefaultLogger_W(t *testing.T) {
	testLevel(t, WARNING, l.W)
}

func TestDefaultLogger_Wf(t *testing.T) {
	testLevelf(t, WARNING, l.Wf)
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

	l.SetLogLevel(ALL^INFO)
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