package l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryWriter_Write(t *testing.T) {
	test := []byte("AXP")
	var w InMemoryWriter
	n, err := w.Write(test)
	assert.NoError(t, err)
	assert.Equal(t, len(test), n)
	assert.Equal(t, test, w.last)
}

func TestInMemoryWriter_Last(t *testing.T) {
	test := []byte("AXP")
	var w InMemoryWriter
	n, err := w.Write(test)
	assert.NoError(t, err)
	assert.Equal(t, len(test), n)
	assert.Equal(t, test, w.Last())
}

func TestInMemoryWriter_Reset(t *testing.T) {
	test := []byte("AXP")
	var w InMemoryWriter
	n, err := w.Write(test)
	assert.NoError(t, err)
	assert.Equal(t, len(test), n)
	assert.Equal(t, test, w.Last())
	assert.Equal(t, test, w.Last())
	w.Reset()
	assert.Nil(t, w.last)
}
