package countingwriter

import (
	"io"
)

type ByteCounter struct {
	w       io.Writer
	counter int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	bc.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w, 0}
	return &bc, &bc.counter
}
