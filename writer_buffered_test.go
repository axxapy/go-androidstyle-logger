package l

import (
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type slowWriter struct{}

func (w *slowWriter) Write(p []byte) (n int, err error) {
	time.Sleep(time.Second)
	return 0, nil
}

func testBufferedWriterWrite(where io.Writer, what []byte) bool {
	done := make(chan struct{})
	go func() {
		where.Write(what)
		close(done)
	}()
	select {
	case <-done:
		time.Sleep(time.Millisecond)
		return true
	case <-time.After(10 * time.Millisecond):
		return false
	}
}

func TestBufferedWriter_Write_Fast(t *testing.T) {
	var out InMemoryWriter

	w := BufferedWriter(&out, 3)

	for i := 0; i < 3; i++ {
		what := []byte(strconv.Itoa(i))
		assert.True(t, testBufferedWriterWrite(w, what))
		assert.Equal(t, what, out.Last())
	}
}

func TestBufferedWriter_Write_Slow(t *testing.T) {
	var out slowWriter

	w := BufferedWriter(&out, 3)

	write := func(what []byte) bool {
		done := make(chan struct{})
		go func() {
			w.Write(what)
			close(done)
		}()
		select {
		case <-done:
			time.Sleep(time.Millisecond)
			return true
		case <-time.After(10 * time.Millisecond):
			return false
		}
	}

	for i := 0; i <= 3; i++ {
		what := []byte(strconv.Itoa(i))
		assert.True(t, write(what))
	}

	assert.False(t, write([]byte{123}))
	assert.False(t, write([]byte{123}))
}
