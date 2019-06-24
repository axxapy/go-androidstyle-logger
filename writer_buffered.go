package l

import "io"

type bufferedWriter chan []byte

func BufferedWriter(dest io.Writer, bufferSize int) bufferedWriter {
	w := make(bufferedWriter, bufferSize)
	go func() {
		for p := range w {
			dest.Write(p)
		}
	}()
	return w
}

func (w bufferedWriter) Write(p []byte) (n int, err error) {
	w <- p
	return len(p), nil
}
