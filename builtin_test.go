package l

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReplaceGoDefaultLogger(t *testing.T) {
	cases := map[string]struct {
		builtinLogLevel LogLevel
		enabledLogLevel LogLevel
		visible         bool
		expectError     bool
	}{
		"unsupported log level": {
			builtinLogLevel: ALL,
			expectError:     true,
		},
		"not visible log level": {
			builtinLogLevel: INFO,
			enabledLogLevel: WARNING,
		},
		"visible log level - DEBUG": {
			builtinLogLevel: DEBUG,
			enabledLogLevel: ALL,
			visible:         true,
		},
		"visible log level - ERROR": {
			builtinLogLevel: ERROR,
			enabledLogLevel: ALL,
			visible:         true,
		},
		"visible log level - INFO": {
			builtinLogLevel: INFO,
			enabledLogLevel: ALL,
			visible:         true,
		},
		"visible log level - VERBOSE": {
			builtinLogLevel: VERBOSE,
			enabledLogLevel: ALL,
			visible:         true,
		},
		"visible log level - WARNING": {
			builtinLogLevel: WARNING,
			enabledLogLevel: ALL,
			visible:         true,
		},
	}

	w := new(InMemoryWriter)
	logger := New().SetWriter(w).SetFlags(FLAG_NO_FILENAME)

	now = func() time.Time {
		t, _ := time.Parse("2006-01-02 15:04:05.999", "2022-10-05 22:53:31.34")
		return t
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			if test.expectError {
				assert.Error(t, ReplaceGoDefaultLogger(logger, test.builtinLogLevel))
				return
			}

			logger.SetLogLevel(test.enabledLogLevel)
			w.Reset()

			assert.NoError(t, ReplaceGoDefaultLogger(logger, test.builtinLogLevel))
			log.Print("some - message - 123")

			if !test.visible {
				assert.Empty(t, w.last)
				return
			}

			expected := string(DefaultFormatter("builtin", test.builtinLogLevel, "", 0, "some - message - 123"))
			assert.Equal(t, expected, string(w.Last()))
		})
	}
}
