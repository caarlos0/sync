package safe_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/caarlos0/sync/safe"
)

func TestReadWriter(t *testing.T) {
	var b bytes.Buffer
	rw := safe.ReadWriter(&b)

	for i := 0; i < 100; i++ {
		go func() {
			_, _ = io.WriteString(rw, "hi")
		}()
		go func() {
			_, _ = io.ReadAll(rw)
		}()
	}
}

func TestWriter(t *testing.T) {
	var b bytes.Buffer
	rw := safe.Writer(&b)

	for i := 0; i < 100; i++ {
		go func() {
			_, _ = io.WriteString(rw, "hi")
		}()
	}
}

func TestReader(t *testing.T) {
	var b bytes.Buffer
	rw := safe.Reader(&b)

	for i := 0; i < 100; i++ {
		go func() {
			_, _ = io.ReadAll(rw)
		}()
	}
}
