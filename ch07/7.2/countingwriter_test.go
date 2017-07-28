package countingwriter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	w, ptrCount := CountingWriter(b)
	data := []byte("hello world")
	w.Write(data)
	if *ptrCount != int64(len(data)) {
		t.Errorf("%d, must be %d\n", *ptrCount, len(data))
	}
}
