package l

type InMemoryWriter struct {
	last []byte
}

func (m *InMemoryWriter) Write(p []byte) (n int, err error) {
	m.last = p
	return len(p), nil
}

func (m *InMemoryWriter) Last() []byte {
	return m.last
}

func (m *InMemoryWriter) Reset() {
	m.last = nil
}
