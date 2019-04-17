package l

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	lSimple = l.WithTag("MY_TAG")
)

func testSimpleLevel(t *testing.T, level LogLevel, f func(msg ...interface{})) {
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

func testSimpleLevelf(t *testing.T, level LogLevel, f func(msg string, args ...interface{})) {
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

func TestSimpleLogger_D(t *testing.T) {
	testSimpleLevel(t, DEBUG, lSimple.D)
}

func TestSimpleLogger_Df(t *testing.T) {
	testSimpleLevelf(t, DEBUG, lSimple.Df)
}

func TestSimpleLogger_E(t *testing.T) {
	testSimpleLevel(t, ERROR, lSimple.E)
}

func TestSimpleLogger_Ef(t *testing.T) {
	testSimpleLevelf(t, ERROR, lSimple.Ef)
}

func TestSimpleLogger_I(t *testing.T) {
	testSimpleLevel(t, INFO, lSimple.I)
}

func TestSimpleLogger_If(t *testing.T) {
	testSimpleLevelf(t, INFO, lSimple.If)
}

func TestSimpleLogger_V(t *testing.T) {
	testSimpleLevel(t, VERBOSE, lSimple.V)
}

func TestSimpleLogger_Vf(t *testing.T) {
	testSimpleLevelf(t, VERBOSE, lSimple.Vf)
}

func TestSimpleLogger_W(t *testing.T) {
	testSimpleLevel(t, WARNING, lSimple.W)
}

func TestSimpleLogger_Wf(t *testing.T) {
	testSimpleLevelf(t, WARNING, lSimple.Wf)
}
