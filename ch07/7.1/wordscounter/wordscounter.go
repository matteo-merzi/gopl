package wordscounter

import (
	"bufio"
	"bytes"
)

// WordsCounter implements the Writer interface and counts words
type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordsCounter(count)
	return count, nil
}
