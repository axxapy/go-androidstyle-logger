package l

import (
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	w = new(InMemoryWriter)
	l = newBaseLogger().SetWriter(w).(*baseLogger)
)

func TestLevelFuncs(t *testing.T) {
	w := new(InMemoryWriter)
	l := newBaseLogger().SetWriter(w).SetFlags(FLAG_NO_FILENAME)

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

		expected := string(DefaultFormatter("MY_TAG", level, "", 0, "some", "message", 123))
		lst := string(w.Last())
		assert.Equal(t, expected, lst)

		w.Reset()
		l.SetLogLevel(ALL ^ level)
		f("MY_TAG", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

type callInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
}

func retrieveCallInfo() callInfo {
	pc, file, line, _ := runtime.Caller(0)
	_, fileName := path.Split(file)
	tmp := runtime.FuncForPC(pc).Name()
	parts := strings.Split(tmp, ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return callInfo{
		packageName: packageName,
		fileName:    fileName,
		funcName:    funcName,
		line:        line,
	}
}

func TestLevelFuncs_f(t *testing.T) {
	w := new(InMemoryWriter)
	l := newBaseLogger().SetWriter(w).SetFlags(FLAG_NO_FILENAME)

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

		expected := string(DefaultFormatter("MY_TAG", level, "", 0, "some - message - 123"))
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
	assert.Equal(t, ALL, l.logLevel)

	l.SetLogLevel(WARNING)
	assert.Equal(t, WARNING, l.logLevel)

	l.SetLogLevel(0, "TAG")
	assert.Equal(t, WARNING, l.logLevel)
	assert.Equal(t, LogLevel(0), l.logLevelPerTag["TAG"])
}

func TestDefaultLogger_ResetLogLevel(t *testing.T) {
	l.SetLogLevel(ALL)
	assert.Equal(t, ALL, l.logLevel)

	l.SetLogLevel(INFO, "TAG")
	assert.Equal(t, ALL, l.logLevel)
	assert.Equal(t, INFO, l.logLevelPerTag["TAG"])

	l.ResetLogLevel()
	assert.Equal(t, LOG_LEVEL_DEFAULT, l.logLevel)
	assert.Equal(t, INFO, l.logLevelPerTag["TAG"])

	l.ResetLogLevel("TAG")
	assert.Equal(t, LogLevel(0), l.logLevelPerTag["TAG"])
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
	assert.Equal(t, "XXX", ll.(*taggedLogger).tag)
}

func TestDefaultLogger_Check(t *testing.T) {
	l := newBaseLogger().SetLogLevel(ALL)
	assert.NotNil(t, l.Check(INFO))

	l.SetLogLevel(ALL ^ INFO)
	assert.Nil(t, l.Check(INFO))
}

func TestDefaultLogger_CheckWithTag(t *testing.T) {
	l := newBaseLogger().SetLogLevel(ALL)
	assert.NotNil(t, l.CheckWithTag(INFO, "MY_TAG"))

	l.SetLogLevel(ALL^INFO, "MY_TAG")
	assert.Nil(t, l.CheckWithTag(INFO, "MY_TAG"))
	assert.NotNil(t, l.CheckWithTag(WARNING, "MY_TAG"))
	assert.NotNil(t, l.CheckWithTag(INFO, "OTHER_TAG"))
}
