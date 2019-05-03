package l

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "go-androidstyle-logger/_mocks"
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

func TestStatic_testLogFuncs(t *testing.T) {
	funcs := map[LogLevel]func(tag string, msg ...interface{}){
		DEBUG:   D,
		ERROR:   E,
		INFO:    I,
		VERBOSE: V,
		WARNING: W,
	}

	w := &mocks.Writer{}

	for level, f := range funcs {
		logger = New().SetLogLevel(ALL).SetWriter(w)
		w.Reset()

		expected := DefaultFormatter(logger, "TAG", level, "some", "message", 123)
		f("TAG", "some", "message", 123)

		assert.Equal(t, string(expected), string(w.Last()))

		w.Reset()
		logger.SetLogLevel(ALL ^ level)
		f("TAG", "some", "message", 123)
		assert.Nil(t, w.Last())

		logger.SetLogLevel(ALL).SetLogLevel(0, "TAG")
		f("TAG", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestStatic_testLogFuncs_f(t *testing.T) {
	funcs := map[LogLevel]func(tag string, msg string, args ...interface{}){
		DEBUG:   Df,
		ERROR:   Ef,
		INFO:    If,
		VERBOSE: Vf,
		WARNING: Wf,
	}

	w := &mocks.Writer{}

	for level, f := range funcs {
		logger = New().SetLogLevel(ALL).SetWriter(w)
		w.Reset()

		expected := DefaultFormatter(logger, "TAG", level, "some - message - 123")
		f("TAG", "%s - %s - %d", "some", "message", 123)

		assert.Equal(t, string(expected), string(w.Last()))

		w.Reset()
		logger.SetLogLevel(ALL ^ level)
		f("TAG", "some", "message", 123)
		assert.Nil(t, w.Last())

		logger.SetLogLevel(ALL).SetLogLevel(0, "TAG")
		f("TAG", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestCheck(t *testing.T) {
	logger = New().SetLogLevel(ALL)
	assert.NotNil(t, Check(INFO))

	SetLogLevel(ALL^INFO)
	assert.Nil(t, Check(INFO))
}

func TestCheckWithTag(t *testing.T) {
	logger = New().SetLogLevel(ALL)
	assert.NotNil(t, CheckWithTag(INFO, "TAG_NAME"))

	SetLogLevel(ALL^INFO, "TAG_NAME")
	assert.Nil(t, CheckWithTag(INFO, "TAG_NAME"))
	assert.NotNil(t, CheckWithTag(INFO, "ANOTHER_TAG"))
}
