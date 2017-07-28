package limit

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitedReader(t *testing.T) {
	limit := 4
	s := "hello world"
	b := &bytes.Buffer{}
	r := LimitedReader(strings.NewReader(s), limit)
	n, _ := b.ReadFrom(r)
	if n != 4 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	trunc := s[0:limit]
	if b.String() != trunc {
		t.Logf(`"%s" != "%s"`, b.String(), trunc)
	}
}
