package cio

import (
	"bytes"
	"io"
	"testing"
)

func TestSafeReadWriter(t *testing.T) {
	var b bytes.Buffer
	rw := SafeReadWriter(&b)

	for i := 0; i < 100; i++ {
		go func() {
			_, _ = io.WriteString(rw, "hi")
		}()
		go func() {
			_, _ = io.ReadAll(rw)
		}()
	}
}
