package linescounter

import (
	"bufio"
	"bytes"
)

// LinesCounter implements the Writer interface and counts lines
type LinesCounter int

func (c *LinesCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	counter := 0
	for scanner.Scan() {
		counter++
	}
	*c += LinesCounter(counter)
	return counter, nil
}
