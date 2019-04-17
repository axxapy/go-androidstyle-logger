package mocks

type Writer struct {
	last []byte
}

func (m *Writer) Write(p []byte) (n int, err error) {
	m.last = p
	return len(p), nil
}

func (m *Writer) Last() []byte {
	return m.last
}

func (m *Writer) Reset() {
	m.last = nil
}