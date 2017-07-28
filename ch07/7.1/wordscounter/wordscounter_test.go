package wordscounter

import "testing"

func TestWrite(t *testing.T) {
	var c WordsCounter
	nwords, _ := c.Write([]byte("One Two Three"))
	if nwords != 3 {
		t.Error("nwords must be 3")
	}
}
