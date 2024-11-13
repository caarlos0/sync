package cio

import (
	"io"
	"sync"
)

// SafeReadWriter is a io.ReadWriter that uses a mutex to lock reads/writes, so
// they don't happen concurrently.
// It can be used to wrap an unsafe ReadWriter.
func SafeReadWriter(rw io.ReadWriter) io.ReadWriter {
	return &safeReadWriter{rw: rw}
}

// safeReadWriter implements io.ReadWriter, but locks reads and writes.
type safeReadWriter struct {
	rw io.ReadWriter
	m  sync.Mutex
}

// Read implements io.ReadWriter.
func (s *safeReadWriter) Read(p []byte) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.rw.Read(p) //nolint: wrapcheck
}

// Write implements io.ReadWriter.
func (s *safeReadWriter) Write(p []byte) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.rw.Write(p) //nolint: wrapcheck
}

// SafeReader is a io.Reader that uses a mutex to lock reads, so
// they it can be called concurrently.
// It can be used to wrap an unsafe Reader.
func SafeReader(r io.Reader) io.Reader {
	return &safeReader{r: r}
}

// safeReader implements io.Reader, but locks reads and writes.
type safeReader struct {
	r io.Reader
	m sync.Mutex
}

// Read implements io.Reader.
func (s *safeReader) Read(p []byte) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.r.Read(p) //nolint: wrapcheck
}

// SafeWriter is a io.Writer that uses a mutex to lock writes, so
// they it can be called concurrently.
// It can be used to wrap an unsafe Writer.
func SafeWriter(w io.Writer) io.Writer {
	return &safeWriter{w: w}
}

// safeWriter implements io.Writer, but locks reads and writes.
type safeWriter struct {
	w io.Writer
	m sync.Mutex
}

// Write implements io.Writer.
func (s *safeWriter) Write(p []byte) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.w.Write(p) //nolint: wrapcheck
}
