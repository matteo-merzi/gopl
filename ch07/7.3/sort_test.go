package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
	r := &tree{value: 2}
	add(r, 1)
	add(r, 3)
	if r.String() != "[1, 2, 3]" {
		t.Error("")
	}
}
