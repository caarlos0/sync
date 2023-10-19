package cio

import (
	"context"
	"io"
	"time"
)

// TimeoutReader returns a reader with a timeout.
func TimeoutReader(r io.Reader, d time.Duration) io.Reader {
	return &timeoutReader{r, d}
}

type timeoutReader struct {
	r io.Reader
	d time.Duration
}

// Read implements io.Reader.
func (r *timeoutReader) Read(p []byte) (n int, err error) {
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), r.d)
	defer func() { cancel() }()

	go func() {
		n, err = r.r.Read(p)
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}
