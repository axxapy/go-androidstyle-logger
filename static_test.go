package l

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestGetLevelName(t *testing.T) {
	assert.Equal(t, "I", GetLevelName(INFO))
	assert.Equal(t, "63", GetLevelName(ALL))
}

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(ALL)
	assert.Equal(t, ALL, logger.(*defaultLogger).log_level)

	SetLogLevel(INFO)
	assert.Equal(t, INFO, logger.(*defaultLogger).log_level)

	SetLogLevel(WARNING, "TAG1", "TAG2")
	assert.Equal(t, INFO, logger.(*defaultLogger).log_level)
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG1"])
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG2"])
}

func TestResetLogLevel(t *testing.T) {
	SetLogLevel(ALL)
	SetLogLevel(WARNING, "TAG1", "TAG2")
	assert.Equal(t, ALL, logger.(*defaultLogger).log_level)
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG1"])
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG2"])

	ResetLogLevel()
	assert.Equal(t, LOG_LEVEL_DEFAULT, logger.(*defaultLogger).log_level)
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG1"])
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG2"])

	ResetLogLevel("TAG1")
	assert.Equal(t, LOG_LEVEL_DEFAULT, logger.(*defaultLogger).log_level)
	assert.Equal(t, LogLevel(0), logger.(*defaultLogger).log_level_tag["TAG1"])
	assert.Equal(t, WARNING, logger.(*defaultLogger).log_level_tag["TAG2"])
}

func TestSetFormatter(t *testing.T) {
	jsFuncPointer := reflect.ValueOf(JsonFormatter).Pointer()
	assert.NotEqual(t, reflect.ValueOf(logger.(*defaultLogger).formatter).Pointer(), jsFuncPointer)

	SetFormatter(JsonFormatter)
	assert.Equal(t, reflect.ValueOf(logger.(*defaultLogger).formatter).Pointer(), jsFuncPointer)
}

func TestSetWriter(t *testing.T) {
	assert.NotEqual(t, os.Stdout, logger.(*defaultLogger).writer)

	SetWriter(os.Stdout)
	assert.Equal(t, os.Stdout, logger.(*defaultLogger).writer)
}

func TestWithTag(t *testing.T) {
	tagged := WithTag("TAG")
	assert.Equal(t, "TAG", tagged.(*simpleLogger).tag)
}

func testStatic_testLogFunc(t *testing.T, level LogLevel, f func (tag string, msg ...interface{})) {
	logger.SetLogLevel(ALL)
	logger.SetWriter(w)
	w.Reset()

	expected := DefaultFormatter(logger, "TAG", level, "some", "message", 123)
	f("TAG", "some", "message", 123)

	assert.Equal(t, string(expected), string(w.Last()))

	w.Reset()
	logger.SetLogLevel(ALL^level)
	f("TAG", "some", "message", 123)
	assert.Nil(t, w.Last())

	logger.SetLogLevel(ALL)
	logger.SetLogLevel(0, "TAG")
	f("TAG", "some", "message", 123)
	assert.Nil(t, w.Last())
}

func testStatic_testLogFuncf(t *testing.T, level LogLevel, f func (tag string, msg string, args ...interface{})) {
	logger.SetLogLevel(ALL)
	logger.SetWriter(w)
	w.Reset()

	expected := DefaultFormatter(logger, "TAG", level, "some - message - 123",)
	f("TAG", "%s - %s - %d", "some", "message", 123)

	assert.Equal(t, string(expected), string(w.Last()))

	w.Reset()
	logger.SetLogLevel(ALL^level)
	f("TAG", "some", "message", 123)
	assert.Nil(t, w.Last())

	logger.SetLogLevel(ALL)
	logger.SetLogLevel(0, "TAG")
	f("TAG", "some", "message", 123)
	assert.Nil(t, w.Last())
}

func TestD(t *testing.T) {
	testStatic_testLogFunc(t, DEBUG, D)
}

func TestDf(t *testing.T) {
	testStatic_testLogFuncf(t, DEBUG, Df)
}

func TestE(t *testing.T) {
	testStatic_testLogFunc(t, ERROR, E)
}

func TestEf(t *testing.T) {
	testStatic_testLogFuncf(t, ERROR, Ef)
}

func TestI(t *testing.T) {
	testStatic_testLogFunc(t, INFO, I)
}

func TestIf(t *testing.T) {
	testStatic_testLogFuncf(t, INFO, If)
}

func TestV(t *testing.T) {
	testStatic_testLogFunc(t, VERBOSE, V)
}

func TestVf(t *testing.T) {
	testStatic_testLogFuncf(t, VERBOSE, Vf)
}

func TestW(t *testing.T) {
	testStatic_testLogFunc(t, WARNING, W)
}

func TestWf(t *testing.T) {
	testStatic_testLogFuncf(t, WARNING, Wf)
}
