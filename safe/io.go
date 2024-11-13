package safe

import (
	"io"
	"sync"
)

// ReadWriter is a io.ReadWriter that uses a mutex to lock reads/writes, so
// they don't happen concurrently.
// It can be used to wrap an unsafe ReadWriter.
func ReadWriter(rw io.ReadWriter) io.ReadWriter {
	return &readWriter{rw: rw}
}

// readWriter implements io.ReadWriter, but locks reads and writes.
type readWriter struct {
	rw io.ReadWriter
	m  sync.Mutex
}

// Read implements io.ReadWriter.
func (s *readWriter) Read(p []byte) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.rw.Read(p) //nolint: wrapcheck
}

// Write implements io.ReadWriter.
func (s *readWriter) Write(p []byte) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.rw.Write(p) //nolint: wrapcheck
}

// Reader is a io.Reader that uses a mutex to lock reads, so
// they it can be called concurrently.
// It can be used to wrap an unsafe Reader.
func Reader(r io.Reader) io.Reader {
	return &reader{r: r}
}

// reader implements io.Reader, but locks reads and writes.
type reader struct {
	r io.Reader
	m sync.Mutex
}

// Read implements io.Reader.
func (s *reader) Read(p []byte) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.r.Read(p) //nolint: wrapcheck
}

// Writer is a io.Writer that uses a mutex to lock writes, so
// they it can be called concurrently.
// It can be used to wrap an unsafe Writer.
func Writer(w io.Writer) io.Writer {
	return &writer{w: w}
}

// writer implements io.Writer, but locks reads and writes.
type writer struct {
	w io.Writer
	m sync.Mutex
}

// Write implements io.Writer.
func (s *writer) Write(p []byte) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.w.Write(p) //nolint: wrapcheck
}
