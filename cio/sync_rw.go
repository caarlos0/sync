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
