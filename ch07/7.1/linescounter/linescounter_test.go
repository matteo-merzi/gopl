package linescounter

import "testing"

func TestWrite(t *testing.T) {
	var c LinesCounter
	nlines, _ := c.Write([]byte(`One
	Two
	Three`))
	if nlines != 3 {
		t.Error("nlines must be 3")
	}
}
