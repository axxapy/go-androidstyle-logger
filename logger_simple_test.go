package l

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleLevelFuncs(t *testing.T) {
	w := new(InMemoryWriter)
	l := New().SetLogLevel(ALL).SetWriter(w)
	lSimple := l.WithTag("MY_TAG")

	funcs := map[LogLevel]func(msg ...interface{}){
		DEBUG:   lSimple.D,
		ERROR:   lSimple.E,
		INFO:    lSimple.I,
		VERBOSE: lSimple.V,
		WARNING: lSimple.W,
	}

	for level, f := range funcs {
		w.Reset()
		l.SetLogLevel(ALL)
		f("some", "message", 123)

		expected := string(DefaultFormatter(l, "MY_TAG", level, "some", "message", 123))
		assert.Equal(t, expected, string(w.Last()))

		w.Reset()
		l.SetLogLevel(ALL ^ level)
		f("some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestSimpleLevelFuncs_f(t *testing.T) {
	w := new(InMemoryWriter)
	l := New().SetLogLevel(ALL).SetWriter(w)
	lSimple := l.WithTag("MY_TAG")

	funcs := map[LogLevel]func(msg string, args ...interface{}){
		DEBUG:   lSimple.Df,
		ERROR:   lSimple.Ef,
		INFO:    lSimple.If,
		VERBOSE: lSimple.Vf,
		WARNING: lSimple.Wf,
	}

	for level, f := range funcs {
		w.Reset()
		l.SetLogLevel(ALL)
		f("%s - %s - %d", "some", "message", 123)

		expected := string(DefaultFormatter(l, "MY_TAG", level, "some - message - 123"))
		assert.Equal(t, expected, string(w.Last()))

		w.Reset()
		l.SetLogLevel(ALL ^ level)
		f("%s - %s - %d", "some", "message", 123)
		assert.Nil(t, w.Last())
	}
}

func TestSimpleLogger_Check(t *testing.T) {
	l := New().SetLogLevel(ALL)

	lt := l.WithTag("MY_TAG")

	assert.NotNil(t, lt.Check(INFO))

	l.SetLogLevel(ALL ^ INFO)
	assert.Nil(t, lt.Check(INFO))

	l.SetLogLevel(ALL, "MY_TAG")
	assert.NotNil(t, lt.Check(INFO))
}
